package tests_test

import (
	"encoding/json"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_avg_aggregation() {
	// Given
	query := es.NewQuery(es.MatchAll()).
		Size(0).
		Aggs(es.Agg("avg_height", es.AvgAgg("height")))

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	raw, exists := response.Aggregations["avg_height"]
	assert.True(s.T(), exists)

	var avgAgg struct {
		Value float64 `json:"value"`
	}
	assert.NoError(s.T(), json.Unmarshal(raw, &avgAgg))
	assert.Greater(s.T(), avgAgg.Value, float64(0))
}

func (s *testSuite) Test_it_should_return_min_aggregation() {
	// Given
	query := es.NewQuery(es.MatchAll()).
		Size(0).
		Aggs(es.Agg("min_weight", es.MinAgg("weight")))

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	raw, exists := response.Aggregations["min_weight"]
	assert.True(s.T(), exists)

	var minAgg struct {
		Value float64 `json:"value"`
	}
	assert.NoError(s.T(), json.Unmarshal(raw, &minAgg))
	assert.GreaterOrEqual(s.T(), minAgg.Value, float64(0))
}

func (s *testSuite) Test_it_should_return_max_aggregation() {
	// Given
	query := es.NewQuery(es.MatchAll()).
		Size(0).
		Aggs(es.Agg("max_base_experience", es.MaxAgg("baseExperience")))

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	raw, exists := response.Aggregations["max_base_experience"]
	assert.True(s.T(), exists)

	var maxAgg struct {
		Value float64 `json:"value"`
	}
	assert.NoError(s.T(), json.Unmarshal(raw, &maxAgg))
	assert.Greater(s.T(), maxAgg.Value, float64(0))
}

func (s *testSuite) Test_it_should_return_sum_aggregation() {
	// Given
	query := es.NewQuery(es.MatchAll()).
		Size(0).
		Aggs(es.Agg("sum_height", es.SumAgg("height")))

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	raw, exists := response.Aggregations["sum_height"]
	assert.True(s.T(), exists)

	var sumAgg struct {
		Value float64 `json:"value"`
	}
	assert.NoError(s.T(), json.Unmarshal(raw, &sumAgg))
	assert.Greater(s.T(), sumAgg.Value, float64(0))
}

func (s *testSuite) Test_it_should_return_stats_aggregation() {
	// Given
	query := es.NewQuery(es.MatchAll()).
		Size(0).
		Aggs(es.Agg("height_stats", es.StatsAgg("height")))

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	raw, exists := response.Aggregations["height_stats"]
	assert.True(s.T(), exists)

	var statsAgg struct {
		Count float64 `json:"count"`
		Min   float64 `json:"min"`
		Max   float64 `json:"max"`
		Avg   float64 `json:"avg"`
		Sum   float64 `json:"sum"`
	}
	assert.NoError(s.T(), json.Unmarshal(raw, &statsAgg))
	assert.Greater(s.T(), statsAgg.Count, float64(0))
	assert.Greater(s.T(), statsAgg.Max, statsAgg.Min)
	assert.Greater(s.T(), statsAgg.Sum, float64(0))
}

func (s *testSuite) Test_it_should_return_extended_stats_aggregation() {
	// Given
	query := es.NewQuery(es.MatchAll()).
		Size(0).
		Aggs(es.Agg("weight_extended_stats", es.ExtendedStatsAgg("weight")))

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	raw, exists := response.Aggregations["weight_extended_stats"]
	assert.True(s.T(), exists)

	var extendedStats struct {
		Count        float64 `json:"count"`
		Min          float64 `json:"min"`
		Max          float64 `json:"max"`
		Avg          float64 `json:"avg"`
		Sum          float64 `json:"sum"`
		StdDeviation float64 `json:"std_deviation"`
		Variance     float64 `json:"variance"`
	}
	assert.NoError(s.T(), json.Unmarshal(raw, &extendedStats))
	assert.Greater(s.T(), extendedStats.Count, float64(0))
	assert.GreaterOrEqual(s.T(), extendedStats.StdDeviation, float64(0))
	assert.GreaterOrEqual(s.T(), extendedStats.Variance, float64(0))
}

func (s *testSuite) Test_it_should_return_cardinality_aggregation() {
	// Given
	query := es.NewQuery(es.MatchAll()).
		Size(0).
		Aggs(es.Agg("unique_names", es.CardinalityAgg("name.keyword")))

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	raw, exists := response.Aggregations["unique_names"]
	assert.True(s.T(), exists)

	var cardinality struct {
		Value float64 `json:"value"`
	}
	assert.NoError(s.T(), json.Unmarshal(raw, &cardinality))
	assert.Greater(s.T(), cardinality.Value, float64(0))
}
