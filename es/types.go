package es

import (
	"unsafe"

	Operator "github.com/GokselKUCUKSAHIN/es-query-builder/es/enums/match/operator"
	ScoreMode "github.com/GokselKUCUKSAHIN/es-query-builder/es/enums/nested/score-mode"
	Mode "github.com/GokselKUCUKSAHIN/es-query-builder/es/enums/sort/mode"
	Order "github.com/GokselKUCUKSAHIN/es-query-builder/es/enums/sort/order"
)

type Object map[string]any

type Array []any

type boolType Object

type filterType Array

type mustType Array

type mustNotType Array

type shouldType Array

type matchType Object

type matchAllType Object

type matchNoneType Object

type termType Object

type termsType Object

type existsType Object

type rangeType Object

type sortType Object

type sourceType Object

type includesType Array

type excludesType Array

type nestedType Object

type aggsType Object

type aggTermType Object

func unsafeIsNil(x any) bool {
	return (*[2]uintptr)(unsafe.Pointer(&x))[1] == 0
}

func correctType(b any) (any, bool) {
	if b == nil || unsafeIsNil(b) {
		return Object{}, false
	}
	switch b.(type) {
	case boolType:
		return Object{"bool": b}, true
	case rangeType:
		return Object{"range": b}, true
	}
	return b, true
}

// NewQuery constructs and returns an es.Object that represents a query.
//
// The function takes an input `queryClause` of any type and attempts to convert it into
// a specific query type based on its value. If the type of `queryClause` matches a
// recognized query type (as determined by the correctType function), the
// function returns an es.Object with the corresponding query type as its key.
//
// If `queryClause` is nil, unsafeIsNil returns true, or the type is not recognized,
// NewQuery returns an es.Object with an empty query.
//
// Recognized types include:
//   - boolType: Creates a "bool" query.
//   - rangeType: Creates a "range" query.
//
// Example usage:
//
//	query := NewQuery(es.Bool())
//	// query now holds an Object with a "bool" key.
//
// If the type of `queryClause` is not recognized:
//
//	query := NewQuery(nil)
//	// query now holds an Object with an empty query.
//
// Parameters:
//   - queryClause: An input of any type that may represent a specific query.
//
// Returns:
//
//	An Object containing the query or an empty query if `queryClause` is not recognized.
func NewQuery(queryClause any) Object {
	if field, ok := correctType(queryClause); ok {
		return Object{
			"query": field,
		}
	}
	return Object{
		"query": Object{},
	}
}

func (o Object) putInQuery(key string, value any) Object {
	if query, exists := o["query"]; exists {
		query.(Object)[key] = value
	}
	return o
}

// Bool creates and returns an empty boolType object.
//
// This function is typically used to initialize a boolType, which can be
// populated later with the appropriate boolean query conditions.
//
// Example usage:
//
//	b := Bool()
//	// b is now an empty boolType object that can be used in a query.
//
// Returns:
//
//	An empty boolType object.
func Bool() boolType {
	return boolType{}
}

// MinimumShouldMatch sets the "minimum_should_match" parameter in a boolType query.
//
// This method allows you to specify the minimum number of "should" clauses
// that must match in a boolean query. The "minimum_should_match" parameter
// is often used to fine-tune the results of a boolean query by requiring a
// certain number of optional (should) clauses to match.
//
// Example usage:
//
//	b := Bool().MinimumShouldMatch(2)
//	// b now includes a "minimum_should_match" parameter with a value of 2.
//
// Parameters:
//   - minimumShouldMatch: An integer specifying the minimum number of "should"
//     clauses that must match for a document to be considered a match.
//
// Returns:
//
//	The updated boolType object with the "minimum_should_match" parameter set.
func (b boolType) MinimumShouldMatch(minimumShouldMatch int) boolType {
	b["minimum_should_match"] = minimumShouldMatch
	return b
}

// Boost sets the "boost" parameter in a boolType query.
//
// This method allows you to assign a boost value to a boolean query, which
// can be used to increase or decrease the relevance score of the query's
// results. The boost parameter is a multiplier that influences how strongly
// this particular query should be considered compared to others.
//
// Example usage:
//
//	b := Bool().Boost(1.5)
//	// b now includes a "boost" parameter with a value of 1.5.
//
// Parameters:
//   - boost: A floating-point number representing the boost factor. Higher
//     values increase the importance of the query in scoring, while lower
//     values decrease it.
//
// Returns:
//
//	The updated boolType object with the "boost" parameter set.
func (b boolType) Boost(boost float64) boolType {
	b["boost"] = boost
	return b
}

// Filter adds one or more filter conditions to the boolType object.
//
// This method updates the "filter" section of the boolType object by appending
// the specified filter conditions. It accepts a variadic number of filter conditions,
// checks their types, and adds them to the "filter" array in the boolType object.
//
// Example usage:
//
//	b := es.Bool().
//	Filter(
//	       es.Term("title", "es-query-builder"),
//	       es.Exists("filter"),
//	)
//	// b now contains the provided conditions in the "filter" array.
//
// Parameters:
//   - items: A variadic list of filter conditions to be added. The type is generic.
//
// Returns:
//
//	The updated boolType object with the new filter conditions added.
func (b boolType) Filter(items ...any) boolType {
	filter, exists := b["filter"]
	if !exists {
		filter = filterType{}
	}
	for i := 0; i < len(items); i++ {
		if field, ok := correctType(items[i]); ok {
			filter = append(filter.(filterType), field)
		}
	}
	b["filter"] = filter
	return b
}

// Must adds one or more conditions to the "must" section of the boolType object.
//
// This method updates the "must" section by appending the specified conditions.
// It accepts a variadic number of conditions, checks their types, and adds them to
// the "must" array in the boolType object.
//
// Example usage:
//
//	b := es.Bool().
//	Must(
//	       es.Term("title", "es-query-builder"),
//	       es.Exists("must"),
//	)
//	// b now contains the provided conditions in the "must" array.
//
// Parameters:
//   - items: A variadic list of conditions to be added. The type is generic.
//
// Returns:
//
//	The updated boolType object with the new conditions added to the "must" section.
func (b boolType) Must(items ...any) boolType {
	must, exists := b["must"]
	if !exists {
		must = mustType{}
	}
	for i := 0; i < len(items); i++ {
		if field, ok := correctType(items[i]); ok {
			must = append(must.(mustType), field)
		}
	}
	b["must"] = must
	return b
}

// MustNot adds one or more conditions to the "must_not" section of the boolType object.
//
// This method updates the "must_not" section by appending the specified conditions.
// It accepts a variadic number of conditions, checks their types, and adds them to
// the "must_not" array in the boolType object.
//
// Example usage:
//
//	b := es.Bool().
//	MustNot(
//	       es.Term("title", "es-query-builder"),
//	       es.Exists("mustNot"),
//	)
//	// b now contains the provided conditions in the "must_not" array.
//
// Parameters:
//   - items: A variadic list of conditions to be added. The type is generic.
//
// Returns:
//
//	The updated boolType object with the new conditions added to the "must_not" section.
func (b boolType) MustNot(items ...any) boolType {
	mustNot, exists := b["must_not"]
	if !exists {
		mustNot = mustNotType{}
	}
	for i := 0; i < len(items); i++ {
		if field, ok := correctType(items[i]); ok {
			mustNot = append(mustNot.(mustNotType), field)
		}
	}
	b["must_not"] = mustNot
	return b
}

// Should adds one or more conditions to the "should" section of the boolType object.
//
// This method updates the "should" section by appending the specified conditions.
// It accepts a variadic number of conditions, checks their types, and adds them to
// the "should" array in the boolType object.
//
// Example usage:
//
//	b := es.Bool().
//	Should(
//	       es.Term("title", "es-query-builder"),
//	       es.Exists("should"),
//	)
//	// b now contains the provided conditions in the "should" array.
//
// Parameters:
//   - items: A variadic list of conditions to be added. The type is generic.
//
// Returns:
//
//	The updated boolType object with the new conditions added to the "should" section.
func (b boolType) Should(items ...any) boolType {
	should, exists := b["should"]
	if !exists {
		should = shouldType{}
	}
	for i := 0; i < len(items); i++ {
		if field, ok := correctType(items[i]); ok {
			should = append(should.(shouldType), field)
		}
	}
	b["should"] = should
	return b
}

// TrackTotalHits sets the "track_total_hits" parameter in an es.Object.
//
// This method allows you to specify whether the total number of hits should
// be tracked in the search results. When set to true, the total number of
// matching documents is included in the response. This is useful for
// pagination and to understand the overall size of the result set.
//
// Example usage:
//
//	query := es.NewQuery(...).TrackTotalHits(true)
//	// query now includes a "track_total_hits" parameter with a value of true.
//
// Parameters:
//   - value: A boolean indicating whether to track the total number of hits.
//     Set to true to include the total count in the response; false to exclude it.
//
// Returns:
//
//	The updated Object with the "track_total_hits" parameter set.
func (o Object) TrackTotalHits(value bool) Object {
	o["track_total_hits"] = value
	return o
}

// Size sets the "size" parameter in an es.Object.
//
// This method specifies the number of search results to return. It controls
// the maximum number of documents that will be included in the search response.
// This is useful for limiting the size of the result set, especially when dealing
// with large datasets or paginating results.
//
// Example usage:
//
//	query := es.NewQuery(...).Size(10)
//	// query now includes a "size" parameter with a value of 10, limiting results to 10 documents.
//
// Parameters:
//   - size: An integer specifying the number of search results to return.
//     Set this value to control the maximum number of documents in the response.
//
// Returns:
//
//	The updated Object with the "size" parameter set.
func (o Object) Size(size int) Object {
	o["size"] = size
	return o
}

// From sets the "from" parameter in an es.Object.
//
// This method specifies the starting point (offset) for the search results.
// It is used to skip a certain number of documents before starting to return
// the results. This is useful for pagination, allowing you to fetch results
// starting from a specific index.
//
// Example usage:
//
//	query := es.NewQuery(...).From(20)
//	// query now includes a "from" parameter with a value of 20, starting results from the 21st document.
//
// Parameters:
//   - from: An integer specifying the starting point (offset) for the search results.
//     Set this value to skip a certain number of documents before beginning the result set.
//
// Returns:
//
//	The updated Object with the "from" parameter set.
func (o Object) From(from int) Object {
	o["from"] = from
	return o
}

// Sort creates a new sortType object with the specified field.
//
// This function initializes a sortType object with a given field name. The
// field is used to specify the sorting criteria in the search query. The
// resulting sortType can be further configured with sorting order and mode.
//
// Example usage:
//
//	s := Sort("age")
//	// s now includes a sortType with a "age" field that can be further configured.
//
// Parameters:
//   - field: A string representing the field to sort by.
//
// Returns:
//
//	A sortType object with the specified field.
func Sort(field string) sortType {
	return sortType{
		field: Object{},
	}
}

func (s sortType) putInTheField(key string, value any) sortType {
	for field := range s {
		if fieldObject, ok := s[field].(Object); ok {
			fieldObject[key] = value
		}
	}
	return s
}

// Order sets the "order" parameter in a sortType object.
//
// This method specifies the order in which the results should be sorted.
// It configures the sortType object to sort the results in ascending or
// descending order.
//
// Example usage:
//
//	s := Sort("age").Order(Order.Desc)
//	// s now includes an "order" parameter with the value "desc".
//
// Parameters:
//   - order: An Order.Order value indicating the sorting order (e.g., ascending or descending).
//
// Returns:
//
//	The updated sortType object with the "order" parameter set.
func (s sortType) Order(order Order.Order) sortType {
	return s.putInTheField("order", order)
}

// Mode sets the "mode" parameter in a sortType object.
//
// This method specifies the mode used for sorting the results. The mode
// determines how sorting should be handled, such as by specifying different
// tie-breaking strategies.
//
// Example usage:
//
//	s := Sort("age").Mode(Mode.Avg)
//	// s now includes a "mode" parameter with the value "avg".
//
// Parameters:
//   - mode: A Mode.Mode value indicating the sorting mode (e.g., average, minimum, maximum).
//
// Returns:
//
//	The updated sortType object with the "mode" parameter set.
func (s sortType) Mode(mode Mode.Mode) sortType {
	return s.putInTheField("mode", mode)
}

// Sort adds one or more sortType objects to an es.Object.
//
// This method allows you to specify multiple sorting criteria for the search query.
// Each sortType object defines how the results should be sorted based on different fields.
//
// Example usage:
//
//	query := es.NewQuery(...).Sort(Sort("age").Order(Order.Desc), Sort("date").Order(Order.Asc))
//	// query now includes a "sort" parameter with multiple sortType objects.
//
// Parameters:
//   - sorts: A variadic list of sortType objects, each specifying sorting criteria.
//
// Returns:
//
//	The updated Object with the "sort" parameter set, containing the provided sortType objects.
func (o Object) Sort(sorts ...sortType) Object {
	o["sort"] = sorts
	return o
}

// Source initializes and returns a sourceType object in the es.Object.
//
// This method sets the "_source" field in the es.Object to a new, empty sourceType object.
// The sourceType object is used to specify which fields should be included or excluded
// from the source data in the search results.
//
// Example usage:
//
//	src := es.NewQuery(...).Source()
//	// src now has a "_source" field set to an empty sourceType object.
//
// Returns:
//
//	A sourceType object, with the "_source" field of the es.Object set to this new object.
func (o Object) Source() sourceType {
	s := sourceType{}
	o["_source"] = s
	return s
}

// SourceFalse sets the "_source" field to false in the es.Object.
//
// This method configures the es.Object to not include the source data in the search results.
// Setting the "_source" field to false excludes the entire source field from the response.
//
// Example usage:
//
//	query := es.NewQuery().SourceFalse()
//	// query now has a "_source" field set to false.
//
// Returns:
//
//	The updated es.Object with the "_source" field set to false.
func (o Object) SourceFalse() Object {
	o["_source"] = false
	return o
}

// Includes adds one or more fields to be included in the sourceType object.
//
// This method updates the sourceType object to specify which fields should be included
// in the search results. If the "includes" key does not already exist, it initializes
// it with an empty includesType slice before appending the new fields.
//
// Example usage:
//
//	s := Source().Includes("title", "author")
//	// s now has an "includes" parameter with "title" and "author" fields.
//
// Parameters:
//   - fields: A variadic list of strings specifying the fields to be included.
//
// Returns:
//
//	The updated sourceType object with the "includes" parameter set to the specified fields.
func (s sourceType) Includes(fields ...string) sourceType {
	includes, exists := s["includes"]
	if !exists {
		includes = includesType{}
	}
	for i := 0; i < len(fields); i++ {
		includes = append(includes.(includesType), fields[i])
	}
	s["includes"] = includes
	return s
}

// Excludes adds one or more fields to be excluded from the sourceType object.
//
// This method updates the sourceType object to specify which fields should be excluded
// from the search results. If the "excludes" key does not already exist, it initializes
// it with an empty excludesType slice before appending the new fields.
//
// Example usage:
//
//	s := Source().Excludes("metadata", "private")
//	// s now has an "excludes" parameter with "metadata" and "private" fields.
//
// Parameters:
//   - fields: A variadic list of strings specifying the fields to be excluded.
//
// Returns:
//
//	The updated sourceType object with the "excludes" parameter set to the specified fields.
func (s sourceType) Excludes(fields ...string) sourceType {
	excludes, exists := s["excludes"]
	if !exists {
		excludes = excludesType{}
	}
	for i := 0; i < len(fields); i++ {
		excludes = append(excludes.(excludesType), fields[i])
	}
	s["excludes"] = excludes
	return s
}

// Term creates a new termType object with the specified key-value pair.
//
// This function initializes a termType object with a single term query, where the
// key is the field name and the value is the term to search for. This is typically
// used to construct a term query in search queries.
//
// Example usage:
//
//	t := Term("category", "books")
//	// t now contains a termType object with a term query for the "category" field.
//
// Parameters:
//   - key: A string representing the field name for the term query.
//   - value: The value to be searched for in the specified field. The type is generic.
//
// Returns:
//
//	A termType object containing the specified term query.
func Term[T any](key string, value T) termType {
	return termType{
		"term": Object{
			key: value,
		},
	}
}

// TermFunc creates a termType object based on a condition evaluated by a function.
//
// This function conditionally creates a termType object if the provided function
// returns true for the given key-value pair. If the function returns false, it
// returns nil instead of creating a termType object.
//
// Example usage:
//
//	t := TermFunc("category", "books", func(key, value string) bool {
//	    return value != ""
//	})
//	// t is either a termType object or nil based on the condition.
//
// Parameters:
//   - key: A string representing the field name for the term query.
//   - value: The value to be searched for in the specified field. The type is generic.
//   - f: A function that takes a key and value, and returns a boolean indicating
//     whether to create the termType object.
//
// Returns:
//
//	A termType object if the condition is true; otherwise, nil.
func TermFunc[T any](key string, value T, f func(key string, value T) bool) termType {
	if !f(key, value) {
		return nil
	}
	return Term(key, value)
}

// TermIf creates a termType object based on a boolean condition.
//
// This function creates a termType object if the provided condition is true. If
// the condition is false, it returns nil instead of creating a termType object.
//
// Example usage:
//
//	t := TermIf("category", "books", true)
//	// t is a termType object if the condition is true; otherwise, it is nil.
//
// Parameters:
//   - key: A string representing the field name for the term query.
//   - value: The value to be searched for in the specified field. The type is generic.
//   - condition: A boolean value that determines whether to create the termType object.
//
// Returns:
//
//	A termType object if the condition is true; otherwise, nil.
func TermIf[T any](key string, value T, condition bool) termType {
	if !condition {
		return nil
	}
	return Term(key, value)
}

// Terms creates a new termsType object with the specified key and values.
//
// This function initializes a termsType object for a terms query, where the key
// is the field name and values is a variadic list of terms to search for in that field.
// This is used to construct queries that match any of the specified terms.
//
// Example usage:
//
//	t := Terms("category", "books", "electronics")
//	// t now contains a termsType object with a terms query for the "category" field.
//
// Parameters:
//   - key: A string representing the field name for the terms query.
//   - values: A variadic list of values to be searched for in the specified field.
//     The type is generic.
//
// Returns:
//
//	A termsType object containing the specified terms query.
func Terms(key string, values ...any) termsType {
	return termsType{
		"terms": Object{
			key: values,
		},
	}
}

// TermsArray creates a new termsType object with the specified key and values as a slice.
//
// This function initializes a termsType object for a terms query, where the key
// is the field name and values is a slice of terms to search for in that field.
// This is useful for cases where the terms are provided as a slice instead of
// a variadic list.
//
// Example usage:
//
//	t := TermsArray("category", []string{"books", "electronics"})
//	// t now contains a termsType object with a terms query for the "category" field.
//
// Parameters:
//   - key: A string representing the field name for the terms query.
//   - values: A slice of values to be searched for in the specified field.
//     The type is generic.
//
// Returns:
//
//	A termsType object containing the specified terms query.
func TermsArray[T any](key string, values []T) termsType {
	return termsType{
		"terms": Object{
			key: values,
		},
	}
}

// TermsFunc creates a termsType object based on a condition evaluated by a function.
//
// This function conditionally creates a termsType object if the provided function
// returns true for the given key and values. If the function returns false, it
// returns nil instead of creating a termsType object.
//
// Example usage:
//
//	t := TermsFunc("category", []string{"books", "electronics"}, func(key string, values []string) bool {
//	    return len(values) > 0
//	})
//	// t is either a termsType object or nil based on the condition.
//
// Parameters:
//   - key: A string representing the field name for the terms query.
//   - values: A slice of values to be searched for in the specified field.
//   - f: A function that takes a key and values, and returns a boolean indicating
//     whether to create the termsType object.
//
// Returns:
//
//	A termsType object if the condition is true; otherwise, nil.
func TermsFunc[T any](key string, values []T, f func(key string, values []T) bool) termsType {
	if !f(key, values) {
		return nil
	}
	return TermsArray(key, values)
}

// TermsIf creates a termsType object based on a boolean condition.
//
// This function creates a termsType object if the provided condition is true. If
// the condition is false, it returns nil instead of creating a termsType object.
//
// Example usage:
//
//	t := TermsIf("category", []string{"books", "electronics"}, true)
//	// t is a termsType object if the condition is true; otherwise, it is nil.
//
// Parameters:
//   - key: A string representing the field name for the terms query.
//   - values: A slice of values to be searched for in the specified field.
//   - condition: A boolean value that determines whether to create the termsType object.
//
// Returns:
//
//	A termsType object if the condition is true; otherwise, nil.
func TermsIf[T any](key string, values []T, condition bool) termsType {
	if !condition {
		return nil
	}
	return TermsArray(key, values)
}

// Exists creates a new existsType object to check if a field exists.
//
// This function initializes an existsType object that specifies a query to check
// if a particular field exists in the documents. The key parameter represents
// the name of the field to check for existence.
//
// Example usage:
//
//	e := Exists("title")
//	// e now contains an existsType object that checks for the existence of the "title" field.
//
// Parameters:
//   - key: A string representing the name of the field to check for existence.
//
// Returns:
//
//	An existsType object that includes the "exists" query for the specified field.
func Exists(key string) existsType {
	return existsType{
		"exists": Object{
			"field": key,
		},
	}
}

// ExistsFunc creates an existsType object based on a condition evaluated by a function.
//
// This function conditionally creates an existsType object if the provided function
// returns true for the given key. If the function returns false, it returns nil
// instead of creating an existsType object.
//
// Example usage:
//
//	e := ExistsFunc("title", func(key string) bool {
//	    return key != ""
//	})
//	// e is either an existsType object or nil based on the condition.
//
// Parameters:
//   - key: A string representing the name of the field to check for existence.
//   - f: A function that takes a key and returns a boolean indicating whether
//     to create the existsType object.
//
// Returns:
//
//	An existsType object if the condition is true; otherwise, nil.
func ExistsFunc(key string, f func(key string) bool) existsType {
	if !f(key) {
		return nil
	}
	return Exists(key)
}

// ExistsIf creates an existsType object based on a boolean condition.
//
// This function creates an existsType object if the provided condition is true.
// If the condition is false, it returns nil instead of creating an existsType object.
//
// Example usage:
//
//	e := ExistsIf("title", true)
//	// e is an existsType object if the condition is true; otherwise, it is nil.
//
// Parameters:
//   - key: A string representing the name of the field to check for existence.
//   - condition: A boolean value that determines whether to create the existsType object.
//
// Returns:
//
//	An existsType object if the condition is true; otherwise, nil.
func ExistsIf(key string, condition bool) existsType {
	if !condition {
		return nil
	}
	return Exists(key)
}

// Match creates a new matchType object with the specified field and query.
//
// This function initializes a matchType object for a match query, where the key
// is the field name and query is the value to search for in that field. This is used
// to construct queries that match the specified value in the given field.
//
// Example usage:
//
//	m := Match("title", "es-query-builder")
//	// m now contains a matchType object that matches the query "es-query-builder" in the "title" field.
//
// Parameters:
//   - key: A string representing the field name for the match query.
//   - query: The value to be matched in the specified field. The type is generic.
//
// Returns:
//
//	A matchType object containing the specified match query.
func Match[T any](key string, query T) matchType {
	return matchType{
		"match": Object{
			key: Object{
				"query": query,
			},
		},
	}
}

func (m matchType) putInTheField(key string, value any) matchType {
	if match, exists := m["match"]; exists {
		if matchObject, moOk := match.(Object); moOk {
			for field := range matchObject {
				if fieldObject, foOk := matchObject[field].(Object); foOk {
					fieldObject[key] = value
				}
			}
		}
	}
	return m
}

// Operator sets the "operator" field in the match query.
//
// This method configures the match query to use a specified operator (e.g., "AND" or "OR")
// for the matching process. It calls putInTheField to add or update the "operator" key
// in the match query object.
//
// Example usage:
//
//	m := Match("title", "es-query-builder").Operator("AND")
//	// m now has an "operator" field set to "AND" in the match query object.
//
// Parameters:
//   - operator: An Operator.Operator value representing the logical operator to be used in the match query.
//
// Returns:
//
//	The updated matchType object with the "operator" field set to the specified value.
func (m matchType) Operator(operator Operator.Operator) matchType {
	return m.putInTheField("operator", operator)
}

// Boost sets the "boost" field in the match query.
//
// This method configures the match query to use a specified boost factor, which influences
// the relevance scoring of the matched documents. It calls putInTheField to add or update
// the "boost" key in the match query object.
//
// Example usage:
//
//	m := Match("title", "es-query-builder").Boost(1.5)
//	// m now has a "boost" field set to 1.5 in the match query object.
//
// Parameters:
//   - boost: A float64 value representing the boost factor to be applied to the match query.
//
// Returns:
//
//	The updated matchType object with the "boost" field set to the specified value.
func (m matchType) Boost(boost float64) matchType {
	return m.putInTheField("boost", boost)
}

// MatchNone creates a new matchNoneType object with the specified field and query.
//
// This function initializes a matchNoneType object for a match_none query, where the key
// represents the field name and query is the value to be matched. This is used to construct
// queries that explicitly match no documents for the specified value in the given field.
//
// Example usage:
//
//	mn := MatchNone("title", "es-query-builder")
//	// mn now contains a matchNoneType object that matches no documents for the "title" field with the query "es-query-builder".
//
// Parameters:
//   - key: A string representing the field name for the match_none query.
//   - query: The value to be used in the match_none query. The type is generic.
//
// Returns:
//
//	A matchNoneType object containing the specified match_none query.
func MatchNone[T any](key string, query T) matchNoneType {
	return matchNoneType{
		"match_none": Object{
			key: Object{
				"query": query,
			},
		},
	}
}

func (m matchNoneType) putInTheField(key string, value any) matchNoneType {
	if matchNone, exists := m["match_none"]; exists {
		if matchNoneObject, mnoOk := matchNone.(Object); mnoOk {
			for field := range matchNoneObject {
				if fieldObject, foOk := matchNoneObject[field].(Object); foOk {
					fieldObject[key] = value
				}
			}
		}
	}
	return m
}

// Operator sets the "operator" field in the match_none query.
//
// This method configures the match_none query to use a specified operator (e.g., "AND" or "OR")
// for the matching process. It calls putInTheField to add or update the "operator" key
// in the match_none query object.
//
// Example usage:
//
//	mn := MatchNone("title", "es-query-builder").Operator("AND")
//	// mn now has an "operator" field set to "AND" in the match_none query object.
//
// Parameters:
//   - operator: An Operator.Operator value representing the logical operator to be used in the match_none query.
//
// Returns:
//
//	The updated matchNoneType object with the "operator" field set to the specified value.
func (m matchNoneType) Operator(operator Operator.Operator) matchNoneType {
	return m.putInTheField("operator", operator)
}

// Boost sets the "boost" field in the match_none query.
//
// This method configures the match_none query to use a specified boost factor, which influences
// the relevance scoring of the matched documents. It calls putInTheField to add or update
// the "boost" key in the match_none query object.
//
// Example usage:
//
//	mn := MatchNone("title", "es-query-builder").Boost(1.5)
//	// mn now has a "boost" field set to 1.5 in the match_none query object.
//
// Parameters:
//   - boost: A float64 value representing the boost factor to be applied to the match_none query.
//
// Returns:
//
//	The updated matchNoneType object with the "boost" field set to the specified value.
func (m matchNoneType) Boost(boost float64) matchNoneType {
	return m.putInTheField("boost", boost)
}

// MatchAll creates a new matchAllType object that matches all documents.
//
// This function initializes a matchAllType object that is used to construct queries
// that match all documents. This can be useful for scenarios where you want to ensure
// that all documents are included in the results.
//
// Example usage:
//
//	ma := MatchAll()
//	// ma now contains a matchAllType object that matches all documents.
//
// Returns:
//
//	A matchAllType object with a match_all query that matches all documents.
func MatchAll() matchAllType {
	return matchAllType{
		"match_all": Object{},
	}
}

func (m matchAllType) putInTheField(key string, value any) matchAllType {
	if matchAll, exists := m["match_all"]; exists {
		matchAll.(Object)[key] = value
	}
	return m
}

// Boost sets the "boost" field in the match_all query.
//
// This method configures the match_all query to use a specified boost factor, which influences
// the relevance scoring of the matched documents. It calls putInTheField to add or update
// the "boost" key in the match_all query object.
//
// Example usage:
//
//	ma := MatchAll().Boost(1.5)
//	// ma now has a "boost" field set to 1.5 in the match_all query object.
//
// Parameters:
//   - boost: A float64 value representing the boost factor to be applied to the match_all query.
//
// Returns:
//
//	The updated matchAllType object with the "boost" field set to the specified value.
func (m matchAllType) Boost(boost float64) matchAllType {
	return m.putInTheField("boost", boost)
}

// Range creates a new rangeType object with the specified field.
//
// This function initializes a rangeType object for specifying range queries. The key represents
// the field name, and the rangeType object is used to define the range conditions for that field.
//
// Example usage:
//
//	r := Range("age")
//	// r now contains a rangeType object with the specified field "age" for range queries.
//
// Parameters:
//   - key: A string representing the field name for the range query.
//
// Returns:
//
//	A rangeType object with the specified field ready for defining range conditions.
func Range(key string) rangeType {
	return rangeType{
		key: Object{},
	}
}

// Range adds a range query to the es.Object with the specified field.
//
// This method initializes a range query for the specified field using the es.Range function,
// and adds it to the "range" section of the es.Object. It allows the es.Object to include
// range-based filtering or searching for the specified field.
//
// Example usage:
//
//	o := es.NewQuery(...).Range("age")
//	// o now contains a rangeType object in the "range" section for the "age" field.
//
// Parameters:
//   - key: A string representing the field name for the range query.
//
// Returns:
//
//	The updated rangeType object that is included in the "range" section of the Object.
func (o Object) Range(key string) rangeType {
	r := Range(key)
	o.putInQuery("range", r)
	return r
}

// LesserThan sets the "lt" (less than) field for the range query.
//
// This method specifies that the range query should match values that are less than
// the provided value. It removes any existing "lte" (less than or equal to) field to ensure
// that only one type of range condition is applied.
//
// Example usage:
//
//	r := Range("age").LesserThan(20)
//	// r now has an "lt" field set to 20 in the range query for the "age" field.
//
// Parameters:
//   - lt: The value that the field should be less than.
//
// Returns:
//
//	The updated rangeType object with the "lt" field set to the specified value.
func (r rangeType) LesserThan(lt any) rangeType {
	for key := range r {
		if rangeObject, ok := r[key].(Object); ok {
			rangeObject["lt"] = lt
			delete(rangeObject, "lte")
		}
	}
	return r
}

// LesserThanOrEqual sets the "lte" (less than or equal to) field for the range query.
//
// This method specifies that the range query should match values that are less than or equal
// to the provided value. It removes any existing "lt" (less than) field to ensure that only
// one type of range condition is applied.
//
// Example usage:
//
//	r := Range("age").LesserThanOrEqual(20)
//	// r now has an "lte" field set to 20 in the range query for the "age" field.
//
// Parameters:
//   - lte: The value that the field should be less than or equal to.
//
// Returns:
//
//	The updated rangeType object with the "lte" field set to the specified value.
func (r rangeType) LesserThanOrEqual(lte any) rangeType {
	for key := range r {
		if rangeObject, ok := r[key].(Object); ok {
			rangeObject["lte"] = lte
			delete(rangeObject, "lt")
		}
	}
	return r
}

// GreaterThan sets the "gt" (greater than) field for the range query.
//
// This method specifies that the range query should match values that are greater than
// the provided value. It removes any existing "gte" (greater than or equal to) field
// to ensure that only one type of range condition is applied.
//
// Example usage:
//
//	r := Range("age").GreaterThan(50)
//	// r now has a "gt" field set to 50 in the range query for the "age" field.
//
// Parameters:
//   - gt: The value that the field should be greater than.
//
// Returns:
//
//	The updated rangeType object with the "gt" field set to the specified value.
func (r rangeType) GreaterThan(gt any) rangeType {
	for key := range r {
		if rangeObject, ok := r[key].(Object); ok {
			rangeObject["gt"] = gt
			delete(rangeObject, "gte")
		}
	}
	return r
}

// GreaterThanOrEqual sets the "gte" (greater than or equal to) field for the range query.
//
// This method specifies that the range query should match values that are greater than or equal
// to the provided value. It removes any existing "gt" (greater than) field to ensure that only
// one type of range condition is applied.
//
// Example usage:
//
//	r := Range("age").GreaterThanOrEqual(50)
//	// r now has a "gte" field set to 50 in the range query for the "age" field.
//
// Parameters:
//   - gte: The value that the field should be greater than or equal to.
//
// Returns:
//
//	The updated rangeType object with the "gte" field set to the specified value.
func (r rangeType) GreaterThanOrEqual(gte any) rangeType {
	for key := range r {
		if rangeObject, ok := r[key].(Object); ok {
			rangeObject["gte"] = gte
			delete(rangeObject, "gt")
		}
	}
	return r
}

// Format sets the "format" field for the range query.
//
// This method specifies a format for the range query values, which can be useful for
// controlling how date or numeric values are interpreted. It applies the format to
// all fields in the range query object.
//
// Example usage:
//
//	r := Range("date").Format("yyyy-MM-dd")
//	// r now has a "format" field set to "yyyy-MM-dd" in the range query for the "date" field.
//
// Parameters:
//   - format: A string representing the format to be applied to the range query values.
//
// Returns:
//
//	The updated rangeType object with the "format" field set to the specified value.
func (r rangeType) Format(format string) rangeType {
	for key := range r {
		if rangeObject, ok := r[key].(Object); ok {
			rangeObject["format"] = format
		}
	}
	return r
}

// Boost sets the "boost" field for the range query.
//
// This method applies a boost factor to the range query, influencing the relevance scoring
// of documents that match the query. It applies the boost to all fields in the range query object.
//
// Example usage:
//
//	r := Range("age").Boost(1.5)
//	// r now has a "boost" field set to 1.5 in the range query for the "age" field.
//
// Parameters:
//   - boost: A float64 value representing the boost factor to be applied to the range query.
//
// Returns:
//
//	The updated rangeType object with the "boost" field set to the specified value.
func (r rangeType) Boost(boost float64) rangeType {
	for key := range r {
		if rangeObject, ok := r[key].(Object); ok {
			rangeObject["boost"] = boost
		}
	}
	return r
}

// Nested creates a new nestedType object for a nested query.
//
// This function initializes a nested query object with the specified path and query.
// The path represents the field path for the nested query, and the nested query is
// constructed using the es.NewQuery function.
//
// Example usage:
//
//	nestedQuery := Nested("comments", es.Bool().Filter(...).MustNot(...))
//	// nestedQuery now contains a nestedType object with the specified path and query.
//
// Parameters:
//   - path: A string representing the path for the nested query.
//   - nestedQuery: The query to be applied within the nested query.
//
// Returns:
//
//	A nestedType object with the "nested" query and specified path.
func Nested[T any](path string, nestedQuery T) nestedType {
	o := NewQuery(nestedQuery)
	o["path"] = path
	return nestedType{
		"nested": o,
	}
}

func (n nestedType) putInNested(key string, value any) nestedType {
	if nested, exists := n["nested"]; exists {
		nested.(Object)[key] = value
	}
	return n
}

// InnerHits sets the "inner_hits" field for the nested query.
//
// This method specifies the inner hits for the nested query, allowing you to control
// the returned inner documents for the nested query. It uses the putInNested method
// to add or update the "inner_hits" key in the nested query object.
//
// Example usage:
//
//	nested := Nested("comments", es.Term("text", "example"))
//	nested = nested.InnerHits(Object{"inner": "hits"})
//	// nested now has an "inner_hits" field with the specified Object in the nested query.
//
// Parameters:
//   - innerHits: An es.Object representing the inner hits configuration for the nested query.
//
// Returns:
//
//	The updated nestedType object with the "inner_hits" field set to the specified value.
func (n nestedType) InnerHits(innerHits Object) nestedType {
	return n.putInNested("inner_hits", innerHits)
}

// ScoreMode sets the "score_mode" field for the nested query.
//
// This method configures how the scores for the nested documents should be calculated
// and combined. It uses the putInNested method to add or update the "score_mode"
// key in the nested query object.
//
// Example usage:
//
//	nested := Nested("comments", es.Term("text", "example"))
//	nested = nested.ScoreMode(ScoreMode.Sum)
//	// nested now has a "score_mode" field with the specified ScoreMode value in the nested query.
//
// Parameters:
//   - scoreMode: A ScoreMode.ScoreMode value representing the scoring mode for the nested query.
//
// Returns:
//
//	The updated nestedType object with the "score_mode" field set to the specified value.
func (n nestedType) ScoreMode(scoreMode ScoreMode.ScoreMode) nestedType {
	return n.putInNested("score_mode", scoreMode)
}

// AggTerm creates a new aggregation term with the specified field.
//
// This function initializes an aggregation term with the given field name.
// It can be used to specify a field for aggregation operations in queries.
//
// Example usage:
//
//	termAgg := AggTerm("fieldName")
//	// termAgg now has the "field" set to "fieldName".
//
// Parameters:
//   - field: The name of the field to aggregate on.
//
// Returns:
//
//	An aggTermType object with the "field" set to the provided value.
func AggTerm(field string) aggTermType {
	return aggTermType{
		"field": field,
	}
}

// Missing sets the "missing" value for an aggregation term.
//
// This method specifies a value to be used when the field is missing in documents.
// It updates the aggTermType object to handle missing values in the aggregation.
//
// Example usage:
//
//	termAgg := AggTerm("fieldName").Missing("N/A")
//	// termAgg now has the "missing" field set to "N/A".
//
// Parameters:
//   - missing: The value to use when the field is missing.
//
// Returns:
//
//	The updated aggTermType object with the "missing" field set to the specified value.
func (aggTerm aggTermType) Missing(missing string) aggTermType {
	aggTerm["missing"] = missing
	return aggTerm
}

// AggTerms creates a new "terms" aggregation.
//
// This function initializes an aggregation for terms. It can be used to perform
// aggregation based on the unique terms of a field.
//
// Example usage:
//
//	termsAgg := AggTerms()
//	// termsAgg now has the "terms" field initialized.
//
// Returns:
//
//	An aggsType object with the "terms" field initialized.
func AggTerms() aggsType {
	return aggsType{
		"terms": Object{},
	}
}

// AggMultiTerms creates a new "multi_terms" aggregation.
//
// This function initializes an aggregation for multiple terms. It can be used
// to perform aggregation based on multiple fields or term combinations.
//
// Example usage:
//
//	multiTermsAgg := AggMultiTerms()
//	// multiTermsAgg now has the "multi_terms" field initialized.
//
// Returns:
//
//	An aggsType object with the "multi_terms" field initialized.
func AggMultiTerms() aggsType {
	return aggsType{
		"multi_terms": Object{},
	}
}

// AggMax creates a new "max" aggregation.
//
// This function initializes an aggregation to calculate the maximum value of a field.
//
// Example usage:
//
//	maxAgg := AggMax()
//	// maxAgg now has the "max" field initialized.
//
// Returns:
//
//	An aggsType object with the "max" field initialized.
func AggMax() aggsType {
	return aggsType{
		"max": Object{},
	}
}

// AggMin creates a new "min" aggregation.
//
// This function initializes an aggregation to calculate the minimum value of a field.
//
// Example usage:
//
//	minAgg := AggMin()
//	// minAgg now has the "min" field initialized.
//
// Returns:
//
//	An aggsType object with the "min" field initialized.
func AggMin() aggsType {
	return aggsType{
		"min": Object{},
	}
}

// AggAvg creates a new "avg" aggregation.
//
// This function initializes an aggregation to calculate the average value of a field.
//
// Example usage:
//
//	avgAgg := AggAvg()
//	// avgAgg now has the "avg" field initialized.
//
// Returns:
//
//	An aggsType object with the "avg" field initialized.
func AggAvg() aggsType {
	return aggsType{
		"avg": Object{},
	}
}

// AggCustom creates a custom aggregation with the provided aggregation object.
//
// This function initializes an aggregation based on the given custom aggregation definition.
//
// Example usage:
//
//	customAgg := AggCustom(Object{"custom": "value"})
//	// customAgg now has the custom aggregation specified.
//
// Parameters:
//   - agg: An es.Object representing a custom aggregation definition.
//
// Returns:
//
//	An aggsType object initialized with the provided custom aggregation.
func AggCustom(agg Object) aggsType {
	return aggsType(agg)
}

func (agg aggsType) putInTheField(key string, value any) aggsType {
	for field := range agg {
		if fieldObject, ok := agg[field].(Object); ok {
			fieldObject[key] = value
		}
	}
	return agg
}

// Aggs adds a nested aggregation to the aggsType object.
//
// This method adds a nested aggregation under the "aggs" field with the given name.
//
// Example usage:
//
//	nestedAgg := AggTerms().Size(5)
//	agg := AggTerms().Aggs("nested", nestedAgg)
//	// agg now has a nested aggregation named "nested" with the specified aggregation.
//
// Parameters:
//   - name: The name of the nested aggregation.
//   - nestedAgg: The nested aggregation to add.
//
// Returns:
//
//	The updated aggsType object with the nested aggregation added.
func (agg aggsType) Aggs(name string, nestedAgg aggsType) aggsType {
	agg["aggs"] = Object{
		name: nestedAgg,
	}
	return agg
}

// Field sets the "field" value in the aggsType object.
//
// This method specifies the field to aggregate on in the aggsType object.
//
// Example usage:
//
//	agg := AggTerms().Field("fieldName")
//	// agg now has the "field" set to "fieldName".
//
// Parameters:
//   - field: The name of the field to aggregate on.
//
// Returns:
//
//	The updated aggsType object with the "field" set to the specified value.
func (agg aggsType) Field(field string) aggsType {
	return agg.putInTheField("field", field)
}

// Size sets the "size" value in the aggsType object.
//
// This method specifies the number of terms to return in the aggregation result.
//
// Example usage:
//
//	agg := AggTerms().Size(10)
//	// agg now has the "size" field set to 10.
//
// Parameters:
//   - size: The number of terms to return.
//
// Returns:
//
//	The updated aggsType object with the "size" field set to the specified value.
func (agg aggsType) Size(size int) aggsType {
	return agg.putInTheField("size", size)
}

// Order sets the "order" field in the aggsType object.
//
// This method specifies the sorting order for the aggregation results.
//
// Example usage:
//
//	agg := AggTerms().Order("fieldName", Order.Desc)
//	// agg now has the "order" field set to "desc" for "fieldName".
//
// Parameters:
//   - field: The name of the field to sort by.
//   - order: The Order value specifying the sorting direction (e.g., Asc or Desc).
//
// Returns:
//
//	The updated aggsType object with the "order" field set to the specified value.
func (agg aggsType) Order(field string, order Order.Order) aggsType {
	return agg.putInTheField("order",
		Object{
			field: order,
		},
	)
}

// Include sets the "include" field in the aggsType object.
//
// This method specifies a pattern to include in the aggregation results.
//
// Example usage:
//
//	agg := AggTerms().Include("pattern*")
//	// agg now has the "include" field set to "pattern*".
//
// Parameters:
//   - include: The pattern to include in the aggregation results.
//
// Returns:
//
//	The updated aggsType object with the "include" field set to the specified value.
func (agg aggsType) Include(include string) aggsType {
	return agg.putInTheField("include", include)
}

// Exclude sets the "exclude" field in the aggsType object.
//
// This method specifies a pattern to exclude from the aggregation results.
//
// Example usage:
//
//	agg := AggTerms().Exclude("pattern*")
//	// agg now has the "exclude" field set to "pattern*".
//
// Parameters:
//   - exclude: The pattern to exclude from the aggregation results.
//
// Returns:
//
//	The updated aggsType object with the "exclude" field set to the specified value.
func (agg aggsType) Exclude(exclude string) aggsType {
	return agg.putInTheField("exclude", exclude)
}

// Terms sets the "terms" field in the aggsType object.
//
// This method adds a list of aggregation terms to the "terms" field of the aggsType object.
// It allows specifying multiple term aggregations for the aggregation query.
//
// Example usage:
//
//	agg := AggTerms().
//		Terms(
//			AggTerm("field1"),
//			AggTerm("field2"),
//		)
//	// agg now has the "terms" field containing the provided term aggregations.
//
// Parameters:
//   - terms: A variadic list of aggTermType objects representing the term aggregations.
//
// Returns:
//
//	The updated aggsType object with the "terms" field set to the provided term aggregations.
func (agg aggsType) Terms(terms ...aggTermType) aggsType {
	return agg.putInTheField("terms", terms)
}

// Aggs adds a named aggregation to the "aggs" field of the es.Object.
//
// This method allows adding a nested aggregation under the "aggs" field in the es.Object.
// It associates the given name with the specified aggregation, enabling complex aggregation queries.
//
// Example usage:
//
//	termAgg := AggTerms().Field("fieldName")
//	query := es.NewQuery().Aggs("myAgg", termAgg)
//	// query now has an "aggs" field with a nested aggregation named "myAgg".
//
// Parameters:
//   - name: The name to associate with the nested aggregation.
//   - agg: The aggsType object representing the nested aggregation.
//
// Returns:
//
//	The updated Object with the "aggs" field containing the new named aggregation.
func (o Object) Aggs(name string, agg aggsType) Object {
	aggs, exists := o["aggs"]
	if !exists {
		aggs = Object{}
	}
	aggs.(Object)[name] = agg
	o["aggs"] = aggs
	return o
}
