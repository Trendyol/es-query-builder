package es_test

import (
	"testing"

	ScriptLanguage "github.com/Trendyol/es-query-builder/es/enums/script-language"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   ScriptID   ////

func Test_ScriptId_should_script_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.ScriptID)
}

func Test_ScriptId_should_create_json_with_id_and_lang_field_inside_script(t *testing.T) {
	t.Parallel()
	// Given
	script := es.ScriptID("custom_match_script", ScriptLanguage.Painless)

	// When Then
	assert.NotNil(t, script)
	scriptJSON := assert.MarshalWithoutError(t, script)
	assert.Equal(t, "{\"id\":\"custom_match_script\",\"lang\":\"painless\"}", scriptJSON)
}

////   ScriptSource   ////

func Test_ScriptSource_should_script_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.ScriptSource)
}

func Test_ScriptSource_should_create_json_with_source_and_lang_field_inside_script(t *testing.T) {
	t.Parallel()
	// Given
	script := es.ScriptSource("Math.max(1, doc['match_threshold'].value - 1)", ScriptLanguage.Expression)

	// When Then
	assert.NotNil(t, script)
	scriptJSON := assert.MarshalWithoutError(t, script)
	assert.Equal(t, "{\"lang\":\"expression\",\"source\":\"Math.max(1, doc['match_threshold'].value - 1)\"}", scriptJSON)
}

func Test_Script_should_have_Option_method(t *testing.T) {
	t.Parallel()
	// Given
	script := es.ScriptID("key", ScriptLanguage.Mustache)

	// When Then
	assert.NotNil(t, script.Option)
}

func Test_Script_Option_should_create_json_with_options_field_inside_script(t *testing.T) {
	t.Parallel()
	// Given
	script := es.ScriptID("key", ScriptLanguage.Mustache).
		Option("retry", "5")

	// When Then
	assert.NotNil(t, script)
	scriptJSON := assert.MarshalWithoutError(t, script)
	assert.Equal(t, "{\"id\":\"key\",\"lang\":\"mustache\",\"options\":{\"retry\":\"5\"}}", scriptJSON)
}

func Test_Script_Option_should_append_options_field_inside_script_when_options_already_exists(t *testing.T) {
	t.Parallel()
	// Given
	script := es.ScriptID("key", ScriptLanguage.Mustache).
		Option("retry", "5").
		Option("timeout", "10s").
		Option("size", "250")

	// When Then
	assert.NotNil(t, script)
	scriptJSON := assert.MarshalWithoutError(t, script)
	assert.Equal(t, "{\"id\":\"key\",\"lang\":\"mustache\",\"options\":{\"retry\":\"5\",\"size\":\"250\",\"timeout\":\"10s\"}}", scriptJSON)
}

func Test_Script_should_have_Parameter_method(t *testing.T) {
	t.Parallel()
	// Given
	script := es.ScriptSource("Math.min(tree[1])", ScriptLanguage.Java)

	// When Then
	assert.NotNil(t, script.Parameter)
}

func Test_Script_Parameter_should_create_json_with_params_field_inside_script(t *testing.T) {
	t.Parallel()
	// Given
	script := es.ScriptSource("Math.min(tree[1])", ScriptLanguage.Java).
		Parameter("p1", 100)

	// When Then
	assert.NotNil(t, script)
	scriptJSON := assert.MarshalWithoutError(t, script)
	assert.Equal(t, "{\"lang\":\"java\",\"params\":{\"p1\":100},\"source\":\"Math.min(tree[1])\"}", scriptJSON)
}

func Test_Script_Parameter_should_append_params_field_inside_script_when_params_already_exists(t *testing.T) {
	t.Parallel()
	// Given
	script := es.ScriptSource("Math.min(tree[1])", ScriptLanguage.Java).
		Parameter("p1", 100).
		Parameter("p2", "hello").
		Parameter("p3", 5.26).
		Parameter("p4", []int{22, 33, 44})

	// When Then
	assert.NotNil(t, script)
	scriptJSON := assert.MarshalWithoutError(t, script)
	// nolint:golint,lll
	assert.Equal(t, "{\"lang\":\"java\",\"params\":{\"p1\":100,\"p2\":\"hello\",\"p3\":5.26,\"p4\":[22,33,44]},\"source\":\"Math.min(tree[1])\"}", scriptJSON)
}
