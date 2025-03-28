package es_test

import (
	"testing"

	ScriptLanguage "github.com/Trendyol/es-query-builder/es/enums/script-language"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_ScriptQuery_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.ScriptQuery)
}

func Test_ScriptQuery_should_create_json_with_script_field_inside_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.ScriptQuery(es.ScriptSource("src", ScriptLanguage.Painless)),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"script\":{\"script\":{\"lang\":\"painless\",\"source\":\"src\"}}}}", bodyJSON)
}

func Test_ScriptQuery_method_should_create_scriptQueryType(t *testing.T) {
	t.Parallel()
	// Given
	scriptQuery := es.ScriptQuery(es.ScriptSource("src", ScriptLanguage.Java))

	// Then
	assert.NotNil(t, scriptQuery)
	assert.IsTypeString(t, "es.scriptQueryType", scriptQuery)
}

func Test_ScriptQuery_should_have_Boost_method(t *testing.T) {
	t.Parallel()
	// Given
	scriptQuery := es.ScriptQuery(es.ScriptSource("src", ScriptLanguage.Mustache))

	// When Then
	assert.NotNil(t, scriptQuery.Boost)
}

func Test_ScriptQuery_Boost_should_create_json_with_boost_field_inside_scriptQuery(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.ScriptQuery(es.ScriptID("12345", ScriptLanguage.Expression)).
			Boost(3.14),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"script\":{\"boost\":3.14,\"script\":{\"id\":\"12345\",\"lang\":\"expression\"}}}}", bodyJSON)
}

func Test_ScriptQuery_should_have_Name_method(t *testing.T) {
	t.Parallel()
	// Given
	scriptQuery := es.ScriptQuery(es.ScriptSource("src", ScriptLanguage.Mustache))

	// When Then
	assert.NotNil(t, scriptQuery.Name)
}

func Test_ScriptQuery_Name_should_create_json_with__name_field_inside_scriptQuery(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.ScriptQuery(es.ScriptID("12345", ScriptLanguage.Painless)).
			Name("Cemil"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"script\":{\"_name\":\"Cemil\",\"script\":{\"id\":\"12345\",\"lang\":\"painless\"}}}}", bodyJSON)
}
