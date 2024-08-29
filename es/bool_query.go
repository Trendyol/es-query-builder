package es

type boolType Object

type filterType Array

type mustType Array

type mustNotType Array

type shouldType Array

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
