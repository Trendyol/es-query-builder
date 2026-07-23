package tests_test

import (
	"integration-tests/model"
	"integration-tests/tests"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_wildcard_query() {
	// Given
	foo := model.FooDocument{ID: "10", Foo: "foobar"}
	bar := model.FooDocument{ID: "20", Foo: "bar"}
	baz := model.FooDocument{ID: "30", Foo: "foobaz"}

	s.FooElasticsearchRepository.BulkInsert(s.TestContext, foo, bar, baz)
	tests.WaitExists(s.TestContext, s.FooElasticsearchRepository, foo.ID)
	tests.WaitExists(s.TestContext, s.FooElasticsearchRepository, bar.ID)
	tests.WaitExists(s.TestContext, s.FooElasticsearchRepository, baz.ID)

	query := es.NewQuery(es.Wildcard("foo", "foo*"))

	// When
	result, err := s.FooElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 2, len(result))
	assert.Equal(s.T(), "foobar", result["10"].Foo)
	assert.Equal(s.T(), "foobaz", result["30"].Foo)

	s.FooElasticsearchRepository.BulkDelete(s.TestContext, foo.ID, bar.ID, baz.ID)
}
