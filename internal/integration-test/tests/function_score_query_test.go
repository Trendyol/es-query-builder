package tests_test

import (
	scriptlanguage "github.com/Trendyol/es-query-builder/es/enums/script-language"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_score_documents_with_weight_function() {
	// Given
	query := es.NewQuery(
		es.FunctionScore(es.MatchAll()).
			Functions(es.WeightFunction(2.0)),
	).Size(5)

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), response.Hits)
	assert.Greater(s.T(), len(response.Hits.Hits), 0)
	assert.Greater(s.T(), response.Hits.Hits[0].Score, float32(0))
}

func (s *testSuite) Test_it_should_score_documents_with_script_score_function() {
	// Given
	query := es.NewQuery(
		es.FunctionScore(es.MatchAll()).
			Functions(
				es.ScriptScoreFunction(
					es.ScriptSource("doc['height'].value", scriptlanguage.Painless),
				),
			),
	).Size(5)

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), response.Hits)
	assert.Greater(s.T(), len(response.Hits.Hits), 0)
	assert.Greater(s.T(), response.Hits.Hits[0].Score, float32(0))
}

func (s *testSuite) Test_it_should_score_documents_with_field_value_factor_function() {
	// Given
	query := es.NewQuery(
		es.FunctionScore(es.MatchAll()).
			Functions(
				es.FieldValueFactorFunction(
					es.FieldValueFactor("height").Factor(1.2).Missing(1),
				),
			),
	).Size(5)

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), response.Hits)
	assert.Greater(s.T(), len(response.Hits.Hits), 0)
	assert.Greater(s.T(), response.Hits.Hits[0].Score, float32(0))
}

func (s *testSuite) Test_it_should_score_documents_with_random_score_function() {
	// Given
	query := es.NewQuery(
		es.FunctionScore(es.MatchAll()).
			Functions(es.RandomScoreFunction().Seed(42).Field("id")),
	).Size(5)

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), response.Hits)
	assert.Greater(s.T(), len(response.Hits.Hits), 0)
	assert.Greater(s.T(), response.Hits.Hits[0].Score, float32(0))
}

func (s *testSuite) Test_it_should_score_documents_with_decay_function() {
	// Given
	query := es.NewQuery(
		es.FunctionScore(es.MatchAll()).
			Functions(
				es.DecayFunction("gauss",
					es.Decay("height").Origin(10).Scale(20).Offset(0).DecayValue(0.5),
				),
			),
	).Size(5)

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), response.Hits)
	assert.Greater(s.T(), len(response.Hits.Hits), 0)
	assert.Greater(s.T(), response.Hits.Hits[0].Score, float32(0))
}
