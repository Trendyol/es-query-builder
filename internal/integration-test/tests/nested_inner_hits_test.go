package tests_test

import (
	"encoding/json"

	ScoreMode "github.com/Trendyol/es-query-builder/es/enums/score-mode"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_inner_hits_for_nested_query() {
	// Given
	query := es.NewQuery(
		es.Nested("abilities", es.Term("abilities.name", "blaze")).
			InnerHits(
				es.InnerHits().
					Name("matched_abilities").
					Size(5),
			),
	)

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), response.Hits)
	assert.Greater(s.T(), len(response.Hits.Hits), 0)

	hit := response.Hits.Hits[0]
	assert.NotEmpty(s.T(), hit.InnerHits)

	raw, ok := hit.InnerHits["matched_abilities"]
	assert.True(s.T(), ok)

	var innerHits struct {
		Hits struct {
			Hits []struct {
				Source map[string]any `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}
	assert.NoError(s.T(), json.Unmarshal(raw, &innerHits))
	assert.Greater(s.T(), len(innerHits.Hits.Hits), 0)
	assert.Equal(s.T(), "blaze", innerHits.Hits.Hits[0].Source["name"])
}

func (s *testSuite) Test_it_should_apply_score_mode_on_nested_query() {
	// Given
	query := es.NewQuery(
		es.Nested("abilities", es.Term("abilities.name", "blaze")).
			ScoreMode(ScoreMode.Avg),
	)

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), response.Hits)
	assert.Greater(s.T(), len(response.Hits.Hits), 0)
	assert.Greater(s.T(), response.Hits.Hits[0].Score, float32(0))
}
