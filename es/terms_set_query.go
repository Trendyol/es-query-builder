package es

type termsSetType Object

// TermsSet creates a new es.termsSetType object with the specified field and terms.
//
// This function initializes an es.termsSetType object for a terms_set query, where the key
// represents the field name, and values are the terms that must be matched in the field.
//
// Example usage:
//
//	termsSet : es.TermsSet("tags", "go", "elasticsearch")
//	// termsSet now contains an es.termsSetType object that matches documents where the "tags" field contains "go" or "elasticsearch".
//
// Parameters:
//   - key: A string representing the field name for the terms_set query.
//   - values: A variadic list of terms to be matched in the specified field.
//
// Returns:
//
//	An es.termsSetType object containing the specified terms_set query.
func TermsSet(key string, values ...any) termsSetType {
	return termsSetType{
		"terms_set": Object{
			key: Object{
				"terms": values,
			},
		},
	}
}

// Boost sets the "boost" field in the terms_set query.
//
// This method configures the terms_set query to use a specified boost factor, which influences
// the relevance scoring of the matched documents.
//
// Example usage:
//
//	termsSet : es.TermsSet("tags", "go", "elasticsearch").Boost(1.2)
//	// termsSet now has a "boost" field set to 1.2 in the terms_set query object.
//
// Parameters:
//   - boost: A float64 value representing the boost factor to be applied to the terms_set query.
//
// Returns:
//
//	The updated es.termsSetType object with the "boost" field set.
func (t termsSetType) Boost(boost float64) termsSetType {
	return t.putInTheField("boost", boost)
}

// MinimumShouldMatchField sets the "minimum_should_match_field" in the terms_set query.
//
// This method specifies a field that determines the minimum number of terms that must be matched.
// The field's value should be an integer representing the required number of matches.
//
// Example usage:
//
//	termsSet : es.TermsSet("tags", "go", "elasticsearch").MinimumShouldMatchField("match_count")
//	// termsSet now has a "minimum_should_match_field" set to "match_count".
//
// Parameters:
//   - minimumShouldMatchField: A string representing the field name that specifies the required number of matches.
//
// Returns:
//
//	The updated es.termsSetType object with the "minimum_should_match_field" set.
func (t termsSetType) MinimumShouldMatchField(minimumShouldMatchField string) termsSetType {
	return t.putInTheField("minimum_should_match_field", minimumShouldMatchField)
}

// MinimumShouldMatchScript sets the "minimum_should_match_script" in the terms_set query.
//
// This method specifies a script that determines the minimum number of terms that must be matched,
// allowing for more dynamic query logic.
//
// Example usage:
//
//	script := es.ScriptSource("return doc['tag_count'].value;", es.ScriptLanguage.Painless)
//	termsSet : es.TermsSet("tags", "go", "elasticsearch").MinimumShouldMatchScript(script)
//	// termsSet now has a "minimum_should_match_script" field set with the provided script.
//
// Parameters:
//   - minimumShouldMatchScript: A scriptType object defining the script to determine the required number of matches.
//
// Returns:
//
//	The updated es.termsSetType object with the "minimum_should_match_script" field set.
func (t termsSetType) MinimumShouldMatchScript(minimumShouldMatchScript scriptType) termsSetType {
	return t.putInTheField("minimum_should_match_script", minimumShouldMatchScript)
}

func (t termsSetType) putInTheField(key string, value any) termsSetType {
	return genericPutInTheFieldOfFirstChild(t, "terms_set", key, value)
}
