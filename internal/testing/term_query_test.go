package testing

import (
	"encoding/json"
	"github.com/Trendyol/es-query-builder/es"
	"github.com/bayraktugrul/go-await"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_term_query() {
	// Given
	foo := FooDocument{
		Id:  "10",
		Foo: "foo",
	}
	bar := FooDocument{
		Id:  "20",
		Foo: "bar",
	}

	s.ElasticsearchRepository.BulkInsert([]FooDocument{foo, bar})
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(foo.Id) })
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(bar.Id) })

	query := es.NewQuery(
		es.Bool().Must(
			es.Term("foo", "foo")),
	)
	bodyJSON, _ := json.Marshal(query)

	// When
	result, err := s.ElasticsearchRepository.Search(string(bodyJSON))

	// Then
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), len(result), 1)
	assert.Equal(s.T(), result[0].Foo, "foo")

	s.ElasticsearchRepository.BulkDelete([]string{foo.Id, bar.Id})
}
