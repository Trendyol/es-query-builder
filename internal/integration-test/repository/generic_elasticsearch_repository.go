package repository

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"integration-tests/constants"
	"integration-tests/model"
	"io"
	"net/http"
)

type pokedexElasticsearchRepository struct {
	client *elasticsearch.Client
	index  string
}

type BaseGenericRepository[ID comparable, T any] interface {
	//Search(ctx context.Context, query es.Object) (*model_repository.SearchResponse[T], error)
	Exists(ctx context.Context, documentID ID, routingID string) (bool, error)
	//GetCount(ctx context.Context, query es.Object) (*model_repository.CountResponse, error)
	//GetSearchHits(ctx context.Context, query es.Object) (map[ID]T, error)
	//GetByID(ctx context.Context, documentID ID, routingID string) (*T, error)
	//Insert(ctx context.Context, document T) error
	//BulkInsert(ctx context.Context, documents ...T) error
	//Delete(ctx context.Context, documentID ID) error
	//DeleteByQuery(ctx context.Context, query es.Object) error
	//BulkDelete(ctx context.Context, documentIDs ...ID) error
}

func NewPokedexElasticsearchRepository(client *elasticsearch.Client) BaseGenericRepository[string, model.Pokemon] {
	return &pokedexElasticsearchRepository{
		client: client,
		index:  constants.PokemonIndex,
	}
}

func (repository *pokedexElasticsearchRepository) Exists(ctx context.Context, documentID string, routingID string) (bool, error) {
	request := esapi.ExistsRequest{
		Index:      repository.index,
		DocumentID: documentID,
		Routing:    routingID,
	}
	res, err := request.Do(ctx, repository.client)
	if err != nil {
		// TODO: ignore 404
		return false, err
	}
	return res.StatusCode == http.StatusOK, nil
}

//func (repository *elasticsearchRepository) Search(query es.Object) ([]model.FooDocument, error) {
//	res, err := repository.client.Search(
//		repository.client.Search.WithIndex(constants.TestIndex),
//		repository.client.Search.WithBody(mapToReader(query)),
//	)
//	defer func() {
//		if err = res.Body.Close(); err != nil {
//			panic(err)
//		}
//	}()
//	if err != nil {
//		return nil, fmt.Errorf("failed to execute search request: %w", err)
//	}
//	if res.StatusCode != http.StatusOK {
//		return nil, fmt.Errorf("unexpected status code %d: %s", res.StatusCode, res.String())
//	}
//
//	var searchResponse model.GenericSearchResponse[model.FooDocument]
//	if err = json.NewDecoder(res.Body).Decode(&searchResponse); err != nil {
//		return nil, fmt.Errorf("failed to decode search response: %w", err)
//	}
//	if len(searchResponse.Hits.Hits) == constants.Zero {
//		return nil, nil
//	}
//
//	result := make([]model.FooDocument, 0)
//	for i := range searchResponse.Hits.Hits {
//		result = append(result, searchResponse.Hits.Hits[i].Source)
//	}
//
//	return result, nil
//}
//
//func (repository *elasticsearchRepository) Insert(ctx context.Context, document model.FooDocument) error {
//	request := esapi.IndexRequest{
//		Index:      constants.TestIndex,
//		DocumentID: document.ID,
//		Refresh:    constants.True,
//		Body:       mapToReader(document),
//	}
//
//	res, err := request.Do(ctx, repository.client)
//	defer func() {
//		if err = res.Body.Close(); err != nil {
//			panic(err)
//		}
//	}()
//	if err != nil {
//		return fmt.Errorf("failed to execute insert request: %w", err)
//	}
//	if res.IsError() {
//		return fmt.Errorf("insert request returned error: %s", res.String())
//	}
//	return err
//}
//
//func (repository *elasticsearchRepository) BulkInsert(ctx context.Context, documents ...model.FooDocument) error {
//	if len(documents) == 0 {
//		return nil
//	}
//	var bulkRequestBody strings.Builder
//	bulkRequestBody.Grow(60 * len(documents))
//	for i := range documents {
//		bulkRequestBody.WriteString(`{"index":{"_index":"`)
//		bulkRequestBody.WriteString(constants.TestIndex)
//		bulkRequestBody.WriteString(`","_id":"`)
//		bulkRequestBody.WriteString(documents[i].ID)
//		bulkRequestBody.WriteString(`"}}`)
//		bulkRequestBody.WriteByte('\n')
//
//		docJson, err := json.Marshal(documents[i])
//		if err != nil {
//			return fmt.Errorf("failed to marshal document with Id %s: %w", documents[i].ID, err)
//		}
//
//		if bulkRequestBody.Cap() < bulkRequestBody.Len()+len(docJson)+1 {
//			bulkRequestBody.Grow(len(docJson) + 1)
//		}
//
//		bulkRequestBody.Write(docJson)
//		bulkRequestBody.WriteByte('\n')
//	}
//	request := esapi.BulkRequest{
//		Refresh: constants.True,
//		Body:    strings.NewReader(bulkRequestBody.String()),
//	}
//
//	res, err := request.Do(ctx, repository.client)
//	defer func() {
//		if err = res.Body.Close(); err != nil {
//			panic(err)
//		}
//	}()
//	if err != nil {
//		return fmt.Errorf("failed to execute bulk insert request: %w", err)
//	}
//	if res.IsError() {
//		return fmt.Errorf("bulk insert request returned error: %s", res.String())
//	}
//	return nil
//}
//
//func (repository *elasticsearchRepository) Delete(ctx context.Context, docID string) error {
//	request := esapi.DeleteRequest{
//		Index:      constants.TestIndex,
//		DocumentID: docID,
//	}
//	res, err := request.Do(ctx, repository.client)
//	defer func() {
//		if err = res.Body.Close(); err != nil {
//			panic(err)
//		}
//	}()
//	if err != nil {
//		return fmt.Errorf("failed to execute delete request: %w", err)
//	}
//	if res.IsError() {
//		return fmt.Errorf("delete request returned error: %s", res.String())
//	}
//	return err
//}
//
//func (repository *elasticsearchRepository) DeleteByQuery(ctx context.Context, query es.Object) error {
//	request := esapi.DeleteByQueryRequest{
//		Index: []string{constants.TestIndex},
//		Body:  mapToReader(query),
//	}
//	res, err := request.Do(ctx, repository.client)
//	defer func() {
//		if err = res.Body.Close(); err != nil {
//			panic(err)
//		}
//	}()
//	if err != nil {
//		return fmt.Errorf("failed to execute delete by query request: %w", err)
//	}
//	if res.IsError() {
//		return fmt.Errorf("delete by query request returned error: %s", res.String())
//	}
//	return err
//}
//
//func (repository *elasticsearchRepository) BulkDelete(ctx context.Context, docIDs ...string) error {
//	if len(docIDs) == 0 {
//		return nil
//	}
//	var bulkRequestBody strings.Builder
//	bulkRequestBody.Grow(50 * len(docIDs))
//	for _, docID := range docIDs {
//		bulkRequestBody.WriteString(`{"delete":{"_index":"`)
//		bulkRequestBody.WriteString(constants.TestIndex)
//		bulkRequestBody.WriteString(`","_id":"`)
//		bulkRequestBody.WriteString(docID)
//		bulkRequestBody.WriteString(`"}}`)
//		bulkRequestBody.WriteByte('\n')
//	}
//	request := esapi.BulkRequest{
//		Refresh: constants.True,
//		Body:    strings.NewReader(bulkRequestBody.String()),
//	}
//
//	res, err := request.Do(ctx, repository.client)
//	defer func() {
//		if err = res.Body.Close(); err != nil {
//			panic(err)
//		}
//	}()
//	if err != nil {
//		return fmt.Errorf("failed to execute bulk delete request: %w", err)
//	}
//	if res.IsError() {
//		return fmt.Errorf("bulk delete request returned error: %s", res.String())
//	}
//	return nil
//}
//
//func (repository *elasticsearchRepository) Exists(ctx context.Context, docID string) bool {
//	request := esapi.ExistsRequest{
//		Index:      constants.TestIndex,
//		DocumentID: docID,
//	}
//	res, _ := request.Do(ctx, repository.client)
//	// TODO: handle error properly
//	return res.StatusCode == http.StatusOK
//}

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
