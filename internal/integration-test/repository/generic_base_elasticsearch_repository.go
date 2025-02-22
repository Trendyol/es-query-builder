package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"integration-tests/constants"
	"integration-tests/errorx"
	"integration-tests/model_repository"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type BaseGenericRepository[ID comparable, T any] interface {
	Search(ctx context.Context, query es.Object) (*model_repository.SearchResponse, error)
	GetSearchHits(ctx context.Context, query es.Object) (map[ID]T, error)
	Exists(ctx context.Context, documentID ID) (bool, error)
	Insert(ctx context.Context, document T) error
	BulkInsert(ctx context.Context, documents ...T) error
	Delete(ctx context.Context, documentID ID) error
	DeleteByQuery(ctx context.Context, query es.Object) error
	BulkDelete(ctx context.Context, documentIDs ...ID) error
}

func mapToReader[T any](object T) io.Reader {
	pipeReader, pipeWriter := io.Pipe()
	go func() {
		defer pipeWriter.Close()
		if err := json.NewEncoder(pipeWriter).Encode(object); err != nil {
			pipeWriter.CloseWithError(err)
		}
	}()
	return pipeReader
}

type baseRepository struct {
	Client    *elasticsearch.Client
	IndexName string
}

func newBaseRepository(
	client *elasticsearch.Client,
	indexName string,
) *baseRepository {
	return &baseRepository{
		Client:    client,
		IndexName: indexName,
	}
}

type baseGenericRepository[ID comparable, T any] struct {
	*baseRepository
	mapFunc        func(ID, model_repository.SearchHit) (ID, T, error)
	mapDocIDFunc   func(model_repository.SearchHit) (ID, error)
	mapModelIdFunc func(T) string
}

func NewBaseGenericRepository[ID comparable, T any](
	client *elasticsearch.Client,
	IndexName string,
	mapFunc func(ID, model_repository.SearchHit) (ID, T, error),
	mapDocIDFunc func(model_repository.SearchHit) (ID, error),
	mapModelIdFunc func(T) string,
) BaseGenericRepository[ID, T] {
	return &baseGenericRepository[ID, T]{
		mapFunc:        mapFunc,
		mapDocIDFunc:   mapDocIDFunc,
		mapModelIdFunc: mapModelIdFunc,
		baseRepository: newBaseRepository(client, IndexName),
	}
}

func (repository *baseGenericRepository[ID, T]) Search(ctx context.Context, query es.Object) (*model_repository.SearchResponse, error) {
	res, err := repository.Client.Search(
		repository.Client.Search.WithContext(ctx),
		repository.Client.Search.WithIndex(repository.baseRepository.IndexName),
		repository.Client.Search.WithBody(mapToReader(query)),
	)
	defer func() {
		if err = res.Body.Close(); err != nil {
			panic(err)
		}
	}()
	if err != nil {
		return nil, fmt.Errorf("#baseGenericRepository/Search - failed to execute search request: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("#baseGenericRepository/Search - unexpected status code %d: %s", res.StatusCode, res.String())
	}

	var searchResponse model_repository.SearchResponse
	if err = json.NewDecoder(res.Body).Decode(&searchResponse); err != nil {
		return nil, fmt.Errorf("#baseGenericRepository/NewDecoder_Decode - failed to decode search response: %w", err)
	}
	return &searchResponse, nil
}

func (repository *baseGenericRepository[ID, T]) GetSearchHits(ctx context.Context, query es.Object) (map[ID]T, error) {
	searchResponse, err := repository.Search(ctx, query)
	if err != nil {
		return nil, err
	}
	searchHitMap := make(map[ID]T)
	for _, searchHit := range searchResponse.Hits.Hits {
		docID, err := repository.mapDocIDFunc(searchHit)
		if err != nil {
			return nil, err
		}
		id, mappedHit, err := repository.mapFunc(docID, searchHit)
		if err != nil {
			return nil, err
		}
		searchHitMap[id] = mappedHit
	}
	return searchHitMap, err
}

func (repository *baseGenericRepository[ID, T]) Exists(ctx context.Context, documentID ID) (bool, error) {
	request := esapi.ExistsRequest{
		Index:      repository.IndexName,
		DocumentID: fmt.Sprintf("%v", documentID),
	}
	res, err := request.Do(ctx, repository.Client)
	if err != nil {
		if errorx.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return res.StatusCode == http.StatusOK, nil
}

func (repository *baseGenericRepository[ID, T]) Insert(ctx context.Context, document T) error {
	request := esapi.IndexRequest{
		Index:      repository.IndexName,
		DocumentID: repository.mapModelIdFunc(document),
		Refresh:    constants.True,
		Body:       mapToReader(document),
	}

	res, err := request.Do(ctx, repository.Client)
	defer func() {
		if err = res.Body.Close(); err != nil {
			panic(err)
		}
	}()
	if err != nil {
		return fmt.Errorf("#pokedexElasticsearchRepository - failed to execute insert request: %w", err)
	}
	if res.IsError() {
		return fmt.Errorf("#pokedexElasticsearchRepository - insert request returned error: %s", res.String())
	}
	return err
}

func (repository *baseGenericRepository[ID, T]) BulkInsert(ctx context.Context, documents ...T) error {
	if len(documents) == 0 {
		return nil
	}
	var bulkRequestBody strings.Builder
	bulkRequestBody.Grow(60 * len(documents))
	for i := range documents {
		bulkRequestBody.WriteString(`{"index":{"_index":"`)
		bulkRequestBody.WriteString(repository.IndexName)
		bulkRequestBody.WriteString(`","_id":"`)
		bulkRequestBody.WriteString(repository.mapModelIdFunc(documents[i]))
		bulkRequestBody.WriteString(`"}}`)
		bulkRequestBody.WriteByte('\n')

		docJson, err := json.Marshal(documents[i])
		if err != nil {
			return fmt.Errorf("#pokedexElasticsearchRepository - failed to marshal document with Id %s: %w", repository.mapModelIdFunc(documents[i]), err)
		}

		if bulkRequestBody.Cap() < bulkRequestBody.Len()+len(docJson)+1 {
			bulkRequestBody.Grow(len(docJson) + 1)
		}

		bulkRequestBody.Write(docJson)
		bulkRequestBody.WriteByte('\n')
	}
	request := esapi.BulkRequest{
		Refresh: constants.True,
		Body:    strings.NewReader(bulkRequestBody.String()),
	}

	res, err := request.Do(ctx, repository.Client)
	defer func() {
		if err = res.Body.Close(); err != nil {
			panic(err)
		}
	}()
	if err != nil {
		return fmt.Errorf("#pokedexElasticsearchRepository - failed to execute bulk insert request: %w", err)
	}
	if res.IsError() {
		return fmt.Errorf("#pokedexElasticsearchRepository - bulk insert request returned error: %s", res.String())
	}
	return nil
}

func (repository *baseGenericRepository[ID, T]) Delete(ctx context.Context, docID ID) error {
	request := esapi.DeleteRequest{
		Index:      repository.IndexName,
		DocumentID: fmt.Sprintf("%v", docID),
	}
	res, err := request.Do(ctx, repository.Client)
	defer func() {
		if err = res.Body.Close(); err != nil {
			panic(err)
		}
	}()
	if err != nil {
		return fmt.Errorf("#pokedexElasticsearchRepository - failed to execute delete request: %w", err)
	}
	if res.IsError() {
		return fmt.Errorf("pokedexElasticsearchRepository - delete request returned error: %s", res.String())
	}
	return err
}

func (repository *baseGenericRepository[ID, T]) DeleteByQuery(ctx context.Context, query es.Object) error {
	request := esapi.DeleteByQueryRequest{
		Index: []string{repository.IndexName},
		Body:  mapToReader(query),
	}
	res, err := request.Do(ctx, repository.Client)
	defer func() {
		if err = res.Body.Close(); err != nil {
			panic(err)
		}
	}()
	if err != nil {
		return fmt.Errorf("#pokedexElasticsearchRepository - failed to execute delete by query request: %w", err)
	}
	if res.IsError() {
		return fmt.Errorf("#pokedexElasticsearchRepository - delete by query request returned error: %s", res.String())
	}
	return err
}

func (repository *baseGenericRepository[ID, T]) BulkDelete(ctx context.Context, docIDs ...ID) error {
	if len(docIDs) == 0 {
		return nil
	}
	var bulkRequestBody strings.Builder
	bulkRequestBody.Grow(50 * len(docIDs))
	for _, docID := range docIDs {
		bulkRequestBody.WriteString(`{"delete":{"_index":"`)
		bulkRequestBody.WriteString(repository.IndexName)
		bulkRequestBody.WriteString(`","_id":"`)
		bulkRequestBody.WriteString(fmt.Sprintf("%v", docID))
		bulkRequestBody.WriteString(`"}}`)
		bulkRequestBody.WriteByte('\n')
	}
	request := esapi.BulkRequest{
		Refresh: constants.True,
		Body:    strings.NewReader(bulkRequestBody.String()),
	}

	res, err := request.Do(ctx, repository.Client)
	defer func() {
		if err = res.Body.Close(); err != nil {
			panic(err)
		}
	}()
	if err != nil {
		return fmt.Errorf("#pokedexElasticsearchRepository - failed to execute bulk delete request: %w", err)
	}
	if res.IsError() {
		return fmt.Errorf("#pokedexElasticsearchRepository - bulk delete request returned error: %s", res.String())
	}
	return nil
}
