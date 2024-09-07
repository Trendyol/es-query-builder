package testing

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"integration-tests/constants"
	"net/http"
	"strings"
)

type elasticsearchRepository struct {
	client *elasticsearch.Client
}

type ElasticsearchRepository interface {
	Search(query string) ([]FooDocument, error)
	BulkInsert(documents []FooDocument) error
	Insert(document FooDocument) error
	BulkDelete(docIds []string) error
	Delete(docId string) error
	DeleteByQuery(query string) error
	Exists(docId string) bool
}

func NewElasticsearchRepository(client *elasticsearch.Client) ElasticsearchRepository {
	return &elasticsearchRepository{client: client}
}

func (e *elasticsearchRepository) Search(query string) ([]FooDocument, error) {
	res, err := e.client.Search(
		e.client.Search.WithIndex(constants.TestIndex),
		e.client.Search.WithBody(strings.NewReader(query)),
	)
	defer func() {
		if err = res.Body.Close(); err != nil {
			panic(err)
		}
	}()
	if err != nil {
		return nil, fmt.Errorf("failed to execute search request: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d: %s", res.StatusCode, res.String())
	}

	var searchResponse SearchResponse
	if err = json.NewDecoder(res.Body).Decode(&searchResponse); err != nil {
		return nil, fmt.Errorf("failed to decode search response: %w", err)
	}
	if len(searchResponse.Hits.Hits) == constants.Zero {
		return nil, nil
	}

	result := make([]FooDocument, 0)
	for i := range searchResponse.Hits.Hits {
		result = append(result, searchResponse.Hits.Hits[i].Source)
	}

	return result, nil
}

func (e *elasticsearchRepository) BulkInsert(documents []FooDocument) error {
	var bulkRequestBody strings.Builder

	for i := range documents {
		meta := fmt.Sprintf(`{"index":{"_index":"%s","_id":"%s"}}%s`, constants.TestIndex, documents[i], "\n")
		bulkRequestBody.WriteString(meta)

		docJson, err := json.Marshal(documents[i])
		if err != nil {
			return fmt.Errorf("failed to marshal document with Id %s: %w", documents[i].Id, err)
		}
		bulkRequestBody.WriteString(string(docJson) + "\n")
	}
	request := esapi.BulkRequest{
		Body:    strings.NewReader(bulkRequestBody.String()),
		Refresh: constants.True,
	}

	res, err := request.Do(context.Background(), e.client)
	if err != nil {
		return fmt.Errorf("failed to execute bulk insert request: %w", err)
	}
	if res.IsError() {
		return fmt.Errorf("bulk insert request returned error: %s", res.String())
	}
	return nil
}

func (e *elasticsearchRepository) Insert(document FooDocument) error {
	fooDoc, _ := json.Marshal(document)

	request := esapi.IndexRequest{
		Index:      constants.TestIndex,
		DocumentID: document.Id,
		Body:       strings.NewReader(string(fooDoc)),
		Refresh:    constants.True,
	}

	res, err := request.Do(context.Background(), e.client)
	if err != nil {
		return fmt.Errorf("failed to execute insert request: %w", err)
	}
	if res.IsError() {
		return fmt.Errorf("insert request returned error: %s", res.String())
	}
	return err
}

func (e *elasticsearchRepository) Delete(docId string) error {
	request := esapi.DeleteRequest{
		Index:      constants.TestIndex,
		DocumentID: docId,
	}
	res, err := request.Do(context.Background(), e.client)
	if err != nil {
		return fmt.Errorf("failed to execute insert request: %w", err)
	}
	if res.IsError() {
		return fmt.Errorf("insert request returned error: %s", res.String())
	}
	return err
}

func (e *elasticsearchRepository) BulkDelete(docIds []string) error {
	var bulkRequestBody strings.Builder

	for _, docId := range docIds {
		meta := fmt.Sprintf(`{"delete":{"_index":"%s","_id":"%s"}}%s`, constants.TestIndex, docId, "\n")
		bulkRequestBody.WriteString(meta)
	}

	request := esapi.BulkRequest{
		Body:    strings.NewReader(bulkRequestBody.String()),
		Refresh: constants.True,
	}

	res, err := request.Do(context.Background(), e.client)
	if err != nil {
		return fmt.Errorf("failed to execute bulk delete request: %w", err)
	}
	if res.IsError() {
		return fmt.Errorf("bulk delete request returned error: %s", res.String())
	}
	return nil
}

func (e *elasticsearchRepository) DeleteByQuery(query string) error {
	request := esapi.DeleteByQueryRequest{
		Index: []string{constants.TestIndex},
		Body:  strings.NewReader(query),
	}
	res, err := request.Do(context.Background(), e.client)
	if err != nil {
		return fmt.Errorf("failed to execute insert request: %w", err)
	}
	if res.IsError() {
		return fmt.Errorf("insert request returned error: %s", res.String())
	}
	return err
}

func (e *elasticsearchRepository) Exists(docId string) bool {
	request := esapi.ExistsRequest{
		Index:      constants.TestIndex,
		DocumentID: docId,
	}
	res, _ := request.Do(context.Background(), e.client)
	return res.StatusCode == constants.StatusOK
}
