package tests_test

import (
	"encoding/json"
	"integration-tests/model"
	"integration-tests/tests"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_collapse_hits_by_field() {
	// Given
	docs := []model.FooDocument{
		{ID: "collapse-1", Foo: "group-a"},
		{ID: "collapse-2", Foo: "group-a"},
		{ID: "collapse-3", Foo: "group-a"},
		{ID: "collapse-4", Foo: "group-b"},
		{ID: "collapse-5", Foo: "group-b"},
	}
	s.FooElasticsearchRepository.BulkInsert(s.TestContext, docs...)
	for _, doc := range docs {
		tests.WaitExists(s.TestContext, s.FooElasticsearchRepository, doc.ID)
	}

	queryWithoutCollapse := es.NewQuery(es.Terms("foo", "group-a", "group-b")).Size(10)
	queryWithCollapse := es.NewQuery(es.Terms("foo", "group-a", "group-b")).
		Size(10).
		Collapse(es.FieldCollapse("foo"))

	ids := make([]string, 0, len(docs))
	for _, doc := range docs {
		ids = append(ids, doc.ID)
	}
	defer s.FooElasticsearchRepository.BulkDelete(s.TestContext, ids...)

	// When
	responseWithoutCollapse, err := s.FooElasticsearchRepository.Search(s.TestContext, queryWithoutCollapse)
	assert.Nil(s.T(), err)

	responseWithCollapse, err := s.FooElasticsearchRepository.Search(s.TestContext, queryWithCollapse)
	assert.Nil(s.T(), err)

	// Then
	assert.NotNil(s.T(), responseWithoutCollapse.Hits)
	assert.NotNil(s.T(), responseWithCollapse.Hits)
	assert.Equal(s.T(), 5, len(responseWithoutCollapse.Hits.Hits))
	assert.Equal(s.T(), 2, len(responseWithCollapse.Hits.Hits))

	seen := make(map[string]bool)
	for _, hit := range responseWithCollapse.Hits.Hits {
		var foo model.FooDocument
		assert.NoError(s.T(), json.Unmarshal(hit.Source, &foo))
		assert.False(s.T(), seen[foo.Foo], "collapse should return at most one hit per foo value")
		seen[foo.Foo] = true
	}
	assert.True(s.T(), seen["group-a"])
	assert.True(s.T(), seen["group-b"])
}

func (s *testSuite) Test_it_should_collapse_hits_with_inner_hits() {
	// Given
	docs := []model.FooDocument{
		{ID: "collapse-ih-1", Foo: "group-x"},
		{ID: "collapse-ih-2", Foo: "group-x"},
		{ID: "collapse-ih-3", Foo: "group-x"},
		{ID: "collapse-ih-4", Foo: "group-y"},
	}
	s.FooElasticsearchRepository.BulkInsert(s.TestContext, docs...)
	for _, doc := range docs {
		tests.WaitExists(s.TestContext, s.FooElasticsearchRepository, doc.ID)
	}

	ids := make([]string, 0, len(docs))
	for _, doc := range docs {
		ids = append(ids, doc.ID)
	}
	defer s.FooElasticsearchRepository.BulkDelete(s.TestContext, ids...)

	query := es.NewQuery(es.Terms("foo", "group-x", "group-y")).
		Size(10).
		Collapse(
			es.FieldCollapse("foo").
				InnerHits(es.InnerHits().Name("collapsed_docs").Size(3)).
				MaxConcurrentGroupSearches(4),
		)

	// When
	response, err := s.FooElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), response.Hits)
	assert.Equal(s.T(), 2, len(response.Hits.Hits))

	hit := response.Hits.Hits[0]
	assert.NotEmpty(s.T(), hit.InnerHits)

	raw, ok := hit.InnerHits["collapsed_docs"]
	assert.True(s.T(), ok)

	var innerHits struct {
		Hits struct {
			Hits []struct {
				ID string `json:"_id"`
			} `json:"hits"`
		} `json:"hits"`
	}
	assert.NoError(s.T(), json.Unmarshal(raw, &innerHits))
	assert.Greater(s.T(), len(innerHits.Hits.Hits), 0)
	assert.LessOrEqual(s.T(), len(innerHits.Hits.Hits), 3)
}
