package testing

import (
	"encoding/json"
	"time"

	"github.com/GokselKUCUKSAHIN/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_term_query() {
	// Given
	foo := FooDocument{
		Foo: "foo",
	}
	bar := FooDocument{
		Foo: "bar",
	}

	fooDoc, _ := json.Marshal(foo)
	barDoc, _ := json.Marshal(bar)

	s.ElasticsearchRepository.Insert(testIndexName, "10", string(fooDoc))
	time.Sleep(2 * time.Second)
	s.ElasticsearchRepository.Insert(testIndexName, "20", string(barDoc))
	time.Sleep(2 * time.Second)

	query := es.NewQuery(
		es.Bool().Must(
			es.Term("foo", "foo")),
	)
	bodyJSON, _ := json.Marshal(query)

	// When
	result, err := s.ElasticsearchRepository.Search(testIndexName, string(bodyJSON))

	// Then
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), len(result), 1)
	assert.Equal(s.T(), result[0].Foo, "foo")

	s.ElasticsearchRepository.Delete(testIndexName, "10")
	s.ElasticsearchRepository.Delete(testIndexName, "20")
}

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_terms_query() {
	// Given
	doc1 := FooDocument{
		Foo: "foo",
	}
	doc2 := FooDocument{
		Foo: "mars",
	}
	doc3 := FooDocument{
		Foo: "earth",
	}

	testDoc1, _ := json.Marshal(doc1)
	testDoc2, _ := json.Marshal(doc2)
	testDoc3, _ := json.Marshal(doc3)

	s.ElasticsearchRepository.Insert(testIndexName, "10", string(testDoc1))
	time.Sleep(2 * time.Second)
	s.ElasticsearchRepository.Insert(testIndexName, "20", string(testDoc2))
	time.Sleep(2 * time.Second)
	s.ElasticsearchRepository.Insert(testIndexName, "30", string(testDoc3))
	time.Sleep(2 * time.Second)

	query := es.NewQuery(
		es.Bool().Must(
			es.Terms("foo", "earth", "mars")),
	)
	bodyJSON, _ := json.Marshal(query)

	// When
	result, err := s.ElasticsearchRepository.Search(testIndexName, string(bodyJSON))

	// Then
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), len(result), 2)
	assert.Contains(s.T(), result, doc2)
	assert.Contains(s.T(), result, doc3)

	s.ElasticsearchRepository.Delete(testIndexName, "10")
	s.ElasticsearchRepository.Delete(testIndexName, "20")
	s.ElasticsearchRepository.Delete(testIndexName, "30")
}
