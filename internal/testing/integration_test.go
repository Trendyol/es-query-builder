package testing

import (
	"encoding/json"
	"time"

	"github.com/GokselKUCUKSAHIN/es-query-builder/es"
	"github.com/GokselKUCUKSAHIN/es-query-builder/test/assert"
)

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_term_query() {
	//given
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

	//when
	result, err := s.ElasticsearchRepository.Search(testIndexName, string(bodyJSON))

	//then
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), len(result), 1)
	assert.Equal(s.T(), result[0].Foo, "foo")
}
