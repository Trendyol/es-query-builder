package tests_test

import (
	"integration-tests/model"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/bayraktugrul/go-await"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_query_string_with_wildcard_or_operator() {
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

	query := es.NewQuery(es.QueryString("ge* OR bar").AnalyzeWildcard(true))

	// When
	result, err := s.ElasticsearchRepository.Search(query)

	// Then
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 2, len(result))
	assert.Equal(s.T(), "george orwell", result[0].Foo)
	assert.Equal(s.T(), "bar", result[1].Foo)

	s.ElasticsearchRepository.BulkDelete(s.TestContext, foo.ID, bar.ID, george.ID)
}
