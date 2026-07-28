package tests_test

import (
	"encoding/json"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_composite_aggregation() {
	// Given
	query := es.NewQuery(es.MatchAll()).
		Size(0).
		Aggs(
			es.Agg("pages",
				es.CompositeAgg(
					es.CompositeTerms("by_name", "name.keyword"),
				).Size(5),
			),
		)

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	raw, exists := response.Aggregations["pages"]
	assert.True(s.T(), exists)

	var compositeAgg struct {
		AfterKey map[string]any `json:"after_key"`
		Buckets  []struct {
			Key      map[string]any `json:"key"`
			DocCount int            `json:"doc_count"`
		} `json:"buckets"`
	}
	assert.NoError(s.T(), json.Unmarshal(raw, &compositeAgg))
	assert.Equal(s.T(), 5, len(compositeAgg.Buckets))
	assert.NotEmpty(s.T(), compositeAgg.AfterKey)
	for _, bucket := range compositeAgg.Buckets {
		assert.NotEmpty(s.T(), bucket.Key["by_name"])
		assert.Greater(s.T(), bucket.DocCount, 0)
	}
}

func (s *testSuite) Test_it_should_paginate_composite_aggregation_with_after() {
	// Given
	firstPageQuery := es.NewQuery(es.MatchAll()).
		Size(0).
		Aggs(
			es.Agg("pages",
				es.CompositeAgg(
					es.CompositeTerms("by_name", "name.keyword"),
				).Size(5),
			),
		)

	firstPageResponse, err := s.PokedexElasticsearchRepository.Search(s.TestContext, firstPageQuery)
	assert.Nil(s.T(), err)

	var firstPage struct {
		AfterKey map[string]any `json:"after_key"`
		Buckets  []struct {
			Key map[string]any `json:"key"`
		} `json:"buckets"`
	}
	assert.NoError(s.T(), json.Unmarshal(firstPageResponse.Aggregations["pages"], &firstPage))
	assert.Equal(s.T(), 5, len(firstPage.Buckets))
	assert.NotEmpty(s.T(), firstPage.AfterKey)

	secondPageQuery := es.NewQuery(es.MatchAll()).
		Size(0).
		Aggs(
			es.Agg("pages",
				es.CompositeAgg(
					es.CompositeTerms("by_name", "name.keyword"),
				).Size(5).After(es.Object(firstPage.AfterKey)),
			),
		)

	// When
	secondPageResponse, err := s.PokedexElasticsearchRepository.Search(s.TestContext, secondPageQuery)

	// Then
	assert.Nil(s.T(), err)
	var secondPage struct {
		Buckets []struct {
			Key map[string]any `json:"key"`
		} `json:"buckets"`
	}
	assert.NoError(s.T(), json.Unmarshal(secondPageResponse.Aggregations["pages"], &secondPage))
	assert.Equal(s.T(), 5, len(secondPage.Buckets))

	firstKeys := make(map[string]bool)
	for _, bucket := range firstPage.Buckets {
		key, ok := bucket.Key["by_name"].(string)
		assert.True(s.T(), ok)
		firstKeys[key] = true
	}
	for _, bucket := range secondPage.Buckets {
		key, ok := bucket.Key["by_name"].(string)
		assert.True(s.T(), ok)
		assert.False(s.T(), firstKeys[key], "second page should not overlap first page keys")
	}
}

func (s *testSuite) Test_it_should_return_composite_aggregation_with_histogram_source() {
	// Given
	query := es.NewQuery(es.MatchAll()).
		Size(0).
		Aggs(
			es.Agg("pages",
				es.CompositeAgg(
					es.CompositeHistogram("by_height", "height", 10),
				).Size(5),
			),
		)

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	raw, exists := response.Aggregations["pages"]
	assert.True(s.T(), exists)

	var compositeAgg struct {
		Buckets []struct {
			Key      map[string]any `json:"key"`
			DocCount int            `json:"doc_count"`
		} `json:"buckets"`
	}
	assert.NoError(s.T(), json.Unmarshal(raw, &compositeAgg))
	assert.Greater(s.T(), len(compositeAgg.Buckets), 0)
	for _, bucket := range compositeAgg.Buckets {
		assert.Contains(s.T(), bucket.Key, "by_height")
		assert.Greater(s.T(), bucket.DocCount, 0)
	}
}
