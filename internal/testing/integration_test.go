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
		Foo: "foo",
	}
	bar := FooDocument{
		Foo: "bar",
	}

	fooDoc, _ := json.Marshal(foo)
	barDoc, _ := json.Marshal(bar)

	s.ElasticsearchRepository.Insert(testIndexName, "10", string(fooDoc))
	s.ElasticsearchRepository.Insert(testIndexName, "20", string(barDoc))
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(testIndexName, "10") })
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(testIndexName, "20") })

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
	s.ElasticsearchRepository.Insert(testIndexName, "20", string(testDoc2))
	s.ElasticsearchRepository.Insert(testIndexName, "30", string(testDoc3))
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(testIndexName, "10") })
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(testIndexName, "20") })
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(testIndexName, "30") })

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

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_query_string_with_wildcard_or_operator() {
	// Given
	foo := FooDocument{
		Foo: "foo",
	}
	bar := FooDocument{
		Foo: "bar",
	}
	george := FooDocument{
		Foo: "george orwell",
	}

	fooDoc, _ := json.Marshal(foo)
	barDoc, _ := json.Marshal(bar)
	georgeDoc, _ := json.Marshal(george)

	s.ElasticsearchRepository.Insert(testIndexName, "10", string(fooDoc))
	s.ElasticsearchRepository.Insert(testIndexName, "20", string(barDoc))
	s.ElasticsearchRepository.Insert(testIndexName, "30", string(georgeDoc))
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(testIndexName, "10") })
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(testIndexName, "20") })
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(testIndexName, "30") })
	//f* OR bar
	query := es.NewQuery(
		es.Bool().Must(
			es.QueryString[string]("ge* OR bar").AnalyzeWildcard(true)),
	)
	bodyJSON, _ := json.Marshal(query)

	// When
	result, err := s.ElasticsearchRepository.Search(testIndexName, string(bodyJSON))

	// Then
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), len(result), 2)
	assert.Equal(s.T(), result[0].Foo, "george orwell")
	assert.Equal(s.T(), result[1].Foo, "bar")

	s.ElasticsearchRepository.Delete(testIndexName, "10")
	s.ElasticsearchRepository.Delete(testIndexName, "20")
	s.ElasticsearchRepository.Delete(testIndexName, "30")
}

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_regexp_query() {
	// Given
	foo := FooDocument{
		Foo: "foo",
	}
	bar := FooDocument{
		Foo: "bar",
	}
	georgeOrwell := FooDocument{
		Foo: "george orwell",
	}
	georgeBest := FooDocument{
		Foo: "george best",
	}
	fooDoc, _ := json.Marshal(foo)
	barDoc, _ := json.Marshal(bar)
	georgeOrwellDoc, _ := json.Marshal(georgeOrwell)
	georgeBestDoc, _ := json.Marshal(georgeBest)

	s.ElasticsearchRepository.Insert(testIndexName, "10", string(fooDoc))
	s.ElasticsearchRepository.Insert(testIndexName, "20", string(barDoc))
	s.ElasticsearchRepository.Insert(testIndexName, "30", string(georgeOrwellDoc))
	s.ElasticsearchRepository.Insert(testIndexName, "40", string(georgeBestDoc))
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(testIndexName, "10") })
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(testIndexName, "20") })
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(testIndexName, "30") })
	await.New().Await(func() bool { return s.ElasticsearchRepository.Exists(testIndexName, "40") })
	//f* OR bar
	query := es.NewQuery(
		es.Regexp("foo", "george.*"),
	)
	bodyJSON, _ := json.Marshal(query)

	// When
	result, err := s.ElasticsearchRepository.Search(testIndexName, string(bodyJSON))

	// Then
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), len(result), 2)
	assert.Equal(s.T(), result[0].Foo, "george orwell")
	assert.Equal(s.T(), result[1].Foo, "george best")

	s.ElasticsearchRepository.Delete(testIndexName, "10")
	s.ElasticsearchRepository.Delete(testIndexName, "20")
	s.ElasticsearchRepository.Delete(testIndexName, "30")
}
