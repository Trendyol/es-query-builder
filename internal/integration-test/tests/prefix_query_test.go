package tests_test

import (
	"integration-tests/model"
	"integration-tests/tests"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_prefix_query() {
	// Given
	tyPhone := model.FooDocument{ID: "10", Foo: "TY-phone"}
	tyWatch := model.FooDocument{ID: "20", Foo: "TY-watch"}
	other := model.FooDocument{ID: "30", Foo: "XX-bag"}

	s.FooElasticsearchRepository.BulkInsert(s.TestContext, tyPhone, tyWatch, other)
	tests.WaitExists(s.TestContext, s.FooElasticsearchRepository, tyPhone.ID)
	tests.WaitExists(s.TestContext, s.FooElasticsearchRepository, tyWatch.ID)
	tests.WaitExists(s.TestContext, s.FooElasticsearchRepository, other.ID)

	query := es.NewQuery(es.Prefix("foo", "TY-"))

	// When
	result, err := s.FooElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 2, len(result))
	assert.Equal(s.T(), "TY-phone", result["10"].Foo)
	assert.Equal(s.T(), "TY-watch", result["20"].Foo)

	s.FooElasticsearchRepository.BulkDelete(s.TestContext, tyPhone.ID, tyWatch.ID, other.ID)
}
