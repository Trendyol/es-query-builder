package es_test

import (
	"testing"

	MultiValuesMode "github.com/Trendyol/es-query-builder/es/enums/multi-values-mode"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   Decay   ////

func Test_Decay_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.Decay)
}

func Test_Decay_method_should_create_decayFunctionType(t *testing.T) {
	t.Parallel()
	// Given
	d := es.Decay("date")

	// Then
	assert.NotNil(t, d)
	assert.IsTypeString(t, "es.decayFunctionType", d)
}

func Test_Decay_should_create_json_with_field(t *testing.T) {
	t.Parallel()
	// Given
	d := es.Decay("date")

	// When Then
	assert.NotNil(t, d)
	bodyJSON := assert.MarshalWithoutError(t, d)
	assert.Equal(t, "{\"date\":{}}", bodyJSON)
}

////   Origin   ////

func Test_Decay_should_have_Origin_method(t *testing.T) {
	t.Parallel()
	// Given
	d := es.Decay("date")

	// When Then
	assert.NotNil(t, d.Origin)
}

func Test_Decay_Origin_should_create_json_with_origin_field(t *testing.T) {
	t.Parallel()
	// Given
	d := es.Decay("date").Origin("now")

	// When Then
	assert.NotNil(t, d)
	bodyJSON := assert.MarshalWithoutError(t, d)
	assert.Equal(t, "{\"date\":{\"origin\":\"now\"}}", bodyJSON)
}

////   Scale   ////

func Test_Decay_should_have_Scale_method(t *testing.T) {
	t.Parallel()
	// Given
	d := es.Decay("date")

	// When Then
	assert.NotNil(t, d.Scale)
}

func Test_Decay_Scale_should_create_json_with_scale_field(t *testing.T) {
	t.Parallel()
	// Given
	d := es.Decay("date").Origin("now").Scale("10d")

	// When Then
	assert.NotNil(t, d)
	bodyJSON := assert.MarshalWithoutError(t, d)
	assert.Equal(t, "{\"date\":{\"origin\":\"now\",\"scale\":\"10d\"}}", bodyJSON)
}

////   Offset   ////

func Test_Decay_should_have_Offset_method(t *testing.T) {
	t.Parallel()
	// Given
	d := es.Decay("date")

	// When Then
	assert.NotNil(t, d.Offset)
}

func Test_Decay_Offset_should_create_json_with_offset_field(t *testing.T) {
	t.Parallel()
	// Given
	d := es.Decay("date").Origin("now").Scale("10d").Offset("5d")

	// When Then
	assert.NotNil(t, d)
	bodyJSON := assert.MarshalWithoutError(t, d)
	assert.Equal(t, "{\"date\":{\"offset\":\"5d\",\"origin\":\"now\",\"scale\":\"10d\"}}", bodyJSON)
}

////   DecayValue   ////

func Test_Decay_should_have_DecayValue_method(t *testing.T) {
	t.Parallel()
	// Given
	d := es.Decay("date")

	// When Then
	assert.NotNil(t, d.DecayValue)
}

func Test_Decay_DecayValue_should_create_json_with_decay_field(t *testing.T) {
	t.Parallel()
	// Given
	d := es.Decay("date").Origin("now").Scale("10d").DecayValue(0.5)

	// When Then
	assert.NotNil(t, d)
	bodyJSON := assert.MarshalWithoutError(t, d)
	assert.Equal(t, "{\"date\":{\"decay\":0.5,\"origin\":\"now\",\"scale\":\"10d\"}}", bodyJSON)
}

////   MultiValueMode   ////

func Test_Decay_should_have_MultiValueMode_method(t *testing.T) {
	t.Parallel()
	// Given
	d := es.Decay("date")

	// When Then
	assert.NotNil(t, d.MultiValueMode)
}

func Test_Decay_MultiValueMode_should_create_json_with_multi_value_mode_field(t *testing.T) {
	t.Parallel()
	// Given
	d := es.Decay("location").Origin("0,0").Scale("5km").MultiValueMode(MultiValuesMode.Min)

	// When Then
	assert.NotNil(t, d)
	bodyJSON := assert.MarshalWithoutError(t, d)
	assert.Equal(t, "{\"location\":{\"origin\":\"0,0\",\"scale\":\"5km\"},\"multi_value_mode\":\"min\"}", bodyJSON)
}

////   Combined   ////

func Test_Decay_should_create_json_with_all_parameters(t *testing.T) {
	t.Parallel()
	// Given
	d := es.Decay("date").
		Origin("now").
		Scale("10d").
		Offset("5d").
		DecayValue(0.5)

	// When Then
	assert.NotNil(t, d)
	bodyJSON := assert.MarshalWithoutError(t, d)
	assert.Equal(t, "{\"date\":{\"decay\":0.5,\"offset\":\"5d\",\"origin\":\"now\",\"scale\":\"10d\"}}", bodyJSON)
}

func Test_DecayFunction_with_filter_and_weight_in_function_score(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.FunctionScore(es.MatchAll()).
			Functions(
				es.DecayFunction("gauss", es.Decay("date").Origin("now").Scale("10d")).
					Filter(es.Term("status", "published")).
					Weight(2),
			),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"function_score\":{\"functions\":[{\"filter\":{\"term\":{\"status\":{\"value\":\"published\"}}},\"gauss\":{\"date\":{\"origin\":\"now\",\"scale\":\"10d\"}},\"weight\":2}],\"query\":{\"match_all\":{}}}}}", bodyJSON)
}
