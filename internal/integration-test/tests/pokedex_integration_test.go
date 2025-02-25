package tests_test

import (
	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_documents_pokedex_it_1() {
	// Given
	query := es.NewQuery(es.Match("name", "pikachu"))

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 1, len(result))
	for pokeID, pokemon := range result {
		switch pokeID {
		case "25_35":
			assert.Equal(s.T(), "pikachu", pokemon.Name)
		default:
			s.T().FailNow()
		}
	}
}

func (s *testSuite) Test_it_should_return_documents_pokedex_it_2() {
	// Given
	query := es.NewQuery(es.Term("id", 25))

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 1, len(result))
	for pokeID, pokemon := range result {
		switch pokeID {
		case "25_35":
			assert.Equal(s.T(), "pikachu", pokemon.Name)
		default:
			s.T().FailNow()
		}
	}
}

func (s *testSuite) Test_it_should_return_documents_pokedex_it_3() {
	// Given
	query := es.NewQuery(es.Nested("abilities", es.Term("abilities.name", "static"))).Size(100)

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 11, len(result))
}

func (s *testSuite) Test_it_should_return_documents_pokedex_it_4() {
	// Given
	query := es.NewQuery(es.Nested("moves", es.Term("moves.name", "thunderbolt"))).Size(100)

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 11, len(result))
}

func (s *testSuite) Test_it_should_return_documents_pokedex_it_5() {
	// Given
	query := es.NewQuery(es.Nested("types", es.Term("types.name", "electric"))).Size(100)

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 17, len(result))
}

func (s *testSuite) Test_it_should_return_documents_pokedex_it_6() {
	// Given
	query := es.NewQuery(es.Nested("moves.versionGroupDetails",
		es.Bool().Must(
			es.Term("moves.versionGroupDetails.moveLearnMethodName", "level-up"),
			es.Term("moves.versionGroupDetails.versionGroupName", "red-blue"),
		),
	)).Size(100)

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 100, len(result))
}

func (s *testSuite) Test_it_should_return_documents_pokedex_it_7() {
	// Given
	query := es.NewQuery(es.Nested("stats",
		es.Range("stats.baseStat").
			GreaterThanOrEqual(100),
	))

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 0, len(result))
}

func (s *testSuite) Test_it_should_return_documents_pokedex_it_8() {
	// Given
	query := es.NewQuery(
		es.Bool().Must(
			es.Range("height").
				GreaterThanOrEqual(10).
				LessThanOrEqual(20),
			es.Range("weight").
				GreaterThanOrEqual(100).
				LessThanOrEqual(200),
		),
	)

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 10, len(result))
}

func (s *testSuite) Test_it_should_return_documents_pokedex_it_9() {
	// Given
	query := es.NewQuery(
		es.Bool().Must(
			es.Nested("abilities", es.Term("abilities.name", "levitate")),
			es.Nested("types", es.Term("types.name", "ghost")),
		),
	)

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 3, len(result))
	for pokeID, pokemon := range result {
		switch pokeID {
		case "92_147":
			assert.Equal(s.T(), "gastly", pokemon.Name)
		case "93_148":
			assert.Equal(s.T(), "haunter", pokemon.Name)
		case "200_295":
			assert.Equal(s.T(), "misdreavus", pokemon.Name)
		default:
			s.T().FailNow()
		}
	}
}

func (s *testSuite) Test_it_should_return_documents_pokedex_it_10() {
	// Given
	query := es.NewQuery(
		es.Nested("moves.versionGroupDetails",
			es.Bool().Must(
				es.Term("moves.versionGroupDetails.moveLearnMethodName", "level-up"),
				es.Term("moves.versionGroupDetails.versionGroupName", "red-blue"),
				es.Range("moves.versionGroupDetails.levelLearnedAt").
					LessThanOrEqual(20),
			),
		),
	)

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 10, len(result))
}
