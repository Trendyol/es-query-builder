package es

type BoolType Object

type FilterType Array

type MustType Array

type MustNotType Array

type ShouldType Array

// Bool creates and returns an empty BoolType object.
//
// This function is typically used to initialize a BoolType, which can be
// populated later with the appropriate boolean query conditions.
//
// Example usage:
//
//	b := Bool()
//	// b is now an empty BoolType object that can be used in a query.
//
// Returns:
//
//	An empty BoolType object.
func Bool() BoolType {
	return BoolType{}
}

// MinimumShouldMatch sets the "minimum_should_match" parameter in a BoolType query.
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
//	The updated BoolType object with the "minimum_should_match" parameter set.
func (b BoolType) MinimumShouldMatch(minimumShouldMatch int) BoolType {
	b["minimum_should_match"] = minimumShouldMatch
	return b
}

// AdjustPureNegative sets the "adjust_pure_negative" parameter in a BoolType query.
//
// This method allows you to specify whether pure negative queries should be
// adjusted or not. When set to true, the query will be adjusted to include
// pure negative queries, which can influence the matching behavior of the
// boolean query.
//
// Example usage:
//
//	b := Bool().AdjustPureNegative(true)
//	// b now includes an "adjust_pure_negative" parameter set to true.
//
// Parameters:
//   - adjustPureNegative: A boolean value indicating whether to adjust
//     pure negative queries in the boolean query.
//
// Returns:
//
//	The updated BoolType object with the "adjust_pure_negative" parameter set.
func (b BoolType) AdjustPureNegative(adjustPureNegative bool) BoolType {
	b["adjust_pure_negative"] = adjustPureNegative
	return b
}

// Boost sets the "boost" parameter in a BoolType query.
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
//	The updated BoolType object with the "boost" parameter set.
func (b BoolType) Boost(boost float64) BoolType {
	b["boost"] = boost
	return b
}

// Filter adds one or more filter conditions to the BoolType object.
//
// This method updates the "filter" section of the BoolType object by appending
// the specified filter conditions. It accepts a variadic number of filter conditions,
// checks their types, and adds them to the "filter" array in the BoolType object.
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
//	The updated BoolType object with the new filter conditions added.
func (b BoolType) Filter(items ...any) BoolType {
	filter, ok := b["filter"].(FilterType)
	if !ok {
		filter = FilterType{}
	}
	for i := 0; i < len(items); i++ {
		if field, fOk := correctType(items[i]); fOk {
			filter = append(filter, field)
		}
	}
	b["filter"] = filter
	return b
}

// Must adds one or more conditions to the "must" section of the BoolType object.
//
// This method updates the "must" section by appending the specified conditions.
// It accepts a variadic number of conditions, checks their types, and adds them to
// the "must" array in the BoolType object.
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
//	The updated BoolType object with the new conditions added to the "must" section.
func (b BoolType) Must(items ...any) BoolType {
	must, ok := b["must"].(MustType)
	if !ok {
		must = MustType{}
	}
	for i := 0; i < len(items); i++ {
		if field, fOk := correctType(items[i]); fOk {
			must = append(must, field)
		}
	}
	b["must"] = must
	return b
}

// MustNot adds one or more conditions to the "must_not" section of the BoolType object.
//
// This method updates the "must_not" section by appending the specified conditions.
// It accepts a variadic number of conditions, checks their types, and adds them to
// the "must_not" array in the BoolType object.
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
//	The updated BoolType object with the new conditions added to the "must_not" section.
func (b BoolType) MustNot(items ...any) BoolType {
	mustNot, ok := b["must_not"].(MustNotType)
	if !ok {
		mustNot = MustNotType{}
	}
	for i := 0; i < len(items); i++ {
		if field, fOk := correctType(items[i]); fOk {
			mustNot = append(mustNot, field)
		}
	}
	b["must_not"] = mustNot
	return b
}

// Should adds one or more conditions to the "should" section of the BoolType object.
//
// This method updates the "should" section by appending the specified conditions.
// It accepts a variadic number of conditions, checks their types, and adds them to
// the "should" array in the BoolType object.
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
//	The updated BoolType object with the new conditions added to the "should" section.
func (b BoolType) Should(items ...any) BoolType {
	should, ok := b["should"].(ShouldType)
	if !ok {
		should = ShouldType{}
	}
	for i := 0; i < len(items); i++ {
		if field, fOk := correctType(items[i]); fOk {
			should = append(should, field)
		}
	}
	b["should"] = should
	return b
}
