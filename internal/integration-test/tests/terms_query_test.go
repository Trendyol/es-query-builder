package tests_test

import (
	"integration-tests/model"
	"integration-tests/tests"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_terms_query() {
	// Given
	doc1 := model.FooDocument{
		ID:  "10",
		Foo: "moon",
	}
	doc2 := model.FooDocument{
		ID:  "20",
		Foo: "mars",
	}
	doc3 := model.FooDocument{
		ID:  "30",
		Foo: "earth",
	}

	s.FooElasticsearchRepository.BulkInsert(s.TestContext, doc1, doc2, doc3)
	tests.WaitExists(s.TestContext, s.FooElasticsearchRepository, doc1.ID)
	tests.WaitExists(s.TestContext, s.FooElasticsearchRepository, doc2.ID)
	tests.WaitExists(s.TestContext, s.FooElasticsearchRepository, doc3.ID)

	query := es.NewQuery(es.Terms("foo", "earth", "mars"))

	// When
	result, err := s.FooElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 2, len(result))
	values := tests.MapValues(result)
	assert.Contains(s.T(), values, doc2)
	assert.Contains(s.T(), values, doc3)

	s.FooElasticsearchRepository.BulkDelete(s.TestContext, doc1.ID, doc2.ID, doc3.ID)
}
