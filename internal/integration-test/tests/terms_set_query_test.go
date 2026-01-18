package tests_test

import (
	"integration-tests/model"
	"integration-tests/tests"

	"github.com/Trendyol/es-query-builder/es"
	scriptlanguage "github.com/Trendyol/es-query-builder/es/enums/script-language"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_terms_set_query() {
	// Given - FooDocument kullanarak terms_set testi
	foo1 := model.FooDocument{
		ID:  "terms_set_1",
		Foo: "electric",
	}
	foo2 := model.FooDocument{
		ID:  "terms_set_2",
		Foo: "fire",
	}
	foo3 := model.FooDocument{
		ID:  "terms_set_3",
		Foo: "water",
	}

	s.FooElasticsearchRepository.BulkInsert(s.TestContext, foo1, foo2, foo3)
	tests.WaitExists(s.TestContext, s.FooElasticsearchRepository, foo1.ID)
	tests.WaitExists(s.TestContext, s.FooElasticsearchRepository, foo2.ID)
	tests.WaitExists(s.TestContext, s.FooElasticsearchRepository, foo3.ID)

	query := es.NewQuery(
		es.TermsSet("foo", "electric", "fire", "grass").
			MinimumShouldMatchScript(
				es.ScriptSource("1", scriptlanguage.Painless),
			),
	)

	// When
	result, err := s.FooElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.GreaterOrEqual(s.T(), len(result), 2)

	// electric ve fire olmalı
	_, hasElectric := result["terms_set_1"]
	_, hasFire := result["terms_set_2"]
	assert.True(s.T(), hasElectric)
	assert.True(s.T(), hasFire)

	s.FooElasticsearchRepository.BulkDelete(s.TestContext, foo1.ID, foo2.ID, foo3.ID)
}
