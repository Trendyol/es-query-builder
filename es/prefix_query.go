package es

type prefixType Object

// Prefix creates a new es.prefixType object with the specified key-value pair.
//
// This function initializes an es.prefixType object with a single prefix query, where the
// key is the field name and the value is the prefix to search for.
//
// Example usage:
//
//	p := es.Prefix("user.id", "ki")
//	// p now contains an es.prefixType object with a prefix query for the "user.id" field.
//
// Parameters:
//   - key: A string representing the field name for the prefix query.
//   - value: The prefix value to be searched for in the specified field.
//
// Returns:
//
//	An es.prefixType object containing the specified prefix query.
func Prefix(key string, value string) prefixType {
	return prefixType{
		"prefix": Object{
			key: Object{
				"value": value,
			},
		},
	}
}

// CaseInsensitive sets the "case_insensitive" parameter in an es.prefixType query.
//
// Example usage:
//
//	p := es.Prefix("user.id", "Ki").CaseInsensitive(true)
//
// Parameters:
//   - caseInsensitive: A boolean value indicating whether the prefix query
//     should be case-insensitive.
//
// Returns:
//
//	The updated es.prefixType object with the "case_insensitive" parameter set.
func (p prefixType) CaseInsensitive(caseInsensitive bool) prefixType {
	return p.putInTheField("case_insensitive", caseInsensitive)
}

// Rewrite sets the "rewrite" parameter in an es.prefixType query.
//
// Example usage:
//
//	p := es.Prefix("user.id", "ki").Rewrite("constant_score")
//
// Parameters:
//   - rewrite: A string value representing the rewrite method to be applied.
//
// Returns:
//
//	The updated es.prefixType object with the "rewrite" parameter set.
func (p prefixType) Rewrite(rewrite string) prefixType {
	return p.putInTheField("rewrite", rewrite)
}

// Boost sets the "boost" parameter in an es.prefixType query.
//
// Example usage:
//
//	p := es.Prefix("user.id", "ki").Boost(1.5)
//
// Parameters:
//   - boost: A float64 value representing the boost factor for the prefix query.
//
// Returns:
//
//	The updated es.prefixType object with the "boost" parameter set.
func (p prefixType) Boost(boost float64) prefixType {
	return p.putInTheField("boost", boost)
}

// PrefixFunc creates an es.prefixType object based on a condition evaluated by a function.
//
// Example usage:
//
//	p := es.PrefixFunc("user.id", "ki", func(key, value string) bool {
//	    return value != ""
//	})
//
// Parameters:
//   - key: A string representing the field name for the prefix query.
//   - value: The prefix value to be searched for in the specified field.
//   - f: A function that takes a key and value, and returns a boolean indicating
//     whether to create the es.prefixType object.
//
// Returns:
//
//	An es.prefixType object if the condition is true; otherwise, nil.
func PrefixFunc(key string, value string, f func(key string, value string) bool) prefixType {
	if !f(key, value) {
		return nil
	}
	return Prefix(key, value)
}

// PrefixIf creates an es.prefixType object based on a boolean condition.
//
// Example usage:
//
//	p := es.PrefixIf("user.id", "ki", true)
//
// Parameters:
//   - key: A string representing the field name for the prefix query.
//   - value: The prefix value to be searched for in the specified field.
//   - condition: A boolean value that determines whether to create the es.prefixType object.
//
// Returns:
//
//	An es.prefixType object if the condition is true; otherwise, nil.
func PrefixIf(key string, value string, condition bool) prefixType {
	if !condition {
		return nil
	}
	return Prefix(key, value)
}

func (p prefixType) putInTheField(key string, value any) prefixType {
	return genericPutInTheFieldOfFirstChild(p, "prefix", key, value)
}
