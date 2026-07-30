package es

type suggestType Object

// Suggest creates a new es.suggestType object for the search suggest API.
//
// Example usage:
//
//	suggest := es.Suggest().
//		Text("nike shos").
//		Suggester("product_suggest",
//			es.CompletionSuggester("name.suggest").Size(5).SkipDuplicates(true),
//		)
//	query := es.NewQuery(es.MatchAll()).Suggest(suggest)
//
// Returns:
//
//	An es.suggestType object ready for further configuration.
func Suggest() suggestType {
	return suggestType{}
}

// Text sets the global suggest text shared by all suggesters unless overridden.
//
// Example usage:
//
//	suggest := es.Suggest().Text("nike shos")
//
// Parameters:
//   - text: The text to get suggestions for.
//
// Returns:
//
//	The updated es.suggestType object with the "text" field set.
func (s suggestType) Text(text string) suggestType {
	s["text"] = text
	return s
}

// Suggester adds a named suggester configuration to the suggest object.
//
// Example usage:
//
//	suggest := es.Suggest().
//		Suggester("my_suggest", es.TermSuggester("message"))
//
// Parameters:
//   - name: The name of the suggester.
//   - suggester: A term, phrase, or completion suggester configuration.
//
// Returns:
//
//	The updated es.suggestType object with the named suggester added.
func (s suggestType) Suggester(name string, suggester any) suggestType {
	s[name] = suggester
	return s
}
