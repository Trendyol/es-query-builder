package es

type fuzzyType Object

// Fuzzy creates a new es.fuzzyType object with the specified key-value pair.
//
// This function initializes an es.fuzzyType object with a single fuzzy query, where the
// key is the field name and the value is the term to search for with fuzziness.
//
// Example usage:
//
//	f := es.Fuzzy("user", "ki")
//	// f now contains an es.fuzzyType object with a fuzzy query for the "user" field.
//
// Parameters:
//   - key: A string representing the field name for the fuzzy query.
//   - value: The term to be searched for in the specified field.
//
// Returns:
//
//	An es.fuzzyType object containing the specified fuzzy query.
func Fuzzy(key string, value string) fuzzyType {
	return fuzzyType{
		"fuzzy": Object{
			key: Object{
				"value": value,
			},
		},
	}
}

// Fuzziness sets the "fuzziness" parameter in an es.fuzzyType query.
//
// Accepts either a string (e.g. "AUTO") or an integer (0, 1, 2).
//
// Example usage:
//
//	f := es.Fuzzy("user", "ki").Fuzziness("AUTO")
//	f := es.Fuzzy("user", "ki").Fuzziness(1)
//
// Parameters:
//   - fuzziness: The fuzziness value (string or int).
//
// Returns:
//
//	The updated es.fuzzyType object with the "fuzziness" parameter set.
func (f fuzzyType) Fuzziness(fuzziness any) fuzzyType {
	return f.putInTheField("fuzziness", fuzziness)
}

// MaxExpansions sets the "max_expansions" parameter in an es.fuzzyType query.
//
// Example usage:
//
//	f := es.Fuzzy("user", "ki").MaxExpansions(50)
//
// Parameters:
//   - maxExpansions: Maximum number of terms the fuzzy query will expand to.
//
// Returns:
//
//	The updated es.fuzzyType object with the "max_expansions" parameter set.
func (f fuzzyType) MaxExpansions(maxExpansions int) fuzzyType {
	return f.putInTheField("max_expansions", maxExpansions)
}

// PrefixLength sets the "prefix_length" parameter in an es.fuzzyType query.
//
// Example usage:
//
//	f := es.Fuzzy("user", "ki").PrefixLength(0)
//
// Parameters:
//   - prefixLength: Number of beginning characters left unchanged when creating expansions.
//
// Returns:
//
//	The updated es.fuzzyType object with the "prefix_length" parameter set.
func (f fuzzyType) PrefixLength(prefixLength int) fuzzyType {
	return f.putInTheField("prefix_length", prefixLength)
}

// Transpositions sets the "transpositions" parameter in an es.fuzzyType query.
//
// Example usage:
//
//	f := es.Fuzzy("user", "ki").Transpositions(true)
//
// Parameters:
//   - transpositions: Whether to include transpositions (ab → ba) in edits.
//
// Returns:
//
//	The updated es.fuzzyType object with the "transpositions" parameter set.
func (f fuzzyType) Transpositions(transpositions bool) fuzzyType {
	return f.putInTheField("transpositions", transpositions)
}

// Rewrite sets the "rewrite" parameter in an es.fuzzyType query.
//
// Example usage:
//
//	f := es.Fuzzy("user", "ki").Rewrite("constant_score")
//
// Parameters:
//   - rewrite: A string value representing the rewrite method to be applied.
//
// Returns:
//
//	The updated es.fuzzyType object with the "rewrite" parameter set.
func (f fuzzyType) Rewrite(rewrite string) fuzzyType {
	return f.putInTheField("rewrite", rewrite)
}

// CaseInsensitive sets the "case_insensitive" parameter in an es.fuzzyType query.
//
// Example usage:
//
//	f := es.Fuzzy("user", "Ki").CaseInsensitive(true)
//
// Parameters:
//   - caseInsensitive: A boolean value indicating whether the fuzzy query
//     should be case-insensitive.
//
// Returns:
//
//	The updated es.fuzzyType object with the "case_insensitive" parameter set.
func (f fuzzyType) CaseInsensitive(caseInsensitive bool) fuzzyType {
	return f.putInTheField("case_insensitive", caseInsensitive)
}

// Boost sets the "boost" parameter in an es.fuzzyType query.
//
// Example usage:
//
//	f := es.Fuzzy("user", "ki").Boost(1.5)
//
// Parameters:
//   - boost: A float64 value representing the boost factor for the fuzzy query.
//
// Returns:
//
//	The updated es.fuzzyType object with the "boost" parameter set.
func (f fuzzyType) Boost(boost float64) fuzzyType {
	return f.putInTheField("boost", boost)
}

// FuzzyFunc creates an es.fuzzyType object based on a condition evaluated by a function.
//
// Example usage:
//
//	f := es.FuzzyFunc("user", "ki", func(key, value string) bool {
//	    return value != ""
//	})
//
// Parameters:
//   - key: A string representing the field name for the fuzzy query.
//   - value: The term to be searched for in the specified field.
//   - fn: A function that takes a key and value, and returns a boolean indicating
//     whether to create the es.fuzzyType object.
//
// Returns:
//
//	An es.fuzzyType object if the condition is true; otherwise, nil.
func FuzzyFunc(key string, value string, fn func(key string, value string) bool) fuzzyType {
	if !fn(key, value) {
		return nil
	}
	return Fuzzy(key, value)
}

// FuzzyIf creates an es.fuzzyType object based on a boolean condition.
//
// Example usage:
//
//	f := es.FuzzyIf("user", "ki", true)
//
// Parameters:
//   - key: A string representing the field name for the fuzzy query.
//   - value: The term to be searched for in the specified field.
//   - condition: A boolean value that determines whether to create the es.fuzzyType object.
//
// Returns:
//
//	An es.fuzzyType object if the condition is true; otherwise, nil.
func FuzzyIf(key string, value string, condition bool) fuzzyType {
	if !condition {
		return nil
	}
	return Fuzzy(key, value)
}

func (f fuzzyType) putInTheField(key string, value any) fuzzyType {
	return genericPutInTheFieldOfFirstChild(f, "fuzzy", key, value)
}
