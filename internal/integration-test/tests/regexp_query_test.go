package tests_test

import (
	"integration-tests/model"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/bayraktugrul/go-await"
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

	s.ElasticsearchRepository.BulkInsert(s.TestContext, foo, bar, georgeOrwell, georgeBest)
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(s.TestContext, foo.ID) })
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(s.TestContext, bar.ID) })
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(s.TestContext, georgeOrwell.ID) })
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(s.TestContext, georgeBest.ID) })

	//f* OR bar
	query := es.NewQuery(es.Regexp("foo", "george.*"))

	// When
	result, err := s.ElasticsearchRepository.Search(query)

	// Then
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), len(result), 2)
	assert.Equal(s.T(), result[0].Foo, "george orwell")
	assert.Equal(s.T(), result[1].Foo, "george best")

	s.ElasticsearchRepository.BulkDelete(s.TestContext, foo.ID, bar.ID, georgeOrwell.ID, georgeBest.ID)
}
