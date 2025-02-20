package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"integration-tests/constants"
	"integration-tests/errorx"
	"integration-tests/model"
	"integration-tests/model_repository"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type pokedexElasticsearchRepository struct {
	client *elasticsearch.Client
	index  string
}

func NewPokedexElasticsearchRepository(client *elasticsearch.Client) BaseGenericRepository[string, model.Pokemon] {
	return &pokedexElasticsearchRepository{
		client: client,
		index:  constants.PokemonIndex,
	}
}

func (repository *pokedexElasticsearchRepository) Search(ctx context.Context, query es.Object) (*model_repository.SearchResponse, error) {
	res, err := repository.client.Search(
		repository.client.Search.WithContext(ctx),
		repository.client.Search.WithIndex(repository.index),
		repository.client.Search.WithBody(mapToReader(query)),
	)
	defer func() {
		if err = res.Body.Close(); err != nil {
			panic(err)
		}
	}()
	if err != nil {
		return nil, fmt.Errorf("#pokedexElasticsearchRepository/Search - failed to execute search request: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("#pokedexElasticsearchRepository - unexpected status code %d: %s", res.StatusCode, res.String())
	}

	var searchResponse model_repository.SearchResponse
	if err = json.NewDecoder(res.Body).Decode(&searchResponse); err != nil {
		return nil, fmt.Errorf("#pokedexElasticsearchRepository/NewDecoder_Decode - failed to decode search response: %w", err)
	}
	return &searchResponse, nil
}

func (repository *pokedexElasticsearchRepository) Exists(ctx context.Context, documentID string) (bool, error) {
	request := esapi.ExistsRequest{
		Index:      repository.index,
		DocumentID: documentID,
	}
	res, err := request.Do(ctx, repository.client)
	if err != nil {
		if errorx.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return res.StatusCode == http.StatusOK, nil
}

func (repository *pokedexElasticsearchRepository) Insert(ctx context.Context, document model.Pokemon) error {
	request := esapi.IndexRequest{
		Index:      repository.index,
		DocumentID: document.GetDocumentID(),
		Refresh:    constants.True,
		Body:       mapToReader(document),
	}

	res, err := request.Do(ctx, repository.client)
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

func (repository *pokedexElasticsearchRepository) BulkInsert(ctx context.Context, documents ...model.Pokemon) error {
	if len(documents) == 0 {
		return nil
	}
	var bulkRequestBody strings.Builder
	bulkRequestBody.Grow(60 * len(documents))
	for i := range documents {
		bulkRequestBody.WriteString(`{"index":{"_index":"`)
		bulkRequestBody.WriteString(repository.index)
		bulkRequestBody.WriteString(`","_id":"`)
		bulkRequestBody.WriteString(documents[i].GetDocumentID())
		bulkRequestBody.WriteString(`"}}`)
		bulkRequestBody.WriteByte('\n')

		docJson, err := json.Marshal(documents[i])
		if err != nil {
			return fmt.Errorf("#pokedexElasticsearchRepository - failed to marshal document with Id %s: %w", documents[i].GetDocumentID(), err)
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

	res, err := request.Do(ctx, repository.client)
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

func (repository *pokedexElasticsearchRepository) Delete(ctx context.Context, docID string) error {
	request := esapi.DeleteRequest{
		Index:      repository.index,
		DocumentID: docID,
	}
	res, err := request.Do(ctx, repository.client)
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

func (repository *pokedexElasticsearchRepository) DeleteByQuery(ctx context.Context, query es.Object) error {
	request := esapi.DeleteByQueryRequest{
		Index: []string{repository.index},
		Body:  mapToReader(query),
	}
	res, err := request.Do(ctx, repository.client)
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

func (repository *pokedexElasticsearchRepository) BulkDelete(ctx context.Context, docIDs ...string) error {
	if len(docIDs) == 0 {
		return nil
	}
	var bulkRequestBody strings.Builder
	bulkRequestBody.Grow(50 * len(docIDs))
	for _, docID := range docIDs {
		bulkRequestBody.WriteString(`{"delete":{"_index":"`)
		bulkRequestBody.WriteString(repository.index)
		bulkRequestBody.WriteString(`","_id":"`)
		bulkRequestBody.WriteString(docID)
		bulkRequestBody.WriteString(`"}}`)
		bulkRequestBody.WriteByte('\n')
	}
	request := esapi.BulkRequest{
		Refresh: constants.True,
		Body:    strings.NewReader(bulkRequestBody.String()),
	}

	res, err := request.Do(ctx, repository.client)
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
