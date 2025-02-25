package es

type nestedSortType Object

// NestedSort creates a new es.nestedSortType object with the specified path.
//
// This function initializes an es.nestedSortType object, which is used to define
// sorting behavior for fields within nested documents. The specified path determines
// which nested field the sorting applies to.
//
// Example usage:
//
//	ns := es.NestedSort("user")
//	// ns now includes a nestedSortType with the "user" path.
//
// Parameters:
//   - path: A string representing the nested field path to sort by.
//
// Returns:
//
//	An es.nestedSortType object with the specified path.
func NestedSort(path string) nestedSortType {
	return nestedSortType{
		"path": path,
	}
}

// Filter sets the "filter" parameter in an es.nestedSortType object.
//
// This method applies a filtering condition to the nested sorting, ensuring that
// only documents matching the filter criteria are considered for sorting.
//
// Example usage:
//
//	ns := es.NestedSort("user").Filter(es.Term("user.active", true))
//	// ns now includes a "filter" parameter with the specified filter.
//
// Parameters:
//   - filter: A filter object defining the condition for filtering nested documents.
//
// Returns:
//
//	The updated es.nestedSortType object with the "filter" parameter set.
func (ns nestedSortType) Filter(filter any) nestedSortType {
	if field, fOk := correctType(filter); fOk {
		ns["filter"] = field
	}
	return ns
}

// MaxChildren sets the "max_children" parameter in an es.nestedSortType object.
//
// This method specifies the maximum number of child documents that will be
// considered when sorting the parent document. It helps control sorting behavior
// in cases where multiple nested documents exist.
//
// Example usage:
//
//	ns := es.NestedSort("user").MaxChildren(3)
//	// ns now includes a "max_children" parameter with the value 3.
//
// Parameters:
//   - maxChildren: An integer representing the maximum number of child documents considered for sorting.
//
// Returns:
//
//	The updated es.nestedSortType object with the "max_children" parameter set.
func (ns nestedSortType) MaxChildren(maxChildren int) nestedSortType {
	ns["max_children"] = maxChildren
	return ns
}

// Nested sets a nested sorting configuration within an es.nestedSortType object.
//
// This method allows defining nested sorting within another nested field, enabling
// multi-level sorting configurations.
//
// Example usage:
//
//	ns := es.NestedSort("user").Nested(es.NestedSort("user.address"))
//	// ns now includes a "nested" parameter with the specified nested sorting configuration.
//
// Parameters:
//   - nested: A nestedSortType object defining the nested sorting behavior.
//
// Returns:
//
//	The updated es.nestedSortType object with the "nested" parameter set.
func (ns nestedSortType) Nested(nested nestedSortType) nestedSortType {
	ns["nested"] = nested
	return ns
}
