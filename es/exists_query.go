package es

type existsType Object

// Exists creates a new existsType object to check if a field exists.
//
// This function initializes an existsType object that specifies a query to check
// if a particular field exists in the documents. The key parameter represents
// the name of the field to check for existence.
//
// Example usage:
//
//	e := es.Exists("title")
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
//	e := es.ExistsFunc("title", func(key string) bool {
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
//	e := es.ExistsIf("title", true)
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
