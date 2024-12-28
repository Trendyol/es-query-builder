package es

type termsType Object

// Terms creates a new termsType object with the specified key and values.
//
// This function initializes a termsType object for a terms query, where the key
// is the field name and values is a variadic list of terms to search for in that field.
// This is used to construct queries that match any of the specified terms.
//
// Example usage:
//
//	t := es.Terms("category", "books", "electronics")
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

// Boost sets the "boost" parameter in a termsType query.
//
// This method allows you to specify a boost factor for the terms query,
// which influences the relevance score of documents matching any of the
// specified terms. A higher boost value increases the importance of the
// terms in the query, resulting in higher scores for documents that match
// any of these terms.
//
// Example usage:
//
//	t := es.Terms().Boost(1.5)
//	// t now includes a "boost" parameter set to 1.5.
//
// Parameters:
//   - boost: A float64 value representing the boost factor for the terms
//     query.
//
// Returns:
//
//	The updated termsType object with the "boost" parameter set.
func (t termsType) Boost(boost float64) termsType {
	return t.putInTheField("boost", boost)
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
//	t := es.TermsArray("category", []string{"books", "electronics"})
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
//	t := es.TermsFunc("category", []string{"books", "electronics"}, func(key string, values []string) bool {
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
//	t := es.TermsIf("category", []string{"books", "electronics"}, true)
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

func (t termsType) putInTheField(key string, value any) termsType {
	if terms, ok := t["terms"].(Object); ok {
		terms[key] = value
	}
	return t
}
