package es_test

import (
	"testing"

	ValidationMethod "github.com/Trendyol/es-query-builder/es/enums/validation-method"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   GeoBoundingBox   ////

func Test_GeoBoundingBox_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.GeoBoundingBox)
}

func Test_GeoBoundingBox_should_create_json_with_geo_bounding_box_field_inside_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.GeoBoundingBox("pin.location", 40.73, -74.1, 40.01, -71.12),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"geo_bounding_box\":{\"pin.location\":{\"bottom_right\":{\"lat\":40.01,\"lon\":-71.12},\"top_left\":{\"lat\":40.73,\"lon\":-74.1}}}}}", bodyJSON)
}

func Test_GeoBoundingBox_method_should_create_geoBoundingBoxType(t *testing.T) {
	t.Parallel()
	// Given
	g := es.GeoBoundingBox("pin.location", 40.73, -74.1, 40.01, -71.12)

	// Then
	assert.NotNil(t, g)
	assert.IsTypeString(t, "es.geoBoundingBoxType", g)
}

func Test_GeoBoundingBox_should_have_ValidationMethod_method(t *testing.T) {
	t.Parallel()
	// Given
	g := es.GeoBoundingBox("pin.location", 40.73, -74.1, 40.01, -71.12)

	// When Then
	assert.NotNil(t, g.ValidationMethod)
}

func Test_GeoBoundingBox_ValidationMethod_should_create_json_with_validation_method_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.GeoBoundingBox("pin.location", 40.73, -74.1, 40.01, -71.12).
			ValidationMethod(ValidationMethod.Strict),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"geo_bounding_box\":{\"pin.location\":{\"bottom_right\":{\"lat\":40.01,\"lon\":-71.12},\"top_left\":{\"lat\":40.73,\"lon\":-74.1}},\"validation_method\":\"STRICT\"}}}", bodyJSON)
}

func Test_GeoBoundingBox_should_have_IgnoreUnmapped_method(t *testing.T) {
	t.Parallel()
	// Given
	g := es.GeoBoundingBox("pin.location", 40.73, -74.1, 40.01, -71.12)

	// When Then
	assert.NotNil(t, g.IgnoreUnmapped)
}

func Test_GeoBoundingBox_IgnoreUnmapped_should_create_json_with_ignore_unmapped_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.GeoBoundingBox("pin.location", 40.73, -74.1, 40.01, -71.12).
			IgnoreUnmapped(true),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"geo_bounding_box\":{\"ignore_unmapped\":true,\"pin.location\":{\"bottom_right\":{\"lat\":40.01,\"lon\":-71.12},\"top_left\":{\"lat\":40.73,\"lon\":-74.1}}}}}", bodyJSON)
}

func Test_GeoBoundingBox_should_have_Boost_method(t *testing.T) {
	t.Parallel()
	// Given
	g := es.GeoBoundingBox("pin.location", 40.73, -74.1, 40.01, -71.12)

	// When Then
	assert.NotNil(t, g.Boost)
}

func Test_GeoBoundingBox_Boost_should_create_json_with_boost_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.GeoBoundingBox("pin.location", 40.73, -74.1, 40.01, -71.12).
			Boost(1.5),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"geo_bounding_box\":{\"boost\":1.5,\"pin.location\":{\"bottom_right\":{\"lat\":40.01,\"lon\":-71.12},\"top_left\":{\"lat\":40.73,\"lon\":-74.1}}}}}", bodyJSON)
}

func Test_GeoBoundingBox_should_have_Name_method(t *testing.T) {
	t.Parallel()
	// Given
	g := es.GeoBoundingBox("pin.location", 40.73, -74.1, 40.01, -71.12)

	// When Then
	assert.NotNil(t, g.Name)
}

func Test_GeoBoundingBox_Name_should_create_json_with_name_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.GeoBoundingBox("pin.location", 40.73, -74.1, 40.01, -71.12).
			Name("map_viewport"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"geo_bounding_box\":{\"_name\":\"map_viewport\",\"pin.location\":{\"bottom_right\":{\"lat\":40.01,\"lon\":-71.12},\"top_left\":{\"lat\":40.73,\"lon\":-74.1}}}}}", bodyJSON)
}

func Test_GeoBoundingBox_should_create_json_with_all_fields(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.GeoBoundingBox("pin.location", 40.73, -74.1, 40.01, -71.12).
			ValidationMethod(ValidationMethod.Coerce).
			IgnoreUnmapped(true).
			Boost(2.0).
			Name("map_viewport"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"geo_bounding_box\":{\"_name\":\"map_viewport\",\"boost\":2,\"ignore_unmapped\":true,\"pin.location\":{\"bottom_right\":{\"lat\":40.01,\"lon\":-71.12},\"top_left\":{\"lat\":40.73,\"lon\":-74.1}},\"validation_method\":\"COERCE\"}}}", bodyJSON)
}
