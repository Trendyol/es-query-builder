package es

import (
	SuggestMode "github.com/Trendyol/es-query-builder/es/enums/suggest-mode"
)

type phraseSuggesterType Object

type directGeneratorType Object

// PhraseSuggester creates a phrase suggester for the given field.
//
// A phrase suggester builds corrected phrases based on individual term suggestions.
//
// Example usage:
//
//	suggester := es.PhraseSuggester("message").Text("noble prize").Size(1)
//
// Parameters:
//   - field: The field to use for phrase suggestions.
//
// Returns:
//
//	An es.phraseSuggesterType object representing the phrase suggester.
func PhraseSuggester(field string) phraseSuggesterType {
	return phraseSuggesterType{
		"phrase": Object{
			"field": field,
		},
	}
}

// Text sets the suggest text for this phrase suggester.
//
// Example usage:
//
//	suggester := es.PhraseSuggester("message").Text("noble prize")
//
// Parameters:
//   - text: The text to get suggestions for.
//
// Returns:
//
//	The updated es.phraseSuggesterType object with the "text" field set.
func (p phraseSuggesterType) Text(text string) phraseSuggesterType {
	p["text"] = text
	return p
}

// Size sets the maximum number of suggestions to return.
//
// Example usage:
//
//	suggester := es.PhraseSuggester("message").Size(1)
//
// Parameters:
//   - size: The maximum number of suggestions.
//
// Returns:
//
//	The updated es.phraseSuggesterType object with the "size" field set.
func (p phraseSuggesterType) Size(size int) phraseSuggesterType {
	return p.putInTheField("size", size)
}

// GramSize sets the max size of the n-grams for the field.
//
// Example usage:
//
//	suggester := es.PhraseSuggester("message").GramSize(3)
//
// Parameters:
//   - gramSize: The max size of the n-grams.
//
// Returns:
//
//	The updated es.phraseSuggesterType object with the "gram_size" field set.
func (p phraseSuggesterType) GramSize(gramSize int) phraseSuggesterType {
	return p.putInTheField("gram_size", gramSize)
}

// Confidence sets the confidence level for suggestions.
//
// Example usage:
//
//	suggester := es.PhraseSuggester("message").Confidence(1.0)
//
// Parameters:
//   - confidence: The confidence level threshold.
//
// Returns:
//
//	The updated es.phraseSuggesterType object with the "confidence" field set.
func (p phraseSuggesterType) Confidence(confidence float64) phraseSuggesterType {
	return p.putInTheField("confidence", confidence)
}

// MaxErrors sets the maximum percentage of terms that can be misspellings.
//
// Example usage:
//
//	suggester := es.PhraseSuggester("message").MaxErrors(2)
//
// Parameters:
//   - maxErrors: Absolute or relative maximum number of misspelled terms.
//
// Returns:
//
//	The updated es.phraseSuggesterType object with the "max_errors" field set.
func (p phraseSuggesterType) MaxErrors(maxErrors any) phraseSuggesterType {
	return p.putInTheField("max_errors", maxErrors)
}

// RealWordErrorLikelihood sets the likelihood of a term being a misspelled real word.
//
// Example usage:
//
//	suggester := es.PhraseSuggester("message").RealWordErrorLikelihood(0.95)
//
// Parameters:
//   - likelihood: The likelihood value between 0 and 1.
//
// Returns:
//
//	The updated es.phraseSuggesterType object with the "real_word_error_likelihood" field set.
func (p phraseSuggesterType) RealWordErrorLikelihood(likelihood float64) phraseSuggesterType {
	return p.putInTheField("real_word_error_likelihood", likelihood)
}

// Separator sets the separator used between terms in the bigram field.
//
// Example usage:
//
//	suggester := es.PhraseSuggester("message").Separator(" ")
//
// Parameters:
//   - separator: The separator string.
//
// Returns:
//
//	The updated es.phraseSuggesterType object with the "separator" field set.
func (p phraseSuggesterType) Separator(separator string) phraseSuggesterType {
	return p.putInTheField("separator", separator)
}

// DirectGenerator sets one or more direct generators for the phrase suggester.
//
// Example usage:
//
//	suggester := es.PhraseSuggester("message").
//		DirectGenerator(es.DirectGenerator("message.trigram").SuggestMode(SuggestMode.Always))
//
// Parameters:
//   - generators: A variadic list of direct generator configurations.
//
// Returns:
//
//	The updated es.phraseSuggesterType object with the "direct_generator" field set.
func (p phraseSuggesterType) DirectGenerator(generators ...directGeneratorType) phraseSuggesterType {
	return p.putInTheField("direct_generator", generators)
}

func (p phraseSuggesterType) putInTheField(key string, value any) phraseSuggesterType {
	return genericPutInTheField(p, "phrase", key, value)
}

// DirectGenerator creates a direct generator for a phrase suggester.
//
// Example usage:
//
//	generator := es.DirectGenerator("message.trigram").MinWordLength(3)
//
// Parameters:
//   - field: The field to use for candidate generation.
//
// Returns:
//
//	An es.directGeneratorType object representing the direct generator.
func DirectGenerator(field string) directGeneratorType {
	return directGeneratorType{
		"field": field,
	}
}

// SuggestMode sets the suggest mode for the direct generator.
//
// Example usage:
//
//	generator := es.DirectGenerator("message").SuggestMode(SuggestMode.Always)
//
// Parameters:
//   - suggestMode: A SuggestMode value (missing, popular, or always).
//
// Returns:
//
//	The updated es.directGeneratorType object with the "suggest_mode" field set.
func (d directGeneratorType) SuggestMode(suggestMode SuggestMode.SuggestMode) directGeneratorType {
	d["suggest_mode"] = suggestMode
	return d
}

// MinWordLength sets the minimum word length for the direct generator.
//
// Example usage:
//
//	generator := es.DirectGenerator("message").MinWordLength(3)
//
// Parameters:
//   - minWordLength: The minimum word length.
//
// Returns:
//
//	The updated es.directGeneratorType object with the "min_word_length" field set.
func (d directGeneratorType) MinWordLength(minWordLength int) directGeneratorType {
	d["min_word_length"] = minWordLength
	return d
}

// PrefixLength sets the prefix length for the direct generator.
//
// Example usage:
//
//	generator := es.DirectGenerator("message").PrefixLength(1)
//
// Parameters:
//   - prefixLength: The number of initial characters that must match.
//
// Returns:
//
//	The updated es.directGeneratorType object with the "prefix_length" field set.
func (d directGeneratorType) PrefixLength(prefixLength int) directGeneratorType {
	d["prefix_length"] = prefixLength
	return d
}

// MaxEdits sets the maximum edit distance for the direct generator.
//
// Example usage:
//
//	generator := es.DirectGenerator("message").MaxEdits(2)
//
// Parameters:
//   - maxEdits: The maximum number of edits.
//
// Returns:
//
//	The updated es.directGeneratorType object with the "max_edits" field set.
func (d directGeneratorType) MaxEdits(maxEdits int) directGeneratorType {
	d["max_edits"] = maxEdits
	return d
}
