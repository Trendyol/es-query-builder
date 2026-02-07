package es_test

import (
	"testing"

	Modifier "github.com/Trendyol/es-query-builder/es/enums/modifier"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   FieldValueFactor   ////

func Test_FieldValueFactor_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.FieldValueFactor)
}

func Test_FieldValueFactor_method_should_create_fieldValueFactorType(t *testing.T) {
	t.Parallel()
	// Given
	fvf := es.FieldValueFactor("likes")

	// Then
	assert.NotNil(t, fvf)
	assert.IsTypeString(t, "es.fieldValueFactorType", fvf)
}

func Test_FieldValueFactor_should_create_json_with_field(t *testing.T) {
	t.Parallel()
	// Given
	fvf := es.FieldValueFactor("likes")

	// When Then
	assert.NotNil(t, fvf)
	bodyJSON := assert.MarshalWithoutError(t, fvf)
	assert.Equal(t, "{\"field\":\"likes\"}", bodyJSON)
}

////   Factor   ////

func Test_FieldValueFactor_should_have_Factor_method(t *testing.T) {
	t.Parallel()
	// Given
	fvf := es.FieldValueFactor("likes")

	// When Then
	assert.NotNil(t, fvf.Factor)
}

func Test_FieldValueFactor_Factor_should_create_json_with_factor_field(t *testing.T) {
	t.Parallel()
	// Given
	fvf := es.FieldValueFactor("likes").Factor(1.2)

	// When Then
	assert.NotNil(t, fvf)
	bodyJSON := assert.MarshalWithoutError(t, fvf)
	assert.Equal(t, "{\"factor\":1.2,\"field\":\"likes\"}", bodyJSON)
}

////   Modifier   ////

func Test_FieldValueFactor_should_have_Modifier_method(t *testing.T) {
	t.Parallel()
	// Given
	fvf := es.FieldValueFactor("likes")

	// When Then
	assert.NotNil(t, fvf.Modifier)
}

func Test_FieldValueFactor_Modifier_should_create_json_with_modifier_field(t *testing.T) {
	t.Parallel()
	// Given
	fvf := es.FieldValueFactor("likes").Modifier(Modifier.Log1p)

	// When Then
	assert.NotNil(t, fvf)
	bodyJSON := assert.MarshalWithoutError(t, fvf)
	assert.Equal(t, "{\"field\":\"likes\",\"modifier\":\"log1p\"}", bodyJSON)
}

////   Missing   ////

func Test_FieldValueFactor_should_have_Missing_method(t *testing.T) {
	t.Parallel()
	// Given
	fvf := es.FieldValueFactor("likes")

	// When Then
	assert.NotNil(t, fvf.Missing)
}

func Test_FieldValueFactor_Missing_should_create_json_with_missing_field(t *testing.T) {
	t.Parallel()
	// Given
	fvf := es.FieldValueFactor("likes").Missing(1)

	// When Then
	assert.NotNil(t, fvf)
	bodyJSON := assert.MarshalWithoutError(t, fvf)
	assert.Equal(t, "{\"field\":\"likes\",\"missing\":1}", bodyJSON)
}

////   Combined   ////

func Test_FieldValueFactor_should_create_json_with_all_parameters(t *testing.T) {
	t.Parallel()
	// Given
	fvf := es.FieldValueFactor("likes").
		Factor(1.2).
		Modifier(Modifier.Log1p).
		Missing(1)

	// When Then
	assert.NotNil(t, fvf)
	bodyJSON := assert.MarshalWithoutError(t, fvf)
	assert.Equal(t, "{\"factor\":1.2,\"field\":\"likes\",\"missing\":1,\"modifier\":\"log1p\"}", bodyJSON)
}
