package tests_test

import (
	"integration-tests/model"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/bayraktugrul/go-await"
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

	s.ElasticsearchRepository.BulkInsert(s.TestContext, foo, bar)
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(s.TestContext, foo.ID) })
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(s.TestContext, bar.ID) })

	query := es.NewQuery(es.Term("foo", "foo"))

	// When
	result, err := s.ElasticsearchRepository.Search(query)

	// Then
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 1, len(result))
	assert.Equal(s.T(), "foo", result[0].Foo)

	s.ElasticsearchRepository.BulkDelete(s.TestContext, foo.ID, bar.ID)
}
