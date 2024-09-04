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
	Search(index, query string) ([]FooDocument, error)
	Insert(indexName, docId, document string) error
	Delete(indexName, docId string) error
	DeleteByQuery(indexName, query string) error
	Exists(indexName, docId string) bool
}

func NewElasticsearchRepository(client *elasticsearch.Client) ElasticsearchRepository {
	return &elasticsearchRepository{client: client}
}

func (e *elasticsearchRepository) Search(index, query string) ([]FooDocument, error) {
	res, err := e.client.Search(
		e.client.Search.WithIndex(index),
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

func (e *elasticsearchRepository) Insert(indexName, docId, document string) error {
	request := esapi.IndexRequest{
		Index:      indexName,
		DocumentID: docId,
		Body:       strings.NewReader(document),
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

func (e *elasticsearchRepository) Delete(indexName, docId string) error {
	request := esapi.DeleteRequest{
		Index:      indexName,
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

func (e *elasticsearchRepository) DeleteByQuery(indexName, query string) error {
	request := esapi.DeleteByQueryRequest{
		Index: []string{indexName},
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

func (e *elasticsearchRepository) Exists(indexName, docId string) bool {
	request := esapi.ExistsRequest{
		Index:      indexName,
		DocumentID: docId,
	}
	res, _ := request.Do(context.Background(), e.client)
	return res.StatusCode == constants.StatusOK
}
