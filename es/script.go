package es

import ScriptLanguage "github.com/Trendyol/es-query-builder/es/enums/script-language"

type scriptType Object

type optionsType GenericObject[string]

// ScriptSource creates a new es.scriptType object with the specified script source and language.
//
// This function initializes an es.scriptType object, defining a script with the provided source code
// and script language. It is used to construct script-based queries in Elasticsearch.
//
// Example usage:
//
//	script := es.ScriptSource("return doc['price'].value * params.factor;", es.ScriptLanguage.Painless)
//	// script now contains an es.scriptType object with a Painless script that multiplies the 'price' field by a parameterized factor.
//
// Parameters:
//   - source: A string representing the script source code.
//   - scriptLanguage: The language in which the script is written, specified using es.ScriptLanguage.
//
// Returns:
//
//	An es.scriptType object containing the defined script.
func ScriptSource(source string, scriptLanguage ScriptLanguage.ScriptLanguage) scriptType {
	return scriptType{
		"lang":   scriptLanguage,
		"source": source,
	}
}

// ScriptID creates a new es.scriptType object with the specified script ID and language.
//
// This function initializes an es.scriptType object that references a stored script by its ID,
// along with the specified script language. It is used when executing pre-defined scripts in Elasticsearch.
//
// Example usage:
//
//	script := es.ScriptID("calculate-discount", es.ScriptLanguage.Painless)
//	// script now contains an es.scriptType object referencing the "calculate-discount" script written in Painless.
//
// Parameters:
//   - id: A string representing the unique identifier of the stored script.
//   - scriptLanguage: The language in which the script is written, specified using es.ScriptLanguage.
//
// Returns:
//
//	An es.scriptType object referencing the stored script.
func ScriptID(id string, scriptLanguage ScriptLanguage.ScriptLanguage) scriptType {
	return scriptType{
		"id":   id,
		"lang": scriptLanguage,
	}
}

// Option sets an additional option in the script configuration.
//
// This method adds or updates a key-value pair in the "options" field of the scriptType object.
// It is used to configure optional script settings, such as caching behavior or execution parameters.
//
// Example usage:
//
//	script := es.ScriptID("calculate-discount", es.ScriptLanguage.Painless).Option("cache", "true")
//	// script now has an "options" field with {"cache": "true"}.
//
// Parameters:
//   - option: A string representing the option key to be set.
//   - value: A string representing the value to be assigned to the option key.
//
// Returns:
//
//	The updated es.scriptType object with the specified option set.
func (s scriptType) Option(option, value string) scriptType {
	options, ok := s["options"].(optionsType)
	if !ok {
		options = optionsType{}
	}
	options[option] = value
	s["options"] = options
	return s
}

// Parameter sets a script parameter in the script configuration.
//
// This method adds or updates a key-value pair in the "params" field of the scriptType object.
// It is used to pass external parameters to the script, making it more dynamic and reusable.
//
// Example usage:
//
//	script := es.ScriptID("calculate-discount", es.ScriptLanguage.Painless).
//		Parameter("factor", 0.9)
//	// script now has a "params" field with {"factor": 0.9}.
//
// Parameters:
//   - parameter: A string representing the parameter name.
//   - value: A generic value (any type) assigned to the parameter.
//
// Returns:
//
//	The updated es.scriptType object with the specified parameter set.
func (s scriptType) Parameter(parameter string, value any) scriptType {
	params, ok := s["params"].(Object)
	if !ok {
		params = Object{}
	}
	params[parameter] = value
	s["params"] = params
	return s
}
