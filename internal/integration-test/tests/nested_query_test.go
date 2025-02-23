package tests_test

import (
	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_nested_query() {
	// Given
	query := es.NewQuery(
		es.Bool().Must(
			es.Nested("types", es.Term("types.name", "fire")),
			es.Nested("abilities", es.Term("abilities.name", "blaze")),
		),
	)

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 6, len(result))
	for pokeID, pokemon := range result {
		switch pokeID {
		case "4_5":
			assert.Equal(s.T(), "charmander", pokemon.Name)
		case "5_6":
			assert.Equal(s.T(), "charmeleon", pokemon.Name)
		case "6_7":
			assert.Equal(s.T(), "charizard", pokemon.Name)
		case "155_252":
			assert.Equal(s.T(), "cyndaquil", pokemon.Name)
		case "156_253":
			assert.Equal(s.T(), "quilava", pokemon.Name)
		case "157_254":
			assert.Equal(s.T(), "typhlosion", pokemon.Name)
		default:
			s.T().FailNow()
		}
	}
}

func (s *testSuite) Test_it_should_return_documents_that_by_nested_nested_query_query() {
	// Given
	query := es.NewQuery(
		es.Nested("moves",
			es.Bool().Must(
				es.Term("moves.name", "thunderbolt"),
				es.Nested("moves.versionGroupDetails",
					es.Range("moves.versionGroupDetails.levelLearnedAt").
						LessThanOrEqual(50),
				),
			),
		),
	).Size(100)

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 11, len(result))
	for pokeID, pokemon := range result {
		switch pokeID {
		case "81_132":
			assert.Equal(s.T(), "magnemite", pokemon.Name)
		case "82_133":
			assert.Equal(s.T(), "magneton", pokemon.Name)
		case "92_147":
			assert.Equal(s.T(), "gastly", pokemon.Name)
		case "93_148":
			assert.Equal(s.T(), "haunter", pokemon.Name)
		case "100_158":
			assert.Equal(s.T(), "voltorb", pokemon.Name)
		case "101_159":
			assert.Equal(s.T(), "electrode", pokemon.Name)
		case "109_173":
			assert.Equal(s.T(), "koffing", pokemon.Name)
		case "110_174":
			assert.Equal(s.T(), "weezing", pokemon.Name)
		case "179_273":
			assert.Equal(s.T(), "mareep", pokemon.Name)
		case "200_295":
			assert.Equal(s.T(), "misdreavus", pokemon.Name)
		case "233_226":
			assert.Equal(s.T(), "porygon2", pokemon.Name)
		default:
			s.T().FailNow()
		}
	}
}
