package tests_test

import (
	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_bool_query_with_must() {
	// Given
	query := es.NewQuery(
		es.Bool().Must(
			es.Term("name.keyword", "pikachu"),
			es.Term("isDefault", true),
		),
	)

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 1, len(result))

	pokemon, exists := result["25_35"]
	assert.True(s.T(), exists)
	assert.Equal(s.T(), "pikachu", pokemon.Name)
	assert.True(s.T(), pokemon.IsDefault)
}

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_bool_query_with_should() {
	// Given
	query := es.NewQuery(
		es.Bool().Should(
			es.Term("name.keyword", "pikachu"),
			es.Term("name.keyword", "bulbasaur"),
		).MinimumShouldMatch(1),
	)

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 2, len(result))

	expectedPokemons := map[string]string{
		"25_35": "pikachu",
		"1_1":   "bulbasaur",
	}

	for pokeID, expectedName := range expectedPokemons {
		pokemon, exists := result[pokeID]
		assert.True(s.T(), exists, "Pokemon %s bulunamadı", pokeID)
		assert.Equal(s.T(), expectedName, pokemon.Name)
	}
}

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_bool_query_with_must_not() {
	// Given
	query := es.NewQuery(
		es.Bool().
			Must(es.Nested("types", es.Term("types.name", "fire"))).
			MustNot(es.Nested("abilities", es.Term("abilities.name", "blaze"))),
	)

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)

	// Fire type ama blaze ability olmayan pokemonlar
	for _, pokemon := range result {
		// Her pokemon için type kontrolü
		hasFireType := false
		for _, pokemonType := range pokemon.Types {
			if pokemonType.Name == "fire" {
				hasFireType = true
				break
			}
		}
		assert.True(s.T(), hasFireType, "Pokemon %s fire type değil", pokemon.Name)

		// Blaze ability olmamalı
		for _, ability := range pokemon.Abilities {
			assert.NotEqual(s.T(), "blaze", ability.Name, "Pokemon %s blaze ability'ye sahip olmamalı", pokemon.Name)
		}
	}
}

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_bool_query_with_filter() {
	// Given
	query := es.NewQuery(
		es.Bool().Filter(
			es.Range("baseExperience").GreaterThanOrEqual(200),
			es.Range("weight").LessThan(1000),
		),
	)

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Greater(s.T(), len(result), 0)

	// Tüm sonuçlar baseExperience >= 200 ve weight < 1000 olmalı
	for _, pokemon := range result {
		assert.GreaterOrEqual(s.T(), pokemon.BaseExperience, uint16(200))
		assert.Less(s.T(), pokemon.Weight, uint16(1000))
	}
}
