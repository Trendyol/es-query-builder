package tests_test

import (
	"integration-tests/model"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/bayraktugrul/go-await"
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
	s.ElasticsearchRepository.BulkInsert(s.TestContext, foo, bar, george)
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(s.TestContext, foo.ID) })
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(s.TestContext, bar.ID) })
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(s.TestContext, george.ID) })

	query := es.NewQuery(es.SimpleQueryString[string]("ge* OR bar").AnalyzeWildcard(true))

	// When
	result, err := s.ElasticsearchRepository.Search(query)

	// Then
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), len(result), 2)
	assert.Equal(s.T(), result[0].Foo, "bar")
	assert.Equal(s.T(), result[1].Foo, "george orwell")

	s.ElasticsearchRepository.BulkDelete(s.TestContext, foo.ID, bar.ID, george.ID)
}
