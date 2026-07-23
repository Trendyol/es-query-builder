package tests_test

import (
	"encoding/json"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_value_count_aggregation() {
	// Given
	query := es.NewQuery(es.MatchAll()).
		Size(0).
		Aggs(
			es.Agg("name_count", es.ValueCountAgg("name.keyword")),
		)

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), response.Aggregations)

	raw, exists := response.Aggregations["name_count"]
	assert.True(s.T(), exists)

	var valueCount struct {
		Value float64 `json:"value"`
	}
	assert.NoError(s.T(), json.Unmarshal(raw, &valueCount))
	assert.Greater(s.T(), valueCount.Value, float64(0))
}
