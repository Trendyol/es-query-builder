package tests_test

import (
	"integration-tests/model"
	"integration-tests/tests"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_geo_bounding_box_query() {
	// Given — box around Istanbul; one point inside, one outside
	inside := model.FooDocument{
		ID:  "10",
		Foo: "taksim",
		Location: &model.GeoPoint{
			Lat: 41.037,
			Lon: 28.985,
		},
	}
	outside := model.FooDocument{
		ID:  "20",
		Foo: "izmir",
		Location: &model.GeoPoint{
			Lat: 38.42,
			Lon: 27.14,
		},
	}

	s.FooElasticsearchRepository.BulkInsert(s.TestContext, inside, outside)
	tests.WaitExists(s.TestContext, s.FooElasticsearchRepository, inside.ID)
	tests.WaitExists(s.TestContext, s.FooElasticsearchRepository, outside.ID)

	// top-left / bottom-right covering Istanbul metro area
	query := es.NewQuery(
		es.GeoBoundingBox("location", 41.20, 28.70, 40.80, 29.20),
	)

	// When
	result, err := s.FooElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 1, len(result))
	assert.Equal(s.T(), "taksim", result["10"].Foo)

	s.FooElasticsearchRepository.BulkDelete(s.TestContext, inside.ID, outside.ID)
}
