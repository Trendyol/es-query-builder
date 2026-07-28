package es

type completionSuggesterType Object

type completionFuzzyType Object

// CompletionSuggester creates a completion suggester for the given field.
//
// A completion suggester provides auto-complete / search-as-you-type functionality.
//
// Example usage:
//
//	suggester := es.CompletionSuggester("name.suggest").
//		Prefix("ni").Size(5).SkipDuplicates(true)
//
// Parameters:
//   - field: The completion field to use for suggestions.
//
// Returns:
//
//	An es.completionSuggesterType object representing the completion suggester.
func CompletionSuggester(field string) completionSuggesterType {
	return completionSuggesterType{
		"completion": Object{
			"field": field,
		},
	}
}

// Text sets the suggest text for this completion suggester.
//
// Example usage:
//
//	suggester := es.CompletionSuggester("name.suggest").Text("nike")
//
// Parameters:
//   - text: The text to get suggestions for.
//
// Returns:
//
//	The updated es.completionSuggesterType object with the "text" field set.
func (c completionSuggesterType) Text(text string) completionSuggesterType {
	c["text"] = text
	return c
}

// Prefix sets the prefix used for completion suggestions.
//
// Example usage:
//
//	suggester := es.CompletionSuggester("name.suggest").Prefix("ni")
//
// Parameters:
//   - prefix: The prefix string to complete.
//
// Returns:
//
//	The updated es.completionSuggesterType object with the "prefix" field set.
func (c completionSuggesterType) Prefix(prefix string) completionSuggesterType {
	c["prefix"] = prefix
	return c
}

// Regex sets a regular expression used for completion suggestions.
//
// Example usage:
//
//	suggester := es.CompletionSuggester("name.suggest").Regex("n[iI]ke.*")
//
// Parameters:
//   - regex: The regular expression pattern.
//
// Returns:
//
//	The updated es.completionSuggesterType object with the "regex" field set.
func (c completionSuggesterType) Regex(regex string) completionSuggesterType {
	c["regex"] = regex
	return c
}

// Size sets the maximum number of suggestions to return.
//
// Example usage:
//
//	suggester := es.CompletionSuggester("name.suggest").Size(5)
//
// Parameters:
//   - size: The maximum number of suggestions.
//
// Returns:
//
//	The updated es.completionSuggesterType object with the "size" field set.
func (c completionSuggesterType) Size(size int) completionSuggesterType {
	return c.putInTheField("size", size)
}

// SkipDuplicates skips duplicate suggestions.
//
// Example usage:
//
//	suggester := es.CompletionSuggester("name.suggest").SkipDuplicates(true)
//
// Parameters:
//   - skipDuplicates: Whether to skip duplicate suggestions.
//
// Returns:
//
//	The updated es.completionSuggesterType object with the "skip_duplicates" field set.
func (c completionSuggesterType) SkipDuplicates(skipDuplicates bool) completionSuggesterType {
	return c.putInTheField("skip_duplicates", skipDuplicates)
}

// Fuzzy sets fuzzy options for the completion suggester.
//
// Example usage:
//
//	suggester := es.CompletionSuggester("name.suggest").
//		Fuzzy(es.CompletionFuzzy().Fuzziness(1).Transpositions(true))
//
// Parameters:
//   - fuzzy: An es.completionFuzzyType object with fuzzy options.
//
// Returns:
//
//	The updated es.completionSuggesterType object with the "fuzzy" field set.
func (c completionSuggesterType) Fuzzy(fuzzy completionFuzzyType) completionSuggesterType {
	return c.putInTheField("fuzzy", fuzzy)
}

func (c completionSuggesterType) putInTheField(key string, value any) completionSuggesterType {
	return genericPutInTheField(c, "completion", key, value)
}

// CompletionFuzzy creates fuzzy options for a completion suggester.
//
// Example usage:
//
//	fuzzy := es.CompletionFuzzy().Fuzziness("AUTO").MinLength(3)
//
// Returns:
//
//	An es.completionFuzzyType object ready for further configuration.
func CompletionFuzzy() completionFuzzyType {
	return completionFuzzyType{}
}

// Fuzziness sets the fuzziness value for completion suggestions.
//
// Accepts either a string (e.g. "AUTO") or an integer (0, 1, 2).
//
// Example usage:
//
//	fuzzy := es.CompletionFuzzy().Fuzziness("AUTO")
//
// Parameters:
//   - fuzziness: The fuzziness value (string or int).
//
// Returns:
//
//	The updated es.completionFuzzyType object with the "fuzziness" field set.
func (f completionFuzzyType) Fuzziness(fuzziness any) completionFuzzyType {
	f["fuzziness"] = fuzziness
	return f
}

// Transpositions sets whether transpositions count as a single edit.
//
// Example usage:
//
//	fuzzy := es.CompletionFuzzy().Transpositions(true)
//
// Parameters:
//   - transpositions: Whether to include transpositions.
//
// Returns:
//
//	The updated es.completionFuzzyType object with the "transpositions" field set.
func (f completionFuzzyType) Transpositions(transpositions bool) completionFuzzyType {
	f["transpositions"] = transpositions
	return f
}

// MinLength sets the minimum length of the input before fuzzy suggestions are produced.
//
// Example usage:
//
//	fuzzy := es.CompletionFuzzy().MinLength(3)
//
// Parameters:
//   - minLength: The minimum input length.
//
// Returns:
//
//	The updated es.completionFuzzyType object with the "min_length" field set.
func (f completionFuzzyType) MinLength(minLength int) completionFuzzyType {
	f["min_length"] = minLength
	return f
}

// PrefixLength sets the minimum length of the input that is not checked for fuzzy options.
//
// Example usage:
//
//	fuzzy := es.CompletionFuzzy().PrefixLength(1)
//
// Parameters:
//   - prefixLength: The prefix length that must match exactly.
//
// Returns:
//
//	The updated es.completionFuzzyType object with the "prefix_length" field set.
func (f completionFuzzyType) PrefixLength(prefixLength int) completionFuzzyType {
	f["prefix_length"] = prefixLength
	return f
}

// UnicodeAware sets whether all measurements are done in unicode code points.
//
// Example usage:
//
//	fuzzy := es.CompletionFuzzy().UnicodeAware(true)
//
// Parameters:
//   - unicodeAware: Whether to use unicode-aware measurements.
//
// Returns:
//
//	The updated es.completionFuzzyType object with the "unicode_aware" field set.
func (f completionFuzzyType) UnicodeAware(unicodeAware bool) completionFuzzyType {
	f["unicode_aware"] = unicodeAware
	return f
}
