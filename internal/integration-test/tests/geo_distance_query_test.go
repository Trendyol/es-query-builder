package tests_test

import (
	"integration-tests/model"
	"integration-tests/tests"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_geo_distance_query() {
	// Given — Istanbul center (~41.01, 28.97), nearby Kadikoy, far Ankara
	nearby := model.FooDocument{
		ID:  "10",
		Foo: "kadikoy",
		Location: &model.GeoPoint{
			Lat: 40.99,
			Lon: 29.03,
		},
	}
	far := model.FooDocument{
		ID:  "20",
		Foo: "ankara",
		Location: &model.GeoPoint{
			Lat: 39.93,
			Lon: 32.85,
		},
	}

	s.FooElasticsearchRepository.BulkInsert(s.TestContext, nearby, far)
	tests.WaitExists(s.TestContext, s.FooElasticsearchRepository, nearby.ID)
	tests.WaitExists(s.TestContext, s.FooElasticsearchRepository, far.ID)

	query := es.NewQuery(
		es.GeoDistance("location", 41.01, 28.97, "20km"),
	)

	// When
	result, err := s.FooElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 1, len(result))
	assert.Equal(s.T(), "kadikoy", result["10"].Foo)

	s.FooElasticsearchRepository.BulkDelete(s.TestContext, nearby.ID, far.ID)
}
