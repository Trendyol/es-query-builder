package es

import (
	Mode "github.com/Trendyol/es-query-builder/es/enums/sort/mode"
	Order "github.com/Trendyol/es-query-builder/es/enums/sort/order"
)

type sortType Object

type sourceType Object

type includesType Array

type excludesType Array

// NewQuery creates a new query es.Object with the provided query clause.
//
// This function takes any query clause as input and attempts to convert it into the correct internal type using the `correctType` function.
// If the conversion is successful, the resulting field is stored under the "query" key in the returned es.Object.
// If the conversion fails or the input is nil, an empty es.Object is returned under the "query" key.
//
// Example usage:
//
//	termQuery := es.Term("fieldName", "value")
//	query := es.NewQuery(termQuery)
//	// query now contains a "query" field with the term query.
//
// Parameters:
//   - queryClause: The query clause to be converted and added to the "query" field. It can be of any type.
//
// Returns:
//
//	An Object containing the "query" field with the processed query clause.
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

// Sort creates a new sortType object with the specified field.
//
// This function initializes a sortType object with a given field name. The
// field is used to specify the sorting criteria in the search query. The
// resulting sortType can be further configured with sorting order and mode.
//
// Example usage:
//
//	s := Sort("age")
//	// s now includes a sortType with an "age" field that can be further configured.
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