package tests_test

import (
	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_range_query_with_gte() {
	// Given
	query := es.NewQuery(
		es.Range("baseExperience").GreaterThanOrEqual(300),
	).Size(50)

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Greater(s.T(), len(result), 0)

	// Tüm sonuçlar baseExperience >= 300 olmalı
	for _, pokemon := range result {
		assert.GreaterOrEqual(s.T(), pokemon.BaseExperience, uint16(300))
	}
}

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_range_query_with_lte() {
	// Given
	query := es.NewQuery(
		es.Range("height").LessThanOrEqual(5),
	).Size(50)

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Greater(s.T(), len(result), 0)

	// Tüm sonuçlar height <= 5 olmalı
	for _, pokemon := range result {
		assert.LessOrEqual(s.T(), pokemon.Height, uint16(5))
	}
}

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_range_query_with_gt_and_lt() {
	// Given
	query := es.NewQuery(
		es.Range("weight").
			GreaterThan(100).
			LessThan(500),
	).Size(50)

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Greater(s.T(), len(result), 0)

	// Tüm sonuçlar 100 < weight < 500 olmalı
	for _, pokemon := range result {
		assert.Greater(s.T(), pokemon.Weight, uint16(100))
		assert.Less(s.T(), pokemon.Weight, uint16(500))
	}
}

func (s *testSuite) Test_it_should_return_starter_pokemons_by_id_range() {
	// Given - İlk 10 pokemon (starter bölgesi)
	query := es.NewQuery(
		es.Range("id").
			GreaterThanOrEqual(1).
			LessThanOrEqual(10),
	)

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 10, len(result))

	// ID'ler 1-10 arasında olmalı
	for _, pokemon := range result {
		assert.GreaterOrEqual(s.T(), pokemon.Id, uint16(1))
		assert.LessOrEqual(s.T(), pokemon.Id, uint16(10))
	}
}
