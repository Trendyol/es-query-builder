package repository

import (
	"encoding/json"

	"integration-tests/constants"
	"integration-tests/model"
	"integration-tests/model_repository"

	"github.com/elastic/go-elasticsearch/v8"
)

type fooElasticsearchRepository struct {
	BaseGenericRepository[string, model.FooDocument]
}

func NewFooElasticsearchRepository(client *elasticsearch.Client) BaseGenericRepository[string, model.FooDocument] {
	return &fooElasticsearchRepository{
		NewBaseGenericRepository(client, constants.TestIndex, "fooRepository", mapToFoo, model_repository.MapToId, mapToFooId),
	}
}

func mapToFooId(foo model.FooDocument) string {
	return foo.ID
}

func mapToFoo(docID string, searchHit model_repository.SearchHit) (string, model.FooDocument, error) {
	var foo model.FooDocument
	if err := json.Unmarshal(searchHit.Source, &foo); err != nil {
		return "", model.FooDocument{}, err
	}
	return docID, foo, nil
}
