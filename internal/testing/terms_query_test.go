package testing

import (
	"encoding/json"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/bayraktugrul/go-await"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_terms_query() {
	// Given
	doc1 := FooDocument{
		Id:  "10",
		Foo: "foo",
	}
	doc2 := FooDocument{
		Id:  "20",
		Foo: "mars",
	}
	doc3 := FooDocument{
		Id:  "30",
		Foo: "earth",
	}

	s.ElasticsearchRepository.BulkInsert([]FooDocument{doc1, doc2, doc3})
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(doc1.Id) })
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(doc2.Id) })
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(doc3.Id) })

	query := es.NewQuery(
		es.Bool().Must(
			es.Terms("foo", "earth", "mars")),
	)
	bodyJSON, _ := json.Marshal(query)

	// When
	result, err := s.ElasticsearchRepository.Search(string(bodyJSON))

	// Then
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), len(result), 2)
	assert.Contains(s.T(), result, doc2)
	assert.Contains(s.T(), result, doc3)

	s.ElasticsearchRepository.BulkDelete([]string{doc1.Id, doc2.Id, doc3.Id})
}
