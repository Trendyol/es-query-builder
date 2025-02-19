package tests_test

import (
	"integration-tests/model"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/bayraktugrul/go-await"
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

	s.ElasticsearchRepository.BulkInsert(s.TestContext, doc1, doc2, doc3)
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(s.TestContext, doc1.ID) })
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(s.TestContext, doc2.ID) })
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(s.TestContext, doc3.ID) })

	query := es.NewQuery(es.Terms("foo", "earth", "mars"))

	// When
	result, err := s.ElasticsearchRepository.Search(query)

	// Then
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 2, len(result))
	assert.Contains(s.T(), result, doc2)
	assert.Contains(s.T(), result, doc3)

	s.ElasticsearchRepository.BulkDelete(s.TestContext, doc1.ID, doc2.ID, doc3.ID)
}
