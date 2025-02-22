package tests_test

import (
	"integration-tests/model"
	"integration-tests/tests"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_simple_query_string_with_wildcard_or_operator() {
	// Given
	foo := model.FooDocument{
		ID:  "10",
		Foo: "foo",
	}
	bar := model.FooDocument{
		ID:  "20",
		Foo: "bar",
	}
	george := model.FooDocument{
		ID:  "30",
		Foo: "george orwell",
	}
	s.FooElasticsearchRepository.BulkInsert(s.TestContext, foo, bar, george)
	tests.WaitExists(s.TestContext, s.FooElasticsearchRepository, foo.ID)
	tests.WaitExists(s.TestContext, s.FooElasticsearchRepository, bar.ID)
	tests.WaitExists(s.TestContext, s.FooElasticsearchRepository, george.ID)

	query := es.NewQuery(es.SimpleQueryString[string]("ge* OR bar").AnalyzeWildcard(true))

	// When
	result, err := s.FooElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), len(result), 2)
	assert.Equal(s.T(), result["20"].Foo, "bar")
	assert.Equal(s.T(), result["30"].Foo, "george orwell")

	s.FooElasticsearchRepository.BulkDelete(s.TestContext, foo.ID, bar.ID, george.ID)
}
