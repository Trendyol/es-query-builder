package testing

import (
	"encoding/json"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/bayraktugrul/go-await"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_simple_query_string_with_wildcard_or_operator() {
	// Given
	foo := FooDocument{
		Id:  "10",
		Foo: "foo",
	}
	bar := FooDocument{
		Id:  "20",
		Foo: "bar",
	}
	george := FooDocument{
		Id:  "30",
		Foo: "george orwell",
	}
	s.ElasticsearchRepository.BulkInsert([]FooDocument{foo, bar, george})
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(foo.Id) })
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(bar.Id) })
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(george.Id) })

	query := es.NewQuery(
		es.Bool().Must(
			es.SimpleQueryString[string]("ge* OR bar").AnalyzeWildcard(true)),
	)
	bodyJSON, _ := json.Marshal(query)

	// When
	result, err := s.ElasticsearchRepository.Search(string(bodyJSON))

	// Then
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 2, len(result))
	assert.Equal(s.T(), "bar", result[0].Foo)
	assert.Equal(s.T(), "george orwell", result[1].Foo)

	s.ElasticsearchRepository.BulkDelete([]string{foo.Id, bar.Id, george.Id})
}
