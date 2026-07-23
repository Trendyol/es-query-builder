package tests_test

import (
	"integration-tests/model"
	"integration-tests/tests"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_fuzzy_query() {
	// Given
	apple := model.FooDocument{ID: "10", Foo: "apple"}
	apply := model.FooDocument{ID: "20", Foo: "apply"}
	banana := model.FooDocument{ID: "30", Foo: "banana"}

	s.FooElasticsearchRepository.BulkInsert(s.TestContext, apple, apply, banana)
	tests.WaitExists(s.TestContext, s.FooElasticsearchRepository, apple.ID)
	tests.WaitExists(s.TestContext, s.FooElasticsearchRepository, apply.ID)
	tests.WaitExists(s.TestContext, s.FooElasticsearchRepository, banana.ID)

	query := es.NewQuery(
		es.Fuzzy("foo", "appel").Fuzziness("AUTO"),
	)

	// When
	result, err := s.FooElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 1, len(result))
	assert.Equal(s.T(), "apple", result["10"].Foo)

	s.FooElasticsearchRepository.BulkDelete(s.TestContext, apple.ID, apply.ID, banana.ID)
}
