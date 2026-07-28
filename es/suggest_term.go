package es

import (
	StringDistance "github.com/Trendyol/es-query-builder/es/enums/string-distance"
	SuggestMode "github.com/Trendyol/es-query-builder/es/enums/suggest-mode"
)

type termSuggesterType Object

// TermSuggester creates a term suggester for the given field.
//
// A term suggester suggests alternative terms based on edit distance.
//
// Example usage:
//
//	suggester := es.TermSuggester("message").Text("triyng out Elasticsearch")
//
// Parameters:
//   - field: The field to use for term suggestions.
//
// Returns:
//
//	An es.termSuggesterType object representing the term suggester.
func TermSuggester(field string) termSuggesterType {
	return termSuggesterType{
		"term": Object{
			"field": field,
		},
	}
}

// Text sets the suggest text for this term suggester.
//
// Example usage:
//
//	suggester := es.TermSuggester("message").Text("triyng")
//
// Parameters:
//   - text: The text to get suggestions for.
//
// Returns:
//
//	The updated es.termSuggesterType object with the "text" field set.
func (t termSuggesterType) Text(text string) termSuggesterType {
	t["text"] = text
	return t
}

// Size sets the maximum number of suggestions to return.
//
// Example usage:
//
//	suggester := es.TermSuggester("message").Size(5)
//
// Parameters:
//   - size: The maximum number of suggestions.
//
// Returns:
//
//	The updated es.termSuggesterType object with the "size" field set.
func (t termSuggesterType) Size(size int) termSuggesterType {
	return t.putInTheField("size", size)
}

// SuggestMode sets which suggestions are included.
//
// Example usage:
//
//	suggester := es.TermSuggester("message").SuggestMode(SuggestMode.Popular)
//
// Parameters:
//   - suggestMode: A SuggestMode value (missing, popular, or always).
//
// Returns:
//
//	The updated es.termSuggesterType object with the "suggest_mode" field set.
func (t termSuggesterType) SuggestMode(suggestMode SuggestMode.SuggestMode) termSuggesterType {
	return t.putInTheField("suggest_mode", suggestMode)
}

// MinWordLength sets the minimum length a term must have to be included.
//
// Example usage:
//
//	suggester := es.TermSuggester("message").MinWordLength(4)
//
// Parameters:
//   - minWordLength: The minimum word length.
//
// Returns:
//
//	The updated es.termSuggesterType object with the "min_word_length" field set.
func (t termSuggesterType) MinWordLength(minWordLength int) termSuggesterType {
	return t.putInTheField("min_word_length", minWordLength)
}

// MaxEdits sets the maximum edit distance for suggestions.
//
// Example usage:
//
//	suggester := es.TermSuggester("message").MaxEdits(2)
//
// Parameters:
//   - maxEdits: The maximum number of edits (1 or 2).
//
// Returns:
//
//	The updated es.termSuggesterType object with the "max_edits" field set.
func (t termSuggesterType) MaxEdits(maxEdits int) termSuggesterType {
	return t.putInTheField("max_edits", maxEdits)
}

// PrefixLength sets the number of initial characters that must match.
//
// Example usage:
//
//	suggester := es.TermSuggester("message").PrefixLength(1)
//
// Parameters:
//   - prefixLength: The number of initial characters that must match exactly.
//
// Returns:
//
//	The updated es.termSuggesterType object with the "prefix_length" field set.
func (t termSuggesterType) PrefixLength(prefixLength int) termSuggesterType {
	return t.putInTheField("prefix_length", prefixLength)
}

// MinDocFreq sets the minimum document frequency for a suggestion to be included.
//
// Example usage:
//
//	suggester := es.TermSuggester("message").MinDocFreq(0.01)
//
// Parameters:
//   - minDocFreq: Absolute or relative minimum document frequency.
//
// Returns:
//
//	The updated es.termSuggesterType object with the "min_doc_freq" field set.
func (t termSuggesterType) MinDocFreq(minDocFreq float64) termSuggesterType {
	return t.putInTheField("min_doc_freq", minDocFreq)
}

// MaxTermFreq sets the maximum document frequency of the original term.
//
// Example usage:
//
//	suggester := es.TermSuggester("message").MaxTermFreq(0.01)
//
// Parameters:
//   - maxTermFreq: Absolute or relative maximum term frequency.
//
// Returns:
//
//	The updated es.termSuggesterType object with the "max_term_freq" field set.
func (t termSuggesterType) MaxTermFreq(maxTermFreq float64) termSuggesterType {
	return t.putInTheField("max_term_freq", maxTermFreq)
}

// StringDistance sets the string distance algorithm.
//
// Example usage:
//
//	suggester := es.TermSuggester("message").StringDistance(StringDistance.Levenshtein)
//
// Parameters:
//   - stringDistance: A StringDistance algorithm value.
//
// Returns:
//
//	The updated es.termSuggesterType object with the "string_distance" field set.
func (t termSuggesterType) StringDistance(stringDistance StringDistance.StringDistance) termSuggesterType {
	return t.putInTheField("string_distance", stringDistance)
}

// ShardSize sets the number of suggestions considered on each shard.
//
// Example usage:
//
//	suggester := es.TermSuggester("message").ShardSize(5)
//
// Parameters:
//   - shardSize: The number of suggestions to consider per shard.
//
// Returns:
//
//	The updated es.termSuggesterType object with the "shard_size" field set.
func (t termSuggesterType) ShardSize(shardSize int) termSuggesterType {
	return t.putInTheField("shard_size", shardSize)
}

func (t termSuggesterType) putInTheField(key string, value any) termSuggesterType {
	return genericPutInTheField(t, "term", key, value)
}
