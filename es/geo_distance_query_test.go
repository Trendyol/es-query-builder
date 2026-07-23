package es_test

import (
	"testing"

	DistanceType "github.com/Trendyol/es-query-builder/es/enums/distance-type"
	ValidationMethod "github.com/Trendyol/es-query-builder/es/enums/validation-method"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   GeoDistance   ////

func Test_GeoDistance_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.GeoDistance)
}

func Test_GeoDistance_should_create_json_with_geo_distance_field_inside_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.GeoDistance("pin.location", 40.0, -70.0, "12km"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"geo_distance\":{\"distance\":\"12km\",\"pin.location\":{\"lat\":40,\"lon\":-70}}}}", bodyJSON)
}

func Test_GeoDistance_method_should_create_geoDistanceType(t *testing.T) {
	t.Parallel()
	// Given
	g := es.GeoDistance("pin.location", 40.0, -70.0, "12km")

	// Then
	assert.NotNil(t, g)
	assert.IsTypeString(t, "es.geoDistanceType", g)
}

func Test_GeoDistance_should_have_DistanceType_method(t *testing.T) {
	t.Parallel()
	// Given
	g := es.GeoDistance("pin.location", 40.0, -70.0, "12km")

	// When Then
	assert.NotNil(t, g.DistanceType)
}

func Test_GeoDistance_DistanceType_should_create_json_with_distance_type_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.GeoDistance("pin.location", 40.0, -70.0, "12km").
			DistanceType(DistanceType.Arc),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"geo_distance\":{\"distance\":\"12km\",\"distance_type\":\"arc\",\"pin.location\":{\"lat\":40,\"lon\":-70}}}}", bodyJSON)
}

func Test_GeoDistance_should_have_ValidationMethod_method(t *testing.T) {
	t.Parallel()
	// Given
	g := es.GeoDistance("pin.location", 40.0, -70.0, "12km")

	// When Then
	assert.NotNil(t, g.ValidationMethod)
}

func Test_GeoDistance_ValidationMethod_should_create_json_with_validation_method_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.GeoDistance("pin.location", 40.0, -70.0, "12km").
			ValidationMethod(ValidationMethod.Strict),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"geo_distance\":{\"distance\":\"12km\",\"pin.location\":{\"lat\":40,\"lon\":-70},\"validation_method\":\"STRICT\"}}}", bodyJSON)
}

func Test_GeoDistance_should_have_IgnoreUnmapped_method(t *testing.T) {
	t.Parallel()
	// Given
	g := es.GeoDistance("pin.location", 40.0, -70.0, "12km")

	// When Then
	assert.NotNil(t, g.IgnoreUnmapped)
}

func Test_GeoDistance_IgnoreUnmapped_should_create_json_with_ignore_unmapped_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.GeoDistance("pin.location", 40.0, -70.0, "12km").
			IgnoreUnmapped(true),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"geo_distance\":{\"distance\":\"12km\",\"ignore_unmapped\":true,\"pin.location\":{\"lat\":40,\"lon\":-70}}}}", bodyJSON)
}

func Test_GeoDistance_should_have_Boost_method(t *testing.T) {
	t.Parallel()
	// Given
	g := es.GeoDistance("pin.location", 40.0, -70.0, "12km")

	// When Then
	assert.NotNil(t, g.Boost)
}

func Test_GeoDistance_Boost_should_create_json_with_boost_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.GeoDistance("pin.location", 40.0, -70.0, "12km").
			Boost(1.5),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"geo_distance\":{\"boost\":1.5,\"distance\":\"12km\",\"pin.location\":{\"lat\":40,\"lon\":-70}}}}", bodyJSON)
}

func Test_GeoDistance_should_have_Name_method(t *testing.T) {
	t.Parallel()
	// Given
	g := es.GeoDistance("pin.location", 40.0, -70.0, "12km")

	// When Then
	assert.NotNil(t, g.Name)
}

func Test_GeoDistance_Name_should_create_json_with_name_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.GeoDistance("pin.location", 40.0, -70.0, "12km").
			Name("nearby"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"geo_distance\":{\"_name\":\"nearby\",\"distance\":\"12km\",\"pin.location\":{\"lat\":40,\"lon\":-70}}}}", bodyJSON)
}

func Test_GeoDistance_should_create_json_with_all_fields(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.GeoDistance("pin.location", 40.0, -70.0, "12km").
			DistanceType(DistanceType.Plane).
			ValidationMethod(ValidationMethod.Coerce).
			IgnoreUnmapped(true).
			Boost(2.0).
			Name("nearby"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"geo_distance\":{\"_name\":\"nearby\",\"boost\":2,\"distance\":\"12km\",\"distance_type\":\"plane\",\"ignore_unmapped\":true,\"pin.location\":{\"lat\":40,\"lon\":-70},\"validation_method\":\"COERCE\"}}}", bodyJSON)
}

func Test_GeoDistance_Plane_distance_type_should_create_correct_json(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.GeoDistance("location", 41.0, 29.0, "5km").
			DistanceType(DistanceType.Plane),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"geo_distance\":{\"distance\":\"5km\",\"distance_type\":\"plane\",\"location\":{\"lat\":41,\"lon\":29}}}}", bodyJSON)
}
