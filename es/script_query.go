package es

type scriptQueryType Object

// ScriptQuery creates a new es.scriptQueryType object with the specified script.
//
// This function initializes an es.scriptQueryType object for a script query, where
// the provided script is used to perform custom computations on document fields
// during query execution. Script queries are useful for implementing advanced filtering
// logic that cannot be achieved with standard query types.
//
// Example usage:
//
//	 s := es.ScriptQuery(
//	     es.ScriptSource("doc['price'].value > params.threshold", ScriptLanguage.Painless).Parameter("threshold", 100),
//	 )
//		// s now contains an es.scriptQueryType object that executes the provided script.
//
// Parameters:
//   - script: A scriptType object representing the script to be executed in the query.
//
// Returns:
//
//	An es.scriptQueryType object containing the specified script query.
func ScriptQuery(script scriptType) scriptQueryType {
	return scriptQueryType{
		"script": Object{
			"script": script,
		},
	}
}

// Boost sets the boost value for the script query and returns the updated query object.
//
// This function allows modifying the relevance score of documents matched by the script query
// by applying a boost factor. Higher boost values increase the importance of the query
// relative to other queries in a compound query.
//
// Example usage:
//
//	sq := es.ScriptQuery(script).Boost(2.0)
//	// sq now contains a script query with a boost factor of 2.0, increasing its influence.
//
// Parameters:
//   - boost: A float64 value representing the boost factor for the script query.
//
// Returns:
//
//	A new es.scriptQueryType object with the specified boost value applied.
func (sq scriptQueryType) Boost(boost float64) scriptQueryType {
	return sq.putInTheField("boost", boost)
}

// Name sets a custom name for the script query and returns the updated query object.
//
// This function assigns a user-defined name to the script query, which can be useful for debugging,
// profiling, and identifying specific queries in complex search requests. The assigned name
// does not affect query execution but helps in logging and response analysis.
//
// Example usage:
//
//	sq := es.ScriptQuery(script).Name("custom_script_query")
//	// sq now contains a script query with the name "custom_script_query" for identification.
//
// Parameters:
//   - name: A string representing the custom name to be assigned to the script query.
//
// Returns:
//
//	A new es.scriptQueryType object with the specified name applied.
func (sq scriptQueryType) Name(name string) scriptQueryType {
	return sq.putInTheField("_name", name)
}

func (sq scriptQueryType) putInTheField(key string, value any) scriptQueryType {
	return genericPutInTheField(sq, "script", key, value)
}
