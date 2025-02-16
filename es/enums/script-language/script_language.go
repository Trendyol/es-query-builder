package scriptlanguage

// ScriptLanguage represents the different scripting languages supported by Elasticsearch.
//
// ScriptLanguage is a string type used to specify the language of a script in Elasticsearch queries.
// It defines various options for scripting, allowing users to perform advanced calculations,
// filtering, and custom scoring within search queries.
//
// Example usage:
//
//	var lang ScriptLanguage = Painless
//
//	// Use lang in a script configuration
//
// Constants:
//   - Painless: The default and most efficient scripting language in Elasticsearch.
//   - Expression: A lightweight, fast scripting language for numeric calculations.
//   - Mustache: A template-based scripting language used for rendering responses.
//   - Java: A scripting option that allows Java code execution within queries.
type ScriptLanguage string

const (
	// Painless is the default scripting language in Elasticsearch, optimized for performance and security.
	Painless ScriptLanguage = "painless"

	// Expression is a simple and fast scripting language for numeric computations.
	Expression ScriptLanguage = "expression"

	// Mustache is a templating language used for response rendering in Elasticsearch.
	Mustache ScriptLanguage = "mustache"

	// Java allows executing Java-based scripts in Elasticsearch.
	Java ScriptLanguage = "java"
)

// String returns the string representation of the ScriptLanguage.
func (scriptLanguage ScriptLanguage) String() string {
	return string(scriptLanguage)
}
