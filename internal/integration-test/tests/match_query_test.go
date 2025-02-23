package tests_test

import (
	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_match_query_with_fuzzy() {
	// Given
	query := es.NewQuery(
		es.Match("name", "pikchu").
			Fuzziness("AUTO"),
	)

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 2, len(result))
	for pokeID, pokemon := range result {
		switch pokeID {
		case "25_35":
			assert.Equal(s.T(), "pikachu", pokemon.Name)
		case "172_34":
			assert.Equal(s.T(), "pichu", pokemon.Name)
		default:
			s.T().FailNow()
		}
	}
}
