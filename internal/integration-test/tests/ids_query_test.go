package tests_test

import (
	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_documents_by_ids_query() {
	// Given
	query := es.NewQuery(
		es.IDs("25_35", "1_1", "4_5"),
	)

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 3, len(result))

	expectedPokemons := map[string]string{
		"25_35": "pikachu",
		"1_1":   "bulbasaur",
		"4_5":   "charmander",
	}

	for pokeID, expectedName := range expectedPokemons {
		pokemon, exists := result[pokeID]
		assert.True(s.T(), exists, "Pokemon %s bulunamadı", pokeID)
		assert.Equal(s.T(), expectedName, pokemon.Name)
	}
}
