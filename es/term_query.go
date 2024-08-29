package es

type termType Object

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
