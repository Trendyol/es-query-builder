package es

type wildcardType Object

// Wildcard creates a new es.wildcardType object with the specified key-value pair.
//
// This function initializes an es.wildcardType object with a single wildcard query, where the
// key is the field name and the value is the wildcard pattern to search for (* and ? supported).
//
// Example usage:
//
//	w := es.Wildcard("user.id", "ki*y")
//	// w now contains an es.wildcardType object with a wildcard query for the "user.id" field.
//
// Parameters:
//   - key: A string representing the field name for the wildcard query.
//   - value: The wildcard pattern to be searched for in the specified field.
//
// Returns:
//
//	An es.wildcardType object containing the specified wildcard query.
func Wildcard(key string, value string) wildcardType {
	return wildcardType{
		"wildcard": Object{
			key: Object{
				"value": value,
			},
		},
	}
}

// CaseInsensitive sets the "case_insensitive" parameter in an es.wildcardType query.
//
// Example usage:
//
//	w := es.Wildcard("user.id", "Ki*Y").CaseInsensitive(true)
//
// Parameters:
//   - caseInsensitive: A boolean value indicating whether the wildcard query
//     should be case-insensitive.
//
// Returns:
//
//	The updated es.wildcardType object with the "case_insensitive" parameter set.
func (w wildcardType) CaseInsensitive(caseInsensitive bool) wildcardType {
	return w.putInTheField("case_insensitive", caseInsensitive)
}

// Rewrite sets the "rewrite" parameter in an es.wildcardType query.
//
// Example usage:
//
//	w := es.Wildcard("user.id", "ki*y").Rewrite("constant_score")
//
// Parameters:
//   - rewrite: A string value representing the rewrite method to be applied.
//
// Returns:
//
//	The updated es.wildcardType object with the "rewrite" parameter set.
func (w wildcardType) Rewrite(rewrite string) wildcardType {
	return w.putInTheField("rewrite", rewrite)
}

// Boost sets the "boost" parameter in an es.wildcardType query.
//
// Example usage:
//
//	w := es.Wildcard("user.id", "ki*y").Boost(1.5)
//
// Parameters:
//   - boost: A float64 value representing the boost factor for the wildcard query.
//
// Returns:
//
//	The updated es.wildcardType object with the "boost" parameter set.
func (w wildcardType) Boost(boost float64) wildcardType {
	return w.putInTheField("boost", boost)
}

// WildcardFunc creates an es.wildcardType object based on a condition evaluated by a function.
//
// Example usage:
//
//	w := es.WildcardFunc("user.id", "ki*y", func(key, value string) bool {
//	    return value != ""
//	})
//
// Parameters:
//   - key: A string representing the field name for the wildcard query.
//   - value: The wildcard pattern to be searched for in the specified field.
//   - f: A function that takes a key and value, and returns a boolean indicating
//     whether to create the es.wildcardType object.
//
// Returns:
//
//	An es.wildcardType object if the condition is true; otherwise, nil.
func WildcardFunc(key string, value string, f func(key string, value string) bool) wildcardType {
	if !f(key, value) {
		return nil
	}
	return Wildcard(key, value)
}

// WildcardIf creates an es.wildcardType object based on a boolean condition.
//
// Example usage:
//
//	w := es.WildcardIf("user.id", "ki*y", true)
//
// Parameters:
//   - key: A string representing the field name for the wildcard query.
//   - value: The wildcard pattern to be searched for in the specified field.
//   - condition: A boolean value that determines whether to create the es.wildcardType object.
//
// Returns:
//
//	An es.wildcardType object if the condition is true; otherwise, nil.
func WildcardIf(key string, value string, condition bool) wildcardType {
	if !condition {
		return nil
	}
	return Wildcard(key, value)
}

func (w wildcardType) putInTheField(key string, value any) wildcardType {
	return genericPutInTheFieldOfFirstChild(w, "wildcard", key, value)
}
