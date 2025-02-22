package tests_test

import (
	"integration-tests/model"
	"integration-tests/tests"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_term_query() {
	// Given
	foo := model.FooDocument{
		ID:  "10",
		Foo: "foo",
	}
	bar := model.FooDocument{
		ID:  "20",
		Foo: "bar",
	}

	s.FooElasticsearchRepository.BulkInsert(s.TestContext, foo, bar)
	tests.WaitExists(s.TestContext, s.FooElasticsearchRepository, foo.ID)
	tests.WaitExists(s.TestContext, s.FooElasticsearchRepository, bar.ID)

	query := es.NewQuery(es.Term("foo", "foo"))

	// When
	result, err := s.FooElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 1, len(result))
	assert.Equal(s.T(), "foo", result["10"].Foo)

	s.FooElasticsearchRepository.BulkDelete(s.TestContext, foo.ID, bar.ID)
}
