package es

import (
	Mode "github.com/Trendyol/es-query-builder/es/enums/sort/mode"
	Order "github.com/Trendyol/es-query-builder/es/enums/sort/order"
)

type sortType Object

// Sort creates a new es.sortType object with the specified field.
//
// This function initializes an es.sortType object with a given field name. The
// field is used to specify the sorting criteria in the search query. The
// resulting es.sortType can be further configured with sorting order and mode.
//
// Example usage:
//
//	s := es.Sort("age")
//	// s now includes an es.sortType with an "age" field that can be further configured.
//
// Parameters:
//   - field: A string representing the field to sort by.
//
// Returns:
//
//	An es.sortType object with the specified field.
func Sort(field string) sortType {
	return sortType{
		field: Object{},
	}
}

// Order sets the "order" parameter in an es.sortType object.
//
// This method specifies the order in which the results should be sorted.
// It configures the es.sortType object to sort the results in ascending or
// descending order.
//
// Example usage:
//
//	s := es.Sort("age").Order(Order.Desc)
//	// s now includes an "order" parameter with the value "desc".
//
// Parameters:
//   - order: An Order.Order value indicating the sorting order (e.g., ascending or descending).
//
// Returns:
//
//	The updated es.sortType object with the "order" parameter set.
func (s sortType) Order(order Order.Order) sortType {
	return s.putInTheField("order", order)
}

// Mode sets the "mode" parameter in an es.sortType object.
//
// This method specifies the mode used for sorting the results. The mode
// determines how sorting should be handled, such as by specifying different
// tie-breaking strategies.
//
// Example usage:
//
//	s := es.Sort("age").Mode(Mode.Avg)
//	// s now includes a "mode" parameter with the value "avg".
//
// Parameters:
//   - mode: A Mode.Mode value indicating the sorting mode (e.g., average, minimum, maximum).
//
// Returns:
//
//	The updated es.sortType object with the "mode" parameter set.
func (s sortType) Mode(mode Mode.Mode) sortType {
	return s.putInTheField("mode", mode)
}

// Nested sets the "nested" parameter in an es.sortType object.
//
// This method specifies a nested sorting configuration for sorting fields
// within nested objects. It allows defining sorting behavior for fields
// inside nested documents.
//
// Example usage:
//
//	s := es.Sort("user.age").Nested(es.NestedSort().Path("user"))
//	// s now includes a "nested" parameter with the specified nested sorting configuration.
//
// Parameters:
//   - nested: A nestedSortType value defining the nested sorting configuration.
//
// Returns:
//
//	The updated es.sortType object with the "nested" parameter set.
func (s sortType) Nested(nested nestedSortType) sortType {
	return s.putInTheField("nested", nested)
}

func (s sortType) putInTheField(key string, value any) sortType {
	return genericPutInTheFieldOfFirstObject(s, key, value)
}
