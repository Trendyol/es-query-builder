package tests_test

import (
	"integration-tests/model"
	"integration-tests/tests"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_regexp_query() {
	// Given
	foo := model.FooDocument{
		ID:  "10",
		Foo: "foo",
	}
	bar := model.FooDocument{
		ID:  "20",
		Foo: "bar",
	}
	georgeOrwell := model.FooDocument{
		ID:  "30",
		Foo: "george orwell",
	}
	georgeBest := model.FooDocument{
		ID:  "40",
		Foo: "george best",
	}

	s.FooElasticsearchRepository.BulkInsert(s.TestContext, foo, bar, georgeOrwell, georgeBest)
	tests.WaitExists(s.TestContext, s.FooElasticsearchRepository, foo.ID)
	tests.WaitExists(s.TestContext, s.FooElasticsearchRepository, bar.ID)
	tests.WaitExists(s.TestContext, s.FooElasticsearchRepository, georgeOrwell.ID)
	tests.WaitExists(s.TestContext, s.FooElasticsearchRepository, georgeBest.ID)

	//f* OR bar
	query := es.NewQuery(es.Regexp("foo", "george.*"))

	// When
	result, err := s.FooElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), len(result), 2)
	assert.Equal(s.T(), result["30"].Foo, "george orwell")
	assert.Equal(s.T(), result["40"].Foo, "george best")

	s.FooElasticsearchRepository.BulkDelete(s.TestContext, foo.ID, bar.ID, georgeOrwell.ID, georgeBest.ID)
}
