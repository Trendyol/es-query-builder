package tests_test

import (
	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_documents_that_have_field_with_exists_query() {
	// Given
	query := es.NewQuery(
		es.Exists("name"),
	).Size(300)

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Greater(s.T(), len(result), 200) // Tüm pokemonların name alanı var
}

func (s *testSuite) Test_it_should_return_only_default_pokemons_with_exists_query() {
	// Given
	query := es.NewQuery(
		es.Bool().Must(
			es.Exists("isDefault"),
			es.Term("isDefault", true),
		),
	).Size(300)

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Greater(s.T(), len(result), 0)

	// Tüm sonuçlar isDefault=true olmalı
	for _, pokemon := range result {
		assert.True(s.T(), pokemon.IsDefault)
	}
}
