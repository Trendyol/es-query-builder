package tests_test

import (
	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_dis_max_query() {
	// Given
	query := es.NewQuery(
		es.DisMax(
			es.Term("name.keyword", "pikachu"),
			es.Term("name.keyword", "bulbasaur"),
		).TieBreaker(0.7),
	)

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 2, len(result))

	expected := map[string]string{
		"25_35": "pikachu",
		"1_1":   "bulbasaur",
	}
	for pokeID, expectedName := range expected {
		pokemon, exists := result[pokeID]
		assert.True(s.T(), exists)
		assert.Equal(s.T(), expectedName, pokemon.Name)
	}
}
