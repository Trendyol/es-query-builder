package testing

import (
	"encoding/json"
	"github.com/Trendyol/es-query-builder/es"
	"github.com/bayraktugrul/go-await"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_regexp_query() {
	// Given
	foo := FooDocument{
		Id:  "10",
		Foo: "foo",
	}
	bar := FooDocument{
		Id:  "20",
		Foo: "bar",
	}
	georgeOrwell := FooDocument{
		Id:  "30",
		Foo: "george orwell",
	}
	georgeBest := FooDocument{
		Id:  "40",
		Foo: "george best",
	}

	s.ElasticsearchRepository.BulkInsert([]FooDocument{foo, bar, georgeOrwell, georgeBest})
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(foo.Id) })
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(bar.Id) })
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(georgeOrwell.Id) })
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(georgeBest.Id) })

	//f* OR bar
	query := es.NewQuery(
		es.Regexp("foo", "george.*"),
	)
	bodyJSON, _ := json.Marshal(query)

	// When
	result, err := s.ElasticsearchRepository.Search(string(bodyJSON))

	// Then
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), len(result), 2)
	assert.Equal(s.T(), result[0].Foo, "george orwell")
	assert.Equal(s.T(), result[1].Foo, "george best")

	s.ElasticsearchRepository.BulkDelete([]string{foo.Id, bar.Id, georgeOrwell.Id, georgeBest.Id})
}
