package es

type regexpType Object

// Regexp creates a new regexpType object with the specified key-value pair.
//
// This function initializes a regexpType object with a single term query, where the
// key is the field name and the value is the term to search for. This is typically
// used to construct a regexp query in search queries.
//
// Example usage:
//
//	t := Regexp("endpoint", "/books/.*")
//	// t now contains a regexpType object with a regexp query for the "endpoint" field.
//
// Parameters:
//   - key: A string representing the field name for the regexp query.
//   - value: The value to be searched for in the specified field. The type is regexp.
//
// Returns:
//
//	A termType object containing the specified term query.
func Regexp(key string, value string) regexpType {
	regexp := regexpType{
		key: Object{},
	}
	regexp.value(value)
	return regexp
}

func (r regexpType) value(value string) regexpType {
	for key := range r {
		if rangeObject, ok := r[key].(Object); ok {
			rangeObject["value"] = value
		}
	}
	return r
}

// Flags Enables optional operators for the regular expression.
// Example usage:
//
//	regexp := Regexp("endpoint", "/books/.*").Flags("ALL")
//	// regexp now a "flags" field set "ALL" in the regexp query object.
//
// Parameters:
//   - flags: A string value representing flags value to be applied to the regexp query.
//
// Returns:
//
//	The updated regexp object with the "flags" field set to the specified value.
func (r regexpType) Flags(flags string) regexpType {
	for key := range r {
		if rangeObject, ok := r[key].(Object); ok {
			rangeObject["flags"] = flags
		}
	}
	return r
}

// CaseInsensitive Allows case insensitive matching of the regular expression
// value with the indexed field values when set to true.
// Example usage:
//
//	regexp := Regexp("endpoint", "/books/.*").CaseInsensitive(true)
//	// regexp now a "case_insensitive" field set true in the regexp query object.
//
// Parameters:
//   - caseInsensitive: A bool value representing case insensitive value to be applied to the regexp query.
//
// Returns:
//
//	The updated regexp object with the "case_insensitive" field set to the specified value.
func (r regexpType) CaseInsensitive(caseInsensitive bool) regexpType {
	for key := range r {
		if rangeObject, ok := r[key].(Object); ok {
			rangeObject["case_insensitive"] = caseInsensitive
		}
	}
	return r
}

// MaxDeterminizedStates Maximum number of automaton states required for the query.
// Example usage:
//
//	regexp := Regexp("endpoint", "/books/.*").MaxDeterminizedStates(10000)
//	// regexp now a "max_determinized_states" field set 10000 in the regexp query object.
//
// Parameters:
//   - maxDeterminizedStates: A bool value representing max_determinized_states value to be applied to the regexp query.
//
// Returns:
//
//	The updated regexp object with the "max_determinized_states" field set to the specified value.
func (r regexpType) MaxDeterminizedStates(maxDeterminizedStates int) regexpType {
	for key := range r {
		if rangeObject, ok := r[key].(Object); ok {
			rangeObject["max_determinized_states"] = maxDeterminizedStates
		}
	}
	return r
}

// ReWrite Method used to rewrite the query.
// Example usage:
//
//	regexp := Regexp("endpoint", "/books/.*").ReWrite("a")
//	// regexp now a "rewrite" field set "a" in the regexp query object.
//
// Parameters:
//   - rewrite: A string value representing rewrite value to be applied to the regexp query.
//
// Returns:
//
//	The updated regexp object with the "max_determinized_states" field set to the specified value.
func (r regexpType) ReWrite(rewrite string) regexpType {
	for key := range r {
		if rangeObject, ok := r[key].(Object); ok {
			rangeObject["rewrite"] = rewrite
		}
	}
	return r
}
