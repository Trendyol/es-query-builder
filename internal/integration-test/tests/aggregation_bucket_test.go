package tests_test

import (
	"encoding/json"

	Order "github.com/Trendyol/es-query-builder/es/enums/sort/order"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_terms_aggregation() {
	// Given
	query := es.NewQuery(es.MatchAll()).
		Size(0).
		Aggs(es.Agg("by_is_default", es.TermsAgg("isDefault").Size(2)))

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	raw, exists := response.Aggregations["by_is_default"]
	assert.True(s.T(), exists)

	var termsAgg struct {
		Buckets []struct {
			KeyAsString string `json:"key_as_string"`
			DocCount    int    `json:"doc_count"`
		} `json:"buckets"`
	}
	assert.NoError(s.T(), json.Unmarshal(raw, &termsAgg))
	assert.Greater(s.T(), len(termsAgg.Buckets), 0)
	for _, bucket := range termsAgg.Buckets {
		assert.Greater(s.T(), bucket.DocCount, 0)
	}
}

func (s *testSuite) Test_it_should_return_histogram_aggregation() {
	// Given
	query := es.NewQuery(es.MatchAll()).
		Size(0).
		Aggs(es.Agg("height_histogram", es.HistogramAgg("height", 10).MinDocCount(1)))

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	raw, exists := response.Aggregations["height_histogram"]
	assert.True(s.T(), exists)

	var histogramAgg struct {
		Buckets []struct {
			Key      float64 `json:"key"`
			DocCount int     `json:"doc_count"`
		} `json:"buckets"`
	}
	assert.NoError(s.T(), json.Unmarshal(raw, &histogramAgg))
	assert.Greater(s.T(), len(histogramAgg.Buckets), 0)
	for _, bucket := range histogramAgg.Buckets {
		assert.Greater(s.T(), bucket.DocCount, 0)
	}
}

func (s *testSuite) Test_it_should_return_range_aggregation() {
	// Given
	query := es.NewQuery(es.MatchAll()).
		Size(0).
		Aggs(
			es.Agg("weight_ranges",
				es.RangeAgg("weight").
					Range(es.RangeEntry().To(50)).
					Range(es.RangeEntry().From(50).To(200)).
					Range(es.RangeEntry().From(200)),
			),
		)

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	raw, exists := response.Aggregations["weight_ranges"]
	assert.True(s.T(), exists)

	var rangeAgg struct {
		Buckets []struct {
			Key      string `json:"key"`
			DocCount int    `json:"doc_count"`
		} `json:"buckets"`
	}
	assert.NoError(s.T(), json.Unmarshal(raw, &rangeAgg))
	assert.Equal(s.T(), 3, len(rangeAgg.Buckets))
	totalDocs := 0
	for _, bucket := range rangeAgg.Buckets {
		totalDocs += bucket.DocCount
	}
	assert.Greater(s.T(), totalDocs, 0)
}

func (s *testSuite) Test_it_should_return_filter_aggregation_with_sub_avg() {
	// Given
	query := es.NewQuery(es.MatchAll()).
		Size(0).
		Aggs(
			es.Agg("default_pokemon",
				es.FilterAgg(es.Term("isDefault", true)).
					Aggs(es.Agg("avg_height", es.AvgAgg("height"))),
			),
		)

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	raw, exists := response.Aggregations["default_pokemon"]
	assert.True(s.T(), exists)

	var filterAgg struct {
		DocCount  int `json:"doc_count"`
		AvgHeight struct {
			Value float64 `json:"value"`
		} `json:"avg_height"`
	}
	assert.NoError(s.T(), json.Unmarshal(raw, &filterAgg))
	assert.Greater(s.T(), filterAgg.DocCount, 0)
	assert.Greater(s.T(), filterAgg.AvgHeight.Value, float64(0))
}

func (s *testSuite) Test_it_should_return_filters_aggregation() {
	// Given
	query := es.NewQuery(es.MatchAll()).
		Size(0).
		Aggs(
			es.Agg("by_name",
				es.FiltersAgg().
					Filter("pikachu", es.Term("name.keyword", "pikachu")).
					Filter("bulbasaur", es.Term("name.keyword", "bulbasaur")),
			),
		)

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	raw, exists := response.Aggregations["by_name"]
	assert.True(s.T(), exists)

	var filtersAgg struct {
		Buckets map[string]struct {
			DocCount int `json:"doc_count"`
		} `json:"buckets"`
	}
	assert.NoError(s.T(), json.Unmarshal(raw, &filtersAgg))
	assert.Equal(s.T(), 1, filtersAgg.Buckets["pikachu"].DocCount)
	assert.Equal(s.T(), 1, filtersAgg.Buckets["bulbasaur"].DocCount)
}

func (s *testSuite) Test_it_should_return_multi_terms_aggregation() {
	// Given
	query := es.NewQuery(es.MatchAll()).
		Size(0).
		Aggs(
			es.Agg("by_name_and_default",
				es.MultiTermsAgg(
					es.TermAgg("name.keyword"),
					es.TermAgg("isDefault"),
				).Size(10),
			),
		)

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	raw, exists := response.Aggregations["by_name_and_default"]
	assert.True(s.T(), exists)

	var multiTermsAgg struct {
		Buckets []struct {
			Key      []any `json:"key"`
			DocCount int   `json:"doc_count"`
		} `json:"buckets"`
	}
	assert.NoError(s.T(), json.Unmarshal(raw, &multiTermsAgg))
	assert.Greater(s.T(), len(multiTermsAgg.Buckets), 0)
	for _, bucket := range multiTermsAgg.Buckets {
		assert.Equal(s.T(), 2, len(bucket.Key))
		assert.Greater(s.T(), bucket.DocCount, 0)
	}
}

func (s *testSuite) Test_it_should_return_top_hits_aggregation_at_root() {
	// Given
	query := es.NewQuery(es.MatchAll()).
		Size(0).
		Aggs(
			es.Agg("top_pokemon",
				es.TopHitsAgg().
					Size(3).
					Sort(es.Sort("id").Order(Order.Asc)).
					SourceIncludes("name", "id"),
			),
		)

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	raw, exists := response.Aggregations["top_pokemon"]
	assert.True(s.T(), exists)

	var topHits struct {
		Hits struct {
			Hits []struct {
				Source struct {
					Name string `json:"name"`
					ID   int    `json:"id"`
				} `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}
	assert.NoError(s.T(), json.Unmarshal(raw, &topHits))
	assert.Equal(s.T(), 3, len(topHits.Hits.Hits))
	assert.NotEmpty(s.T(), topHits.Hits.Hits[0].Source.Name)
}
