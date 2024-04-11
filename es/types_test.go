package es_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/GokselKUCUKSAHIN/es-query-builder/test/assert"

	"github.com/GokselKUCUKSAHIN/es-query-builder/es"
	Mode "github.com/GokselKUCUKSAHIN/es-query-builder/es/enums/sort/mode"
	Order "github.com/GokselKUCUKSAHIN/es-query-builder/es/enums/sort/order"
)

////   NewQuery   ////

func Test_NewQuery_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.NewQuery)
}

func Test_NewQuery_should_create_a_new_Object(t *testing.T) {
	// Given When
	bodyA := es.NewQuery(nil)
	bodyB := es.NewQuery(nil)

	// Then
	assert.NotNil(t, bodyA)
	assert.NotNil(t, bodyB)
	assert.Equal(t, bodyA, bodyB)
	assert.NotEqualReference(t, bodyA, bodyB)
	assert.MarshalWithoutError(t, bodyA)
	assert.MarshalWithoutError(t, bodyB)
}

func Test_NewQuery_should_return_type_of_Object(t *testing.T) {
	// Given
	body := es.NewQuery(nil)

	// When
	bodyType := reflect.TypeOf(body).String()

	// Then
	assert.NotNil(t, body)
	assert.Equal(t, "es.Object", bodyType)
	assert.MarshalWithoutError(t, body)
}

func Test_NewQuery_should_add_query_field_into_Object(t *testing.T) {
	// Given
	body := es.NewQuery(nil)

	// When
	q, exists := body["query"]

	// Then
	assert.True(t, exists)
	assert.NotNil(t, q)
}

func Test_NewQuery_should_create_json_with_query_field(t *testing.T) {
	// Given
	body := es.NewQuery(nil)

	// When Then
	assert.NotNil(t, body)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, "{\"query\":{}}", bodyJSON)
}

func Test_NewQuery_Bool_should_create_json_with_bool_field_inside_query(t *testing.T) {
	// Given
	body := es.NewQuery(
		es.Bool(),
	)

	// When Then
	assert.NotNil(t, body)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, "{\"query\":{\"bool\":{}}}", bodyJSON)
}

////   Bool   ////

func Test_Bool_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.Bool)
}

func Test_Bool_method_should_create_boolType(t *testing.T) {
	// Given
	b := es.Bool()

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.boolType", b)
}

func Test_Bool_should_have_SetMinimumShouldMatch_method(t *testing.T) {
	// Given
	b := es.Bool()

	// When Then
	assert.NotNil(t, b.SetMinimumShouldMatch)
}

func Test_Bool_SetMinimumShouldMatch_should_create_json_with_minimum_should_match_field_inside_bool(t *testing.T) {
	// Given
	body := es.NewQuery(
		es.Bool().
			SetMinimumShouldMatch(7),
	)

	// When Then
	assert.NotNil(t, body)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, "{\"query\":{\"bool\":{\"minimum_should_match\":7}}}", bodyJSON)
}

func Test_Bool_should_have_SetBoost_method(t *testing.T) {
	// Given
	b := es.Bool()

	// When Then
	assert.NotNil(t, b.SetBoost)
}

func Test_Bool_SetBoost_should_create_json_with_minimum_should_match_field_inside_bool(t *testing.T) {
	// Given
	body := es.NewQuery(
		es.Bool().
			SetBoost(3.1415),
	)

	// When Then
	assert.NotNil(t, body)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, "{\"query\":{\"bool\":{\"boost\":3.1415}}}", bodyJSON)
}

func Test_Bool_should_have_Filter_method(t *testing.T) {
	// Given
	b := es.Bool()

	// When Then
	assert.NotNil(t, b.Filter)
}

func Test_Bool_should_have_Must_method(t *testing.T) {
	// Given
	b := es.Bool()

	// When Then
	assert.NotNil(t, b.Must)
}

func Test_Bool_should_have_MustNot_method(t *testing.T) {
	// Given
	b := es.Bool()

	// When Then
	assert.NotNil(t, b.MustNot)
}

func Test_Bool_should_have_Should_method(t *testing.T) {
	// Given
	b := es.Bool()

	// When Then
	assert.NotNil(t, b.Should)
}

////   Object   ////

func Test_Object_should_have_SetTrackTotalHits_method(t *testing.T) {
	// Given
	b := es.NewQuery(nil)

	// When Then
	assert.NotNil(t, b.SetTrackTotalHits)
}

func Test_SetTrackTotalHits_should_add_track_total_hits_field_into_Object(t *testing.T) {
	// Given
	body := es.NewQuery(nil)

	// When
	_, beforeExists := body["track_total_hits"]
	object := body.SetTrackTotalHits(true)
	trackTotalHits, afterExists := body["track_total_hits"]

	// Then
	assert.NotNil(t, object)
	assert.False(t, beforeExists)
	assert.True(t, afterExists)
	assert.True(t, trackTotalHits.(bool))
}

func Test_Object_should_have_Size_method(t *testing.T) {
	// Given
	b := es.NewQuery(nil)

	// When Then
	assert.NotNil(t, b.Size)
}

func Test_Size_should_add_size_field_into_Object(t *testing.T) {
	// Given
	body := es.NewQuery(nil)

	// When
	_, beforeExists := body["size"]
	object := body.Size(123)
	size, afterExists := body["size"]

	// Then
	assert.NotNil(t, object)
	assert.False(t, beforeExists)
	assert.True(t, afterExists)
	assert.Equal(t, 123, size.(int))
}

func Test_Object_should_have_From_method(t *testing.T) {
	// Given
	b := es.NewQuery(nil)

	// When Then
	assert.NotNil(t, b.From)
}

func Test_From_should_add_from_field_into_Object(t *testing.T) {
	// Given
	body := es.NewQuery(nil)

	// When
	_, beforeExists := body["from"]
	object := body.From(1500)
	from, afterExists := body["from"]

	// Then
	assert.NotNil(t, object)
	assert.False(t, beforeExists)
	assert.True(t, afterExists)
	assert.Equal(t, 1500, from.(int))
}

func Test_Object_should_have_Sort_method(t *testing.T) {
	// Given
	b := es.NewQuery(nil)

	// When Then
	assert.NotNil(t, b.Sort)
}

func Test_Sort_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.Sort)
}

func Test_Sort_should_return_sortType(t *testing.T) {
	// Given
	sort := es.Sort("name", Order.Asc)

	// When
	bodyType := reflect.TypeOf(sort).String()

	// Then
	assert.NotNil(t, sort)
	assert.Equal(t, "es.sortType", bodyType)
	bodyJSON := assert.MarshalWithoutError(t, sort)
	assert.Equal(t, "{\"name\":{\"order\":\"asc\"}}", bodyJSON)
}

func Test_SortWithMode_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.SortWithMode)
}

func Test_SortWithMode_should_return_sortType(t *testing.T) {
	// Given
	sort := es.SortWithMode("name", Order.Asc, Mode.Sum)

	// When
	bodyType := reflect.TypeOf(sort).String()

	// Then
	assert.NotNil(t, sort)
	assert.Equal(t, "es.sortType", bodyType)
	bodyJSON := assert.MarshalWithoutError(t, sort)
	assert.Equal(t, "{\"name\":{\"mode\":\"sum\",\"order\":\"asc\"}}", bodyJSON)
}

func Test_Sort_should_add_sort_field_into_Object(t *testing.T) {
	// Given
	body := es.NewQuery(nil)

	// When
	_, beforeExists := body["sort"]
	body.Sort(es.Sort("name", Order.Desc))
	sort, afterExists := body["sort"]

	// Then
	assert.NotNil(t, sort)
	assert.False(t, beforeExists)
	assert.True(t, afterExists)
	assert.IsTypeString(t, "[]es.sortType", sort)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, "{\"query\":{},\"sort\":[{\"name\":{\"order\":\"desc\"}}]}", bodyJSON)
}

func Test_Object_should_have_Source_method(t *testing.T) {
	// Given
	b := es.NewQuery(nil)

	// When Then
	assert.NotNil(t, b.Source)
}

func Test_Source_should_add_source_field_into_Object(t *testing.T) {
	// Given
	body := es.NewQuery(nil)

	// When
	_, beforeExists := body["_source"]
	body.Source()
	source, afterExists := body["_source"]

	// Then
	assert.NotNil(t, source)
	assert.False(t, beforeExists)
	assert.True(t, afterExists)
	assert.IsTypeString(t, "es.sourceType", source)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, "{\"_source\":{},\"query\":{}}", bodyJSON)
}

func Test_Source_should_have_Includes_method(t *testing.T) {
	// Given
	body := es.NewQuery(nil)

	// When
	source := body.Source()

	// Then
	assert.NotNil(t, source)
	assert.IsTypeString(t, "es.sourceType", source)
	assert.NotNil(t, source.Includes)
}

func Test_Source_should_have_Excludes_method(t *testing.T) {
	// Given
	body := es.NewQuery(nil)

	// When
	source := body.Source()

	// Then
	assert.NotNil(t, source)
	assert.IsTypeString(t, "es.sourceType", source)
	assert.NotNil(t, source.Excludes)
}

func Test_Source_should_create_json_with_source_field(t *testing.T) {
	// Given
	body := es.NewQuery(nil)

	// When
	body.Source().
		Includes("hello", "world").
		Excludes("Lorem", "Ipsum")

	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, "{\"_source\":{\"excludes\":[\"Lorem\",\"Ipsum\"],\"includes\":[\"hello\",\"world\"]},\"query\":{}}", bodyJSON)
}

func Test_Source_should_append_existing_fields(t *testing.T) {
	// Given
	body := es.NewQuery(nil)

	// When
	body.Source().
		Includes("hello", "world").
		Excludes("Lorem", "Ipsum").
		Includes("golang", "gopher").
		Excludes("Metallica", "Iron Maiden")

	bodyJSON := assert.MarshalWithoutError(t, body)
	//nolint:golint,lll
	assert.Equal(t, "{\"_source\":{\"excludes\":[\"Lorem\",\"Ipsum\",\"Metallica\",\"Iron Maiden\"],\"includes\":[\"hello\",\"world\",\"golang\",\"gopher\"]},\"query\":{}}", bodyJSON)
}

func Test_Object_should_have_SourceFalse_method(t *testing.T) {
	// Given
	b := es.NewQuery(nil)

	// When Then
	assert.NotNil(t, b.SourceFalse)
}

func Test_SourceFalse_should_set_source_field_as_false(t *testing.T) {
	// Given
	body := es.NewQuery(nil)

	// When
	_, beforeExists := body["_source"]
	body.SourceFalse()
	source, afterExists := body["_source"]

	// Then
	assert.NotNil(t, source)
	assert.False(t, beforeExists)
	assert.True(t, afterExists)
	assert.False(t, source.(bool))
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, "{\"_source\":false,\"query\":{}}", bodyJSON)
}

func Test_Object_should_have_Range_method(t *testing.T) {
	// Given
	b := es.NewQuery(nil)

	// When Then
	assert.NotNil(t, b.Range)
}

////   Term   ////

func Test_Term_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.Term[any])
}

func Test_Term_should_create_json_with_term_field_inside_query(t *testing.T) {
	// Given
	body := es.NewQuery(
		es.Term("key", "value"),
	)

	// When Then
	assert.NotNil(t, body)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, "{\"query\":{\"term\":{\"key\":\"value\"}}}", bodyJSON)
}

func Test_Term_method_should_create_termType(t *testing.T) {
	// Given
	b := es.Term("key", "value")

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.termType", b)
}

////   TermFunc   ////

func Test_TermFunc_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.TermFunc[any])
}

func Test_TermFunc_should_create_json_with_term_field_inside_query(t *testing.T) {
	// Given
	body := es.NewQuery(
		es.TermFunc("key", "value", func(key string, value string) bool {
			return true
		}),
	)

	// When Then
	assert.NotNil(t, body)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, "{\"query\":{\"term\":{\"key\":\"value\"}}}", bodyJSON)
}

func Test_TermFunc_should_not_add_term_field_inside_query_when_callback_result_is_false(t *testing.T) {
	// Given
	body := es.NewQuery(
		es.TermFunc("key", "value", func(key string, value string) bool {
			return false
		}),
	)

	// When Then
	assert.NotNil(t, body)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, "{\"query\":{}}", bodyJSON)
}

func Test_TermFunc_should_add_only_term_fields_inside_the_query_when_callback_result_is_true(t *testing.T) {
	// Given
	body := es.NewQuery(
		es.Bool().
			Filter(
				es.TermFunc("a", "b", func(key string, value string) bool {
					return true
				}),
				es.TermFunc("c", "d", func(key string, value string) bool {
					return false
				}),
				es.TermFunc("e", "f", func(key string, value string) bool {
					return true
				}),
			),
	)

	// When Then
	assert.NotNil(t, body)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, "{\"query\":{\"bool\":{\"filter\":[{\"term\":{\"a\":\"b\"}},{\"term\":{\"e\":\"f\"}}]}}}", bodyJSON)
}

func Test_TermFunc_method_should_create_termType(t *testing.T) {
	// Given
	b := es.TermFunc("key", "value", func(key string, value string) bool {
		return true
	})

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.termType", b)
}

////   Terms   ////

func Test_Terms_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.Terms)
}

func Test_Terms_should_create_json_with_terms_field_inside_query(t *testing.T) {
	// Given
	body := es.NewQuery(
		es.Terms("key", "value1", "value2", "value3"),
	)

	// When Then
	assert.NotNil(t, body)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, "{\"query\":{\"terms\":{\"key\":[\"value1\",\"value2\",\"value3\"]}}}", bodyJSON)
}

func Test_Terms_method_should_create_termsType(t *testing.T) {
	// Given
	b := es.Terms("key", "value1", "value2", "value3")

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.termsType", b)
}

////   TermsArray   ////

func Test_TermsArray_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.TermsArray[string])
}

func Test_TermsArray_should_create_json_with_terms_field_inside_query(t *testing.T) {
	// Given
	body := es.NewQuery(
		es.TermsArray("key", []any{"value1", "value2", "value3"}),
	)

	// When Then
	assert.NotNil(t, body)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, "{\"query\":{\"terms\":{\"key\":[\"value1\",\"value2\",\"value3\"]}}}", bodyJSON)
}

func Test_TermsArray_method_should_create_termsType(t *testing.T) {
	// Given
	b := es.TermsArray("key", []any{"value1", "value2", "value3"})

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.termsType", b)
}

////   TermsArrayFunc   ////

func Test_TermsArrayFunc_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.TermsArrayFunc[string])
}

func Test_TermsArrayFunc_should_create_json_with_terms_field_inside_query(t *testing.T) {
	// Given
	body := es.NewQuery(
		es.TermsArrayFunc("key", []string{"a", "b", "c"}, func(key string, values []string) bool {
			return true
		}),
	)

	// When Then
	assert.NotNil(t, body)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, "{\"query\":{\"terms\":{\"key\":[\"a\",\"b\",\"c\"]}}}", bodyJSON)
}

func Test_TermsArrayFunc_should_not_add_terms_field_inside_query_when_callback_result_is_false(t *testing.T) {
	// Given
	body := es.NewQuery(
		es.TermsArrayFunc("key", []string{"a", "b", "c"}, func(key string, value []string) bool {
			return false
		}),
	)

	// When Then
	assert.NotNil(t, body)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, "{\"query\":{}}", bodyJSON)
}

func Test_TermsArrayFunc_should_add_only_terms_fields_inside_the_query_when_callback_result_is_true(t *testing.T) {
	// Given
	body := es.NewQuery(
		es.Bool().
			Filter(
				es.TermsArrayFunc("a", []string{"10", "11", "12"}, func(key string, value []string) bool {
					return false
				}),
				es.TermsArrayFunc("c", []string{"20", "21", "22"}, func(key string, value []string) bool {
					return false
				}),
				es.TermsArrayFunc("e", []string{"30", "31", "32"}, func(key string, value []string) bool {
					return true
				}),
			),
	)

	// When Then
	assert.NotNil(t, body)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, "{\"query\":{\"bool\":{\"filter\":[{\"terms\":{\"e\":[\"30\",\"31\",\"32\"]}}]}}}", bodyJSON)
}

func Test_TermsArrayFunc_method_should_create_termType(t *testing.T) {
	// Given
	b := es.TermsArrayFunc("key", []string{"a", "b", "c"}, func(key string, value []string) bool {
		return true
	})

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.termsType", b)
}

////   Exists   ////

func Test_Exists_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.Exists)
}

func Test_Exists_should_create_json_with_exists_field_inside_query(t *testing.T) {
	// Given
	body := es.NewQuery(
		es.Exists("key"),
	)

	// When Then
	assert.NotNil(t, body)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, "{\"query\":{\"exists\":{\"field\":\"key\"}}}", bodyJSON)
}

func Test_Exists_method_should_create_existsType(t *testing.T) {
	// Given
	b := es.Exists("key")

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.existsType", b)
}

////   ExistsFunc   ////

func Test_ExistsFunc_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.ExistsFunc)
}

func Test_ExistsFunc_should_create_json_with_exists_field_inside_query(t *testing.T) {
	// Given
	body := es.NewQuery(
		es.ExistsFunc("key", func(key string) bool {
			return true
		}),
	)

	// When Then
	assert.NotNil(t, body)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, "{\"query\":{\"exists\":{\"field\":\"key\"}}}", bodyJSON)
}

func Test_ExistsFunc_should_not_add_exists_field_inside_query_when_callback_result_is_false(t *testing.T) {
	// Given
	body := es.NewQuery(
		es.ExistsFunc("key", func(key string) bool {
			return false
		}),
	)

	// When Then
	assert.NotNil(t, body)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, "{\"query\":{}}", bodyJSON)
}

func Test_ExistsFunc_should_add_only_exists_fields_inside_the_query_when_callback_result_is_true(t *testing.T) {
	// Given
	body := es.NewQuery(
		es.Bool().
			Filter(
				es.ExistsFunc("a", func(key string) bool {
					return false
				}),
				es.ExistsFunc("c", func(key string) bool {
					return true
				}),
				es.ExistsFunc("e", func(key string) bool {
					return true
				}),
			),
	)

	// When Then
	assert.NotNil(t, body)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, "{\"query\":{\"bool\":{\"filter\":[{\"exists\":{\"field\":\"c\"}},{\"exists\":{\"field\":\"e\"}}]}}}", bodyJSON)
}

func Test_ExistsFunc_method_should_create_existsType(t *testing.T) {
	// Given
	b := es.ExistsFunc("key", func(key string) bool {
		return true
	})

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.existsType", b)
}

////   Range   ////

func Test_Range_method_should_create_rangeType(t *testing.T) {
	// Given
	body := es.NewQuery(nil)
	b := body.Range("age")

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.rangeType", b)
}

func Test_Range_should_create_json_with_range_field_inside_query(t *testing.T) {
	// Given
	body := es.NewQuery(nil)
	body.Range("age").
		GreaterThanOrEqual(10).
		LesserThanOrEqual(20)

	// When Then
	assert.NotNil(t, body)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, "{\"query\":{\"range\":{\"age\":{\"gte\":10,\"lte\":20}}}}", bodyJSON)
}

func Test_Range_gte_should_override_gt_and_vise_versa(t *testing.T) {
	// Given
	body := es.NewQuery(nil)
	body.Range("age").
		GreaterThanOrEqual(10).
		GreaterThan(20)

	// When Then
	assert.NotNil(t, body)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, "{\"query\":{\"range\":{\"age\":{\"gt\":20}}}}", bodyJSON)
}

func Test_Range_lte_should_override_lt_and_vise_versa(t *testing.T) {
	// Given
	body := es.NewQuery(nil)
	body.Range("age").
		LesserThan(11).
		LesserThanOrEqual(23)

	// When Then
	assert.NotNil(t, body)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, "{\"query\":{\"range\":{\"age\":{\"lte\":23}}}}", bodyJSON)
}

func Test_Range_should_not_range_field_when_no_query_field_in_Object(t *testing.T) {
	// Given
	body := es.Object{}
	body.Range("age").
		GreaterThanOrEqual(10).
		LesserThanOrEqual(20)

	// When Then
	assert.NotNil(t, body)
	bodyJSON := assert.MarshalWithoutError(t, body)
	assert.Equal(t, "{}", bodyJSON)
}

////   Bool.Filter   ////

func Test_Filter_method_should_return_boolType(t *testing.T) {
	// Given
	b := es.Bool()

	// When
	filter := b.Filter()

	// Then
	assert.NotNil(t, filter)
	assert.IsTypeString(t, "es.boolType", filter)
}

func Test_Filter_method_should_add_filter_if_doesnt_exist_before(t *testing.T) {
	// Given
	b := es.Bool()

	// When
	_, beforeExists := b["filter"]
	filter := b.Filter()
	_, afterExists := b["filter"]

	// Then
	assert.NotNil(t, filter)
	assert.False(t, beforeExists)
	assert.True(t, afterExists)
}

func Test_Filter_method_should_hold_items(t *testing.T) {
	// Given
	b := es.Bool().
		Filter(
			es.Term("id", 12345),
		)

	// When
	filter, exists := b["filter"]

	// Then
	assert.True(t, exists)
	assert.IsTypeString(t, "es.filterType", filter)

	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"filter\":[{\"term\":{\"id\":12345}}]}", bodyJSON)
}

////   Bool.Must   ////

func Test_Must_method_should_return_boolType(t *testing.T) {
	// Given
	b := es.Bool()

	// When
	must := b.Must()

	// Then
	assert.NotNil(t, must)
	assert.IsTypeString(t, "es.boolType", must)
}

func Test_Must_method_should_add_must_if_doesnt_exist_before(t *testing.T) {
	// Given
	b := es.Bool()

	// When
	_, beforeExists := b["must"]
	filter := b.Must()
	_, afterExists := b["must"]

	// Then
	assert.NotNil(t, filter)
	assert.False(t, beforeExists)
	assert.True(t, afterExists)
}

func Test_Must_method_should_hold_items(t *testing.T) {
	// Given
	b := es.Bool().
		Must(
			es.Term("id", 12345),
		)

	// When
	must, exists := b["must"]

	// Then
	assert.True(t, exists)
	assert.IsTypeString(t, "es.mustType", must)

	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"must\":[{\"term\":{\"id\":12345}}]}", bodyJSON)
}

////   Bool.MustNot   ////

func Test_MustNot_method_should_return_boolType(t *testing.T) {
	// Given
	b := es.Bool()

	// When
	mustNot := b.MustNot()

	// Then
	assert.NotNil(t, mustNot)
	assert.IsTypeString(t, "es.boolType", mustNot)
}

func Test_MustNot_method_should_add_must_not_if_doesnt_exist_before(t *testing.T) {
	// Given
	b := es.Bool()

	// When
	_, beforeExists := b["must_not"]
	filter := b.MustNot()
	_, afterExists := b["must_not"]

	// Then
	assert.NotNil(t, filter)
	assert.False(t, beforeExists)
	assert.True(t, afterExists)
}

func Test_MustNot_method_should_hold_items(t *testing.T) {
	// Given
	b := es.Bool().
		MustNot(
			es.Term("id", 12345),
		)

	// When
	mustNot, exists := b["must_not"]

	// Then
	assert.True(t, exists)
	assert.IsTypeString(t, "es.mustNotType", mustNot)

	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"must_not\":[{\"term\":{\"id\":12345}}]}", bodyJSON)
}

////   Bool.Should   ////

func Test_Should_method_should_return_boolType(t *testing.T) {
	// Given
	b := es.Bool()

	// When
	should := b.Should()

	// Then
	assert.NotNil(t, should)
	assert.IsTypeString(t, "es.boolType", should)
}

func Test_Should_method_should_add_should_if_doesnt_exist_before(t *testing.T) {
	// Given
	b := es.Bool()

	// When
	_, beforeExists := b["should"]
	filter := b.Should()
	_, afterExists := b["should"]

	// Then
	assert.NotNil(t, filter)
	assert.False(t, beforeExists)
	assert.True(t, afterExists)
}

func Test_Should_method_should_hold_items(t *testing.T) {
	// Given
	b := es.Bool().
		Should(
			es.Term("id", 12345),
		)

	// When
	should, exists := b["should"]

	// Then
	assert.True(t, exists)
	assert.IsTypeString(t, "es.shouldType", should)

	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"should\":[{\"term\":{\"id\":12345}}]}", bodyJSON)
}

// CTE

func Test_It_Runs(t *testing.T) {
	query := es.NewQuery(
		es.Bool().Filter(
			es.Term("content.culture", "tr-TR"),
			es.Nested("listings", func() any {
				return es.Bool().
					Filter(
						es.Term("listings.supplierId", 73),
						es.Term("listings.statusTypes", "ALL"),
					)
			}).SetInnerHits(es.Object{"size": 10000}),
		),
	)

	marshal, _ := json.Marshal(query)

	print(string(marshal))
}
