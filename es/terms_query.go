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
