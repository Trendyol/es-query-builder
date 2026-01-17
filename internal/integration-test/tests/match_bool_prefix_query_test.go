package tests_test

import (
	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_match_bool_prefix_query() {
	// Given
	query := es.NewQuery(
		es.MatchBoolPrefix("name", "char"),
	)

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.GreaterOrEqual(s.T(), len(result), 3)

	// Charmander, Charmeleon, Charizard olmalı
	expectedPokemons := map[string]string{
		"4_5": "charmander",
		"5_6": "charmeleon",
		"6_7": "charizard",
	}

	for pokeID, expectedName := range expectedPokemons {
		pokemon, exists := result[pokeID]
		assert.True(s.T(), exists, "Pokemon %s bulunamadı", pokeID)
		assert.Equal(s.T(), expectedName, pokemon.Name)
	}
}
