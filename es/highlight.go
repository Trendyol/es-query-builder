package es

import (
	BoundaryScanner "github.com/Trendyol/es-query-builder/es/enums/boundary-scanner"
	Fragmenter "github.com/Trendyol/es-query-builder/es/enums/fragmenter"
	HighlighterType "github.com/Trendyol/es-query-builder/es/enums/highlighter-type"
)

type highlightType Object

type highlightFieldType Object

// Highlight creates a new es.highlightType object.
//
// This function initializes a highlight configuration object that can be used
// to configure how search results are highlighted. The highlight object is
// typically added to a query using the Highlight method on es.Object.
//
// Example usage:
//
//	h := es.Highlight().
//		PreTags("<em>").
//		PostTags("</em>").
//		Field(es.HighlightField("title"))
//	query := es.NewQuery(es.MatchAll()).Highlight(h)
//
// Returns:
//
//	An es.highlightType object ready for further configuration.
func Highlight() highlightType {
	return highlightType{}
}

// PreTags sets the "pre_tags" parameter in an es.highlightType.
//
// This method specifies the HTML tags or strings to insert before the highlighted text.
// Multiple tags can be provided for use with different highlight levels.
//
// Example usage:
//
//	h := es.Highlight().PreTags("<em>", "<strong>")
//	// h now includes a "pre_tags" parameter with the specified tags.
//
// Parameters:
//   - tags: A variadic list of strings representing the pre-tags.
//
// Returns:
//
//	The updated es.highlightType object with the "pre_tags" parameter set.
func (h highlightType) PreTags(tags ...string) highlightType {
	h["pre_tags"] = tags
	return h
}

// PostTags sets the "post_tags" parameter in an es.highlightType.
//
// This method specifies the HTML tags or strings to insert after the highlighted text.
// Multiple tags can be provided for use with different highlight levels.
//
// Example usage:
//
//	h := es.Highlight().PostTags("</em>", "</strong>")
//	// h now includes a "post_tags" parameter with the specified tags.
//
// Parameters:
//   - tags: A variadic list of strings representing the post-tags.
//
// Returns:
//
//	The updated es.highlightType object with the "post_tags" parameter set.
func (h highlightType) PostTags(tags ...string) highlightType {
	h["post_tags"] = tags
	return h
}

// Field adds a highlight field configuration to the es.highlightType.
//
// This method adds one or more field configurations to the "fields" object in the
// highlight configuration. Each field can have its own highlight settings that
// override the global highlight settings.
//
// Example usage:
//
//	h := es.Highlight().
//		Field(es.HighlightField("title").NumberOfFragments(0)).
//		Field(es.HighlightField("content").FragmentSize(150).NumberOfFragments(3))
//	// h now includes a "fields" object with configurations for "title" and "content".
//
// Parameters:
//   - fields: A variadic list of es.highlightFieldType objects.
//
// Returns:
//
//	The updated es.highlightType object with the field configurations added.
func (h highlightType) Field(fields ...highlightFieldType) highlightType {
	highlightFields, ok := h["fields"].(Object)
	if !ok {
		highlightFields = Object{}
	}
	for _, field := range fields {
		for key, value := range field {
			highlightFields[key] = value
		}
	}
	h["fields"] = highlightFields
	return h
}

// Type sets the "type" parameter in an es.highlightType.
//
// This method specifies which highlighter implementation to use. Elasticsearch
// provides three highlighter implementations: unified, plain, and fvh.
//
// Example usage:
//
//	h := es.Highlight().Type(HighlighterType.Unified)
//	// h now includes a "type" parameter set to "unified".
//
// Parameters:
//   - highlighterType: A HighlighterType.HighlighterType value representing the highlighter type.
//
// Returns:
//
//	The updated es.highlightType object with the "type" parameter set.
func (h highlightType) Type(highlighterType HighlighterType.HighlighterType) highlightType {
	h["type"] = highlighterType
	return h
}

// Order sets the "order" parameter in an es.highlightType.
//
// This method specifies the order in which highlighted fragments are sorted.
// Setting it to "score" sorts highlighted fragments by relevance score.
//
// Example usage:
//
//	h := es.Highlight().Order("score")
//	// h now includes an "order" parameter set to "score".
//
// Parameters:
//   - order: A string representing the order for highlighted fragments.
//
// Returns:
//
//	The updated es.highlightType object with the "order" parameter set.
func (h highlightType) Order(order string) highlightType {
	h["order"] = order
	return h
}

// Encoder sets the "encoder" parameter in an es.highlightType.
//
// This method specifies how the highlighted text should be encoded. The "default"
// encoder provides no encoding, while "html" encodes the highlighted text as HTML.
//
// Example usage:
//
//	h := es.Highlight().Encoder("html")
//	// h now includes an "encoder" parameter set to "html".
//
// Parameters:
//   - encoder: A string representing the encoder type ("default" or "html").
//
// Returns:
//
//	The updated es.highlightType object with the "encoder" parameter set.
func (h highlightType) Encoder(encoder string) highlightType {
	h["encoder"] = encoder
	return h
}

// RequireFieldMatch sets the "require_field_match" parameter in an es.highlightType.
//
// This method specifies whether only fields that match the query should be highlighted.
// When set to false, all fields will be highlighted regardless of whether they contributed
// to the query match.
//
// Example usage:
//
//	h := es.Highlight().RequireFieldMatch(false)
//	// h now includes a "require_field_match" parameter set to false.
//
// Parameters:
//   - requireFieldMatch: A boolean indicating whether to require field match for highlighting.
//
// Returns:
//
//	The updated es.highlightType object with the "require_field_match" parameter set.
func (h highlightType) RequireFieldMatch(requireFieldMatch bool) highlightType {
	h["require_field_match"] = requireFieldMatch
	return h
}

// FragmentSize sets the "fragment_size" parameter in an es.highlightType.
//
// This method specifies the size of the highlighted fragment in characters.
// The default value is 100.
//
// Example usage:
//
//	h := es.Highlight().FragmentSize(150)
//	// h now includes a "fragment_size" parameter set to 150.
//
// Parameters:
//   - fragmentSize: An integer representing the fragment size in characters.
//
// Returns:
//
//	The updated es.highlightType object with the "fragment_size" parameter set.
func (h highlightType) FragmentSize(fragmentSize int) highlightType {
	h["fragment_size"] = fragmentSize
	return h
}

// NumberOfFragments sets the "number_of_fragments" parameter in an es.highlightType.
//
// This method specifies the maximum number of fragments to return. If set to 0,
// the entire field content is returned as a single fragment.
//
// Example usage:
//
//	h := es.Highlight().NumberOfFragments(5)
//	// h now includes a "number_of_fragments" parameter set to 5.
//
// Parameters:
//   - numberOfFragments: An integer representing the maximum number of fragments.
//
// Returns:
//
//	The updated es.highlightType object with the "number_of_fragments" parameter set.
func (h highlightType) NumberOfFragments(numberOfFragments int) highlightType {
	h["number_of_fragments"] = numberOfFragments
	return h
}

// NoMatchSize sets the "no_match_size" parameter in an es.highlightType.
//
// This method specifies the amount of text to return from the beginning of the field
// if there are no matching fragments to highlight. The default value is 0, which means
// nothing is returned.
//
// Example usage:
//
//	h := es.Highlight().NoMatchSize(150)
//	// h now includes a "no_match_size" parameter set to 150.
//
// Parameters:
//   - noMatchSize: An integer representing the number of characters to return when no match is found.
//
// Returns:
//
//	The updated es.highlightType object with the "no_match_size" parameter set.
func (h highlightType) NoMatchSize(noMatchSize int) highlightType {
	h["no_match_size"] = noMatchSize
	return h
}

// BoundaryScanner sets the "boundary_scanner" parameter in an es.highlightType.
//
// This method specifies how highlighted fragments are bounded. The boundary scanner
// determines the strategy for finding the boundaries of highlighted snippets.
//
// Example usage:
//
//	h := es.Highlight().BoundaryScanner(BoundaryScanner.Sentence)
//	// h now includes a "boundary_scanner" parameter set to "sentence".
//
// Parameters:
//   - boundaryScanner: A BoundaryScanner.BoundaryScanner value representing the boundary scanner type.
//
// Returns:
//
//	The updated es.highlightType object with the "boundary_scanner" parameter set.
func (h highlightType) BoundaryScanner(boundaryScanner BoundaryScanner.BoundaryScanner) highlightType {
	h["boundary_scanner"] = boundaryScanner
	return h
}

// BoundaryChars sets the "boundary_chars" parameter in an es.highlightType.
//
// This method specifies a string containing each boundary character for the chars
// boundary scanner. The default value is ".,!? \t\n".
//
// Example usage:
//
//	h := es.Highlight().BoundaryChars(".,!? \t\n")
//	// h now includes a "boundary_chars" parameter with the specified characters.
//
// Parameters:
//   - boundaryChars: A string containing the boundary characters.
//
// Returns:
//
//	The updated es.highlightType object with the "boundary_chars" parameter set.
func (h highlightType) BoundaryChars(boundaryChars string) highlightType {
	h["boundary_chars"] = boundaryChars
	return h
}

// BoundaryMaxScan sets the "boundary_max_scan" parameter in an es.highlightType.
//
// This method specifies how far to scan for boundary characters when using the chars
// boundary scanner. The default value is 20.
//
// Example usage:
//
//	h := es.Highlight().BoundaryMaxScan(20)
//	// h now includes a "boundary_max_scan" parameter set to 20.
//
// Parameters:
//   - boundaryMaxScan: An integer representing the maximum number of characters to scan.
//
// Returns:
//
//	The updated es.highlightType object with the "boundary_max_scan" parameter set.
func (h highlightType) BoundaryMaxScan(boundaryMaxScan int) highlightType {
	h["boundary_max_scan"] = boundaryMaxScan
	return h
}

// BoundaryScannerLocale sets the "boundary_scanner_locale" parameter in an es.highlightType.
//
// This method specifies the locale to use for the sentence and word boundary scanners.
// This allows proper handling of word and sentence boundaries for different languages.
//
// Example usage:
//
//	h := es.Highlight().BoundaryScannerLocale("en-US")
//	// h now includes a "boundary_scanner_locale" parameter set to "en-US".
//
// Parameters:
//   - locale: A string representing the locale (e.g., "en-US").
//
// Returns:
//
//	The updated es.highlightType object with the "boundary_scanner_locale" parameter set.
func (h highlightType) BoundaryScannerLocale(locale string) highlightType {
	h["boundary_scanner_locale"] = locale
	return h
}

// Fragmenter sets the "fragmenter" parameter in an es.highlightType.
//
// This method specifies how text should be broken up into highlight fragments.
// The simple fragmenter breaks text into same-sized fragments, while the span
// fragmenter tries to avoid breaking up highlighted terms.
//
// Example usage:
//
//	h := es.Highlight().Fragmenter(Fragmenter.Span)
//	// h now includes a "fragmenter" parameter set to "span".
//
// Parameters:
//   - fragmenter: A Fragmenter.Fragmenter value representing the fragmenter type.
//
// Returns:
//
//	The updated es.highlightType object with the "fragmenter" parameter set.
func (h highlightType) Fragmenter(fragmenter Fragmenter.Fragmenter) highlightType {
	h["fragmenter"] = fragmenter
	return h
}

// FragmentOffset sets the "fragment_offset" parameter in an es.highlightType.
//
// This method specifies the margin from which to start highlighting. This is only
// valid when using the fvh highlighter.
//
// Example usage:
//
//	h := es.Highlight().FragmentOffset(10)
//	// h now includes a "fragment_offset" parameter set to 10.
//
// Parameters:
//   - fragmentOffset: An integer representing the fragment offset.
//
// Returns:
//
//	The updated es.highlightType object with the "fragment_offset" parameter set.
func (h highlightType) FragmentOffset(fragmentOffset int) highlightType {
	h["fragment_offset"] = fragmentOffset
	return h
}

// MaxFragmentLength sets the "max_fragment_length" parameter in an es.highlightType.
//
// This method specifies the maximum length of a fragment in characters. This is only
// valid when using the unified highlighter.
//
// Example usage:
//
//	h := es.Highlight().MaxFragmentLength(200)
//	// h now includes a "max_fragment_length" parameter set to 200.
//
// Parameters:
//   - maxFragmentLength: An integer representing the maximum fragment length.
//
// Returns:
//
//	The updated es.highlightType object with the "max_fragment_length" parameter set.
func (h highlightType) MaxFragmentLength(maxFragmentLength int) highlightType {
	h["max_fragment_length"] = maxFragmentLength
	return h
}

// MaxAnalyzedOffset sets the "max_analyzed_offset" parameter in an es.highlightType.
//
// This method specifies the maximum number of characters that will be analyzed for
// highlighting. The remaining text will not be processed. This setting is particularly
// useful for large text fields to avoid excessive memory usage.
//
// Example usage:
//
//	h := es.Highlight().MaxAnalyzedOffset(1000000)
//	// h now includes a "max_analyzed_offset" parameter set to 1000000.
//
// Parameters:
//   - maxAnalyzedOffset: An integer representing the maximum analyzed offset.
//
// Returns:
//
//	The updated es.highlightType object with the "max_analyzed_offset" parameter set.
func (h highlightType) MaxAnalyzedOffset(maxAnalyzedOffset int) highlightType {
	h["max_analyzed_offset"] = maxAnalyzedOffset
	return h
}

// HighlightQuery sets the "highlight_query" parameter in an es.highlightType.
//
// This method specifies an additional query to use for highlighting. This allows
// highlighting based on a different query than the main search query.
//
// Example usage:
//
//	h := es.Highlight().HighlightQuery(es.Match("content", "elasticsearch"))
//	// h now includes a "highlight_query" parameter with the specified query.
//
// Parameters:
//   - query: An object representing the highlight query. It can be of any type.
//
// Returns:
//
//	The updated es.highlightType object with the "highlight_query" parameter set.
func (h highlightType) HighlightQuery(query any) highlightType {
	if field, ok := correctType(query); ok {
		h["highlight_query"] = field
	}
	return h
}

// TagsSchema sets the "tags_schema" parameter in an es.highlightType.
//
// This method specifies the tags schema to use for highlighting. Setting it to "styled"
// uses a built-in set of pre/post tags.
//
// Example usage:
//
//	h := es.Highlight().TagsSchema("styled")
//	// h now includes a "tags_schema" parameter set to "styled".
//
// Parameters:
//   - tagsSchema: A string representing the tags schema.
//
// Returns:
//
//	The updated es.highlightType object with the "tags_schema" parameter set.
func (h highlightType) TagsSchema(tagsSchema string) highlightType {
	h["tags_schema"] = tagsSchema
	return h
}

// HighlightField creates a new es.highlightFieldType object with the specified field name.
//
// This function initializes a highlight field configuration for use in the "fields" object
// of a highlight configuration. Each field can have its own highlight settings that
// override the global highlight settings.
//
// Example usage:
//
//	hf := es.HighlightField("title")
//	// hf now contains a highlight field configuration for the "title" field.
//
// Parameters:
//   - field: A string representing the field name to highlight.
//
// Returns:
//
//	An es.highlightFieldType object with the specified field.
func HighlightField(field string) highlightFieldType {
	return highlightFieldType{
		field: Object{},
	}
}

// PreTags sets the "pre_tags" parameter in an es.highlightFieldType.
//
// This method specifies the HTML tags or strings to insert before the highlighted text
// for this specific field, overriding the global pre_tags setting.
//
// Example usage:
//
//	hf := es.HighlightField("title").PreTags("<b>")
//	// hf now includes a "pre_tags" parameter for the "title" field.
//
// Parameters:
//   - tags: A variadic list of strings representing the pre-tags.
//
// Returns:
//
//	The updated es.highlightFieldType object with the "pre_tags" parameter set.
func (hf highlightFieldType) PreTags(tags ...string) highlightFieldType {
	return hf.putInTheField("pre_tags", tags)
}

// PostTags sets the "post_tags" parameter in an es.highlightFieldType.
//
// This method specifies the HTML tags or strings to insert after the highlighted text
// for this specific field, overriding the global post_tags setting.
//
// Example usage:
//
//	hf := es.HighlightField("title").PostTags("</b>")
//	// hf now includes a "post_tags" parameter for the "title" field.
//
// Parameters:
//   - tags: A variadic list of strings representing the post-tags.
//
// Returns:
//
//	The updated es.highlightFieldType object with the "post_tags" parameter set.
func (hf highlightFieldType) PostTags(tags ...string) highlightFieldType {
	return hf.putInTheField("post_tags", tags)
}

// Type sets the "type" parameter in an es.highlightFieldType.
//
// This method specifies which highlighter implementation to use for this specific field,
// overriding the global type setting.
//
// Example usage:
//
//	hf := es.HighlightField("title").Type(HighlighterType.Plain)
//	// hf now includes a "type" parameter for the "title" field.
//
// Parameters:
//   - highlighterType: A HighlighterType.HighlighterType value representing the highlighter type.
//
// Returns:
//
//	The updated es.highlightFieldType object with the "type" parameter set.
func (hf highlightFieldType) Type(highlighterType HighlighterType.HighlighterType) highlightFieldType {
	return hf.putInTheField("type", highlighterType)
}

// FragmentSize sets the "fragment_size" parameter in an es.highlightFieldType.
//
// This method specifies the size of the highlighted fragment in characters for this
// specific field, overriding the global fragment_size setting.
//
// Example usage:
//
//	hf := es.HighlightField("content").FragmentSize(150)
//	// hf now includes a "fragment_size" parameter for the "content" field.
//
// Parameters:
//   - fragmentSize: An integer representing the fragment size in characters.
//
// Returns:
//
//	The updated es.highlightFieldType object with the "fragment_size" parameter set.
func (hf highlightFieldType) FragmentSize(fragmentSize int) highlightFieldType {
	return hf.putInTheField("fragment_size", fragmentSize)
}

// NumberOfFragments sets the "number_of_fragments" parameter in an es.highlightFieldType.
//
// This method specifies the maximum number of fragments to return for this specific field,
// overriding the global number_of_fragments setting. If set to 0, the entire field content
// is returned as a single fragment.
//
// Example usage:
//
//	hf := es.HighlightField("title").NumberOfFragments(0)
//	// hf now includes a "number_of_fragments" parameter for the "title" field.
//
// Parameters:
//   - numberOfFragments: An integer representing the maximum number of fragments.
//
// Returns:
//
//	The updated es.highlightFieldType object with the "number_of_fragments" parameter set.
func (hf highlightFieldType) NumberOfFragments(numberOfFragments int) highlightFieldType {
	return hf.putInTheField("number_of_fragments", numberOfFragments)
}

// NoMatchSize sets the "no_match_size" parameter in an es.highlightFieldType.
//
// This method specifies the amount of text to return from the beginning of the field
// if there are no matching fragments for this specific field.
//
// Example usage:
//
//	hf := es.HighlightField("content").NoMatchSize(150)
//	// hf now includes a "no_match_size" parameter for the "content" field.
//
// Parameters:
//   - noMatchSize: An integer representing the number of characters to return when no match is found.
//
// Returns:
//
//	The updated es.highlightFieldType object with the "no_match_size" parameter set.
func (hf highlightFieldType) NoMatchSize(noMatchSize int) highlightFieldType {
	return hf.putInTheField("no_match_size", noMatchSize)
}

// Order sets the "order" parameter in an es.highlightFieldType.
//
// This method specifies the order in which highlighted fragments are sorted for this
// specific field, overriding the global order setting.
//
// Example usage:
//
//	hf := es.HighlightField("content").Order("score")
//	// hf now includes an "order" parameter for the "content" field.
//
// Parameters:
//   - order: A string representing the order for highlighted fragments.
//
// Returns:
//
//	The updated es.highlightFieldType object with the "order" parameter set.
func (hf highlightFieldType) Order(order string) highlightFieldType {
	return hf.putInTheField("order", order)
}

// HighlightQuery sets the "highlight_query" parameter in an es.highlightFieldType.
//
// This method specifies an additional query to use for highlighting this specific field.
// This allows highlighting based on a different query than the main search query.
//
// Example usage:
//
//	hf := es.HighlightField("content").HighlightQuery(es.Match("content", "elasticsearch"))
//	// hf now includes a "highlight_query" parameter for the "content" field.
//
// Parameters:
//   - query: An object representing the highlight query. It can be of any type.
//
// Returns:
//
//	The updated es.highlightFieldType object with the "highlight_query" parameter set.
func (hf highlightFieldType) HighlightQuery(query any) highlightFieldType {
	if field, ok := correctType(query); ok {
		return hf.putInTheField("highlight_query", field)
	}
	return hf
}

// RequireFieldMatch sets the "require_field_match" parameter in an es.highlightFieldType.
//
// This method specifies whether only this field should be highlighted if it matches the query,
// overriding the global require_field_match setting.
//
// Example usage:
//
//	hf := es.HighlightField("content").RequireFieldMatch(false)
//	// hf now includes a "require_field_match" parameter for the "content" field.
//
// Parameters:
//   - requireFieldMatch: A boolean indicating whether to require field match for highlighting.
//
// Returns:
//
//	The updated es.highlightFieldType object with the "require_field_match" parameter set.
func (hf highlightFieldType) RequireFieldMatch(requireFieldMatch bool) highlightFieldType {
	return hf.putInTheField("require_field_match", requireFieldMatch)
}

// Fragmenter sets the "fragmenter" parameter in an es.highlightFieldType.
//
// This method specifies how text should be broken up into highlight fragments for this
// specific field, overriding the global fragmenter setting.
//
// Example usage:
//
//	hf := es.HighlightField("content").Fragmenter(Fragmenter.Span)
//	// hf now includes a "fragmenter" parameter for the "content" field.
//
// Parameters:
//   - fragmenter: A Fragmenter.Fragmenter value representing the fragmenter type.
//
// Returns:
//
//	The updated es.highlightFieldType object with the "fragmenter" parameter set.
func (hf highlightFieldType) Fragmenter(fragmenter Fragmenter.Fragmenter) highlightFieldType {
	return hf.putInTheField("fragmenter", fragmenter)
}

// MatchedFields sets the "matched_fields" parameter in an es.highlightFieldType.
//
// This method specifies the fields to combine matches from for highlighting. This is
// useful when using the fvh highlighter to highlight a field based on matches from
// multiple fields.
//
// Example usage:
//
//	hf := es.HighlightField("content").MatchedFields("content", "content.plain")
//	// hf now includes a "matched_fields" parameter for the "content" field.
//
// Parameters:
//   - fields: A variadic list of strings representing the fields to combine matches from.
//
// Returns:
//
//	The updated es.highlightFieldType object with the "matched_fields" parameter set.
func (hf highlightFieldType) MatchedFields(fields ...string) highlightFieldType {
	return hf.putInTheField("matched_fields", fields)
}

func (hf highlightFieldType) putInTheField(key string, value any) highlightFieldType {
	return genericPutInTheFieldOfFirstObject(hf, key, value)
}
