package testing

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"net/http"
	"strings"
)

const zero = 0

type elasticsearchRepository struct {
	client *elasticsearch.Client
}

type ElasticsearchRepository interface {
	Search(index, query string) ([]FooDocument, error)
	Insert(indexName, docId, document string) error
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
		if err := res.Body.Close(); err != nil {
			panic(err)
		}
	}()
	if err != nil || res.StatusCode != http.StatusOK {
		return []FooDocument{}, errors.New(fmt.Sprintf("error getting search response, status: %d err: %+v", res.StatusCode, err))
	}

	var searchResponse SearchResponse
	if err := json.NewDecoder(res.Body).Decode(&searchResponse); err != nil {
		return []FooDocument{}, err
	}
	if len(searchResponse.Hits.Hits) == zero {
		return []FooDocument{}, nil
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
		Refresh:    "true",
	}

	_, err := request.Do(context.Background(), e.client)
	return err
}
