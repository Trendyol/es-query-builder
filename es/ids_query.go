package es

type idsType Object

// Ids creates a new es.idsType object representing an Elasticsearch "ids" query.
//
// This function initializes an es.idsType object for an "ids" query, which allows
// you to match documents by their unique document IDs. It accepts a variadic list
// of strings, each representing a document ID to include in the query.
//
// Example usage:
//
//	q := es.Ids("1", "2", "3")
//	// q now contains an es.idsType object with an "ids" query
//	// matching documents with IDs "1", "2", or "3".
//
// Parameters:
//   - values: A variadic list of strings representing document IDs to be matched.
//
// Returns:
//
//	An es.idsType object containing the specified "ids" query.
func Ids[T ~string](values ...T) idsType {
	return idsType{
		"ids": Object{
			"values": values,
		},
	}
}

// IdsArray creates a new es.idsType object representing an Elasticsearch "ids" query,
// using a slice of strings as input.
//
// This function is an alternative to Ids and is useful when the document IDs are
// already available in a slice. It avoids the need to unpack the slice manually.
//
// Example usage:
//
//	ids := []string{"doc1", "doc2", "doc3"}
//	q := es.IdsArray(ids)
//	// q now contains an es.idsType object with an "ids" query
//	// matching documents with IDs "doc1", "doc2", or "doc3".
//
// Parameters:
//   - values: A slice of strings representing document IDs to be matched.
//
// Returns:
//
//	An es.idsType object containing the specified "ids" query.
func IdsArray[T ~string](values []T) idsType {
	return idsType{
		"ids": Object{
			"values": values,
		},
	}
}

// Boost sets the "boost" parameter in an es.idsType query.
//
// This method allows you to specify a boost factor for the "ids" query,
// which influences the relevance score of documents matching the given IDs.
// A higher boost value increases the importance of this query when scoring results.
//
// Example usage:
//
//	q := es.Ids("a", "b").Boost(1.5)
//	// q now includes a "boost" parameter set to 1.5.
//
// Parameters:
//   - boost: A float64 value representing the boost factor for the "ids" query.
//
// Returns:
//
//	The updated es.idsType object with the "boost" parameter set.
func (i idsType) Boost(boost float64) idsType {
	return i.putInTheField("boost", boost)
}

// Name sets the "_name" parameter in an es.idsType query.
//
// This method assigns a custom name to the "ids" query using the "_name" parameter.
// Named queries are useful for identifying which query clauses matched when analyzing
// the results of a search, especially in scenarios involving complex or nested queries.
//
// Example usage:
//
//	q := es.Ids("1", "2").Name("important-ids")
//	// q now includes a "_name" parameter set to "important-ids".
//
// Parameters:
//   - name: A string value that names the query clause.
//
// Returns:
//
//	The updated es.idsType object with the "_name" parameter set.
func (i idsType) Name(name string) idsType {
	return i.putInTheField("_name", name)
}

func (i idsType) putInTheField(key string, value any) idsType {
	if ids, ok := i["ids"].(Object); ok {
		ids[key] = value
	}
	return i
}
