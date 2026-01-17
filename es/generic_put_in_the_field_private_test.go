package es

import (
	"testing"

	"github.com/Trendyol/es-query-builder/test/assert"
)

////   genericPutInTheField   ////

func Test_genericPutInTheField_should_put_value_on_desired_key_of_root_object(t *testing.T) {
	t.Parallel()

	// Given
	myObject := Object{"trendyol": Object{"team": "crm"}}

	// When
	passThrough := genericPutInTheField(myObject, "trendyol", "emoji", "🧡")

	// Then
	assert.NotNil(t, myObject)
	assert.NotNil(t, passThrough)

	originalBodyJSON := assert.MarshalWithoutError(t, myObject)
	passThroughBodyJSON := assert.MarshalWithoutError(t, passThrough)
	assert.Equal(t, originalBodyJSON, passThroughBodyJSON)
	assert.Equal(t, "{\"trendyol\":{\"emoji\":\"🧡\",\"team\":\"crm\"}}", passThroughBodyJSON)
}

func Test_genericPutInTheField_should_add_field_to_existing_object(t *testing.T) {
	t.Parallel()

	// Given
	root := avgAggType{"avg": Object{"field": "price"}}

	// When
	result := genericPutInTheField(root, "avg", "missing", 0)

	// Then
	assert.NotNil(t, result)
	bodyJSON := assert.MarshalWithoutError(t, result)
	assert.Equal(t, "{\"avg\":{\"field\":\"price\",\"missing\":0}}", bodyJSON)
}

func Test_genericPutInTheField_should_return_unchanged_when_parent_key_not_exists(t *testing.T) {
	t.Parallel()

	// Given
	root := Object{"other": Object{"field": "value"}}

	// When
	result := genericPutInTheField(root, "nonexistent", "key", "value")

	// Then
	assert.NotNil(t, result)
	bodyJSON := assert.MarshalWithoutError(t, result)
	assert.Equal(t, "{\"other\":{\"field\":\"value\"}}", bodyJSON)
}

func Test_genericPutInTheField_should_return_unchanged_when_parent_key_is_not_object(t *testing.T) {
	t.Parallel()

	// Given
	root := Object{"parent": "string_value"}

	// When
	result := genericPutInTheField(root, "parent", "key", "value")

	// Then
	assert.NotNil(t, result)
	bodyJSON := assert.MarshalWithoutError(t, result)
	assert.Equal(t, "{\"parent\":\"string_value\"}", bodyJSON)
}

func Test_genericPutInTheField_should_work_with_different_types(t *testing.T) {
	t.Parallel()

	// Given
	root := termsAggType{"terms": Object{"field": "category"}}

	// When
	result := genericPutInTheField(root, "terms", "size", 10)

	// Then
	assert.NotNil(t, result)
	bodyJSON := assert.MarshalWithoutError(t, result)
	assert.Equal(t, "{\"terms\":{\"field\":\"category\",\"size\":10}}", bodyJSON)
}

func Test_genericPutInTheField_should_overwrite_existing_key(t *testing.T) {
	t.Parallel()

	// Given
	root := Object{"parent": Object{"key": "old_value"}}

	// When
	result := genericPutInTheField(root, "parent", "key", "new_value")

	// Then
	assert.NotNil(t, result)
	bodyJSON := assert.MarshalWithoutError(t, result)
	assert.Equal(t, "{\"parent\":{\"key\":\"new_value\"}}", bodyJSON)
}

////   genericPutInTheFieldOfFirstChild   ////

func Test_genericPutInTheFieldOfFirstChild_should_put_value_in_first_child_object(t *testing.T) {
	t.Parallel()

	// Given
	root := termType{"term": Object{"status": Object{"value": "active"}}}

	// When
	result := genericPutInTheFieldOfFirstChild(root, "term", "boost", 2.0)

	// Then
	assert.NotNil(t, result)
	bodyJSON := assert.MarshalWithoutError(t, result)
	assert.Equal(t, "{\"term\":{\"status\":{\"boost\":2,\"value\":\"active\"}}}", bodyJSON)
}

func Test_genericPutInTheFieldOfFirstChild_should_work_with_match_query(t *testing.T) {
	t.Parallel()

	// Given
	root := matchType{"match": Object{"message": Object{"query": "hello"}}}

	// When
	result := genericPutInTheFieldOfFirstChild(root, "match", "operator", "and")

	// Then
	assert.NotNil(t, result)
	bodyJSON := assert.MarshalWithoutError(t, result)
	assert.Equal(t, "{\"match\":{\"message\":{\"operator\":\"and\",\"query\":\"hello\"}}}", bodyJSON)
}

func Test_genericPutInTheFieldOfFirstChild_should_return_unchanged_when_parent_key_not_exists(t *testing.T) {
	t.Parallel()

	// Given
	root := Object{"other": Object{"field": Object{"value": "test"}}}

	// When
	result := genericPutInTheFieldOfFirstChild(root, "nonexistent", "key", "value")

	// Then
	assert.NotNil(t, result)
	bodyJSON := assert.MarshalWithoutError(t, result)
	assert.Equal(t, "{\"other\":{\"field\":{\"value\":\"test\"}}}", bodyJSON)
}

func Test_genericPutInTheFieldOfFirstChild_should_return_unchanged_when_parent_is_not_object(t *testing.T) {
	t.Parallel()

	// Given
	root := Object{"parent": "string_value"}

	// When
	result := genericPutInTheFieldOfFirstChild(root, "parent", "key", "value")

	// Then
	assert.NotNil(t, result)
	bodyJSON := assert.MarshalWithoutError(t, result)
	assert.Equal(t, "{\"parent\":\"string_value\"}", bodyJSON)
}

func Test_genericPutInTheFieldOfFirstChild_should_return_unchanged_when_no_child_object_found(t *testing.T) {
	t.Parallel()

	// Given
	root := Object{"parent": Object{"field": "string_value"}}

	// When
	result := genericPutInTheFieldOfFirstChild(root, "parent", "key", "value")

	// Then
	assert.NotNil(t, result)
	bodyJSON := assert.MarshalWithoutError(t, result)
	assert.Equal(t, "{\"parent\":{\"field\":\"string_value\"}}", bodyJSON)
}

func Test_genericPutInTheFieldOfFirstChild_should_work_with_range_query(t *testing.T) {
	t.Parallel()

	// Given
	root := rangeType{"range": Object{"age": Object{"gte": 18}}}

	// When
	result := genericPutInTheFieldOfFirstChild(root, "range", "lte", 65)

	// Then
	assert.NotNil(t, result)
	bodyJSON := assert.MarshalWithoutError(t, result)
	assert.Equal(t, "{\"range\":{\"age\":{\"gte\":18,\"lte\":65}}}", bodyJSON)
}

func Test_genericPutInTheFieldOfFirstChild_should_only_modify_first_child(t *testing.T) {
	t.Parallel()

	// Given
	root := matchBoolPrefixType{"match_bool_prefix": Object{"message": Object{"query": "hello"}}}

	// When
	result := genericPutInTheFieldOfFirstChild(root, "match_bool_prefix", "analyzer", "standard")

	// Then
	assert.NotNil(t, result)
	bodyJSON := assert.MarshalWithoutError(t, result)
	assert.Equal(t, "{\"match_bool_prefix\":{\"message\":{\"analyzer\":\"standard\",\"query\":\"hello\"}}}", bodyJSON)
}

////   genericPutInTheFieldOfFirstObject   ////

func Test_genericPutInTheFieldOfFirstObject_should_put_value_in_first_object(t *testing.T) {
	t.Parallel()

	// Given
	root := queryStringType{"query_string": Object{"query": "hello"}}

	// When
	result := genericPutInTheFieldOfFirstObject(root, "default_field", "message")

	// Then
	assert.NotNil(t, result)
	bodyJSON := assert.MarshalWithoutError(t, result)
	assert.Equal(t, "{\"query_string\":{\"default_field\":\"message\",\"query\":\"hello\"}}", bodyJSON)
}

func Test_genericPutInTheFieldOfFirstObject_should_work_with_simple_query_string(t *testing.T) {
	t.Parallel()

	// Given
	root := simpleQueryStringType{"simple_query_string": Object{"query": "test"}}

	// When
	result := genericPutInTheFieldOfFirstObject(root, "default_operator", "and")

	// Then
	assert.NotNil(t, result)
	bodyJSON := assert.MarshalWithoutError(t, result)
	assert.Equal(t, "{\"simple_query_string\":{\"default_operator\":\"and\",\"query\":\"test\"}}", bodyJSON)
}

func Test_genericPutInTheFieldOfFirstObject_should_work_with_sort(t *testing.T) {
	t.Parallel()

	// Given
	root := sortType{"timestamp": Object{"order": "desc"}}

	// When
	result := genericPutInTheFieldOfFirstObject(root, "mode", "avg")

	// Then
	assert.NotNil(t, result)
	bodyJSON := assert.MarshalWithoutError(t, result)
	assert.Equal(t, "{\"timestamp\":{\"mode\":\"avg\",\"order\":\"desc\"}}", bodyJSON)
}

func Test_genericPutInTheFieldOfFirstObject_should_return_unchanged_when_no_object_found(t *testing.T) {
	t.Parallel()

	// Given
	root := Object{"field": "string_value"}

	// When
	result := genericPutInTheFieldOfFirstObject(root, "key", "value")

	// Then
	assert.NotNil(t, result)
	bodyJSON := assert.MarshalWithoutError(t, result)
	assert.Equal(t, "{\"field\":\"string_value\"}", bodyJSON)
}

func Test_genericPutInTheFieldOfFirstObject_should_return_unchanged_when_empty_map(t *testing.T) {
	t.Parallel()

	// Given
	root := Object{}

	// When
	result := genericPutInTheFieldOfFirstObject(root, "key", "value")

	// Then
	assert.NotNil(t, result)
	bodyJSON := assert.MarshalWithoutError(t, result)
	assert.Equal(t, "{}", bodyJSON)
}

func Test_genericPutInTheFieldOfFirstObject_should_overwrite_existing_key(t *testing.T) {
	t.Parallel()

	// Given
	root := Object{"wrapper": Object{"key": "old_value"}}

	// When
	result := genericPutInTheFieldOfFirstObject(root, "key", "new_value")

	// Then
	assert.NotNil(t, result)
	bodyJSON := assert.MarshalWithoutError(t, result)
	assert.Equal(t, "{\"wrapper\":{\"key\":\"new_value\"}}", bodyJSON)
}

func Test_genericPutInTheFieldOfFirstObject_should_add_multiple_fields(t *testing.T) {
	t.Parallel()

	// Given
	root := queryStringType{"query_string": Object{"query": "search term"}}

	// When
	result := genericPutInTheFieldOfFirstObject(root, "default_field", "title")
	result = genericPutInTheFieldOfFirstObject(result, "default_operator", "AND")

	// Then
	assert.NotNil(t, result)
	bodyJSON := assert.MarshalWithoutError(t, result)
	assert.Equal(t, "{\"query_string\":{\"default_field\":\"title\",\"default_operator\":\"AND\",\"query\":\"search term\"}}", bodyJSON)
}
