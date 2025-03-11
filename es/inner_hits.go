package es

type innerHitsType Object

type scriptFieldType Object

type scriptFieldsType GenericObject[scriptFieldType]

type fieldAndFormatType Object

type fieldCollapseType Object

// InnerHits initializes and returns a new instance of innerHitsType.
//
// This method creates an inner hits object, which is used to retrieve nested or
// parent-child documents matching the query. Inner hits allow you to include
// relevant sub-documents in the search response for better insights.
//
// Example usage:
//
//	innerHits := es.InnerHits()
//	// Use `hits` to configure and include inner hits in a query.
//
// Returns:
//
//	A new instance of es.innerHitsType, ready for further configuration.
func InnerHits() innerHitsType {
	return innerHitsType{}
}

// Collapse sets the "collapse" field for the inner hits configuration.
//
// This method specifies a field-based collapsing strategy for inner hits, allowing
// you to group or collapse results based on the values of a specified field.
//
// Example usage:
//
//	collapse := es.FieldCollapse("author")
//	innerHits := es.InnerHits().Collapse(collapse)
//	// The inner hits object now includes a "collapse" field configured for the "author" field.
//
// Parameters:
//   - fieldCollapse: An es.fieldCollapseType object defining the collapsing behavior based on a specific field.
//
// Returns:
//
//	The updated es.innerHitsType object with the "collapse" field set to the specified value.
func (ih innerHitsType) Collapse(fieldCollapse fieldCollapseType) innerHitsType {
	ih["collapse"] = fieldCollapse
	return ih
}

// DocvalueFields sets the "docvalue_fields" field for the inner hits configuration.
//
// This method specifies a list of fields and their formats to include as part of the
// inner hits results. Docvalue fields are field values stored as columnar data for
// efficient retrieval.
//
// Example usage:
//
//	innerHits := es.InnerHits().
//				DocvalueFields(
//					es.FieldAndFormatType("timestamp").Format("epoch_millis"),
//					es.FieldAndFormatType("price"),
//				)
//	// The inner hits object now includes "docvalue_fields" configured with the specified fields.
//
// Parameters:
//   - fieldAndFormat: A variadic list of `es.fieldAndFormatType` objects, each defining
//     a field and an optional format.
//
// Returns:
//
//	The updated es.innerHitsType object with the "docvalue_fields" field set to the specified values.
func (ih innerHitsType) DocvalueFields(fieldAndFormat ...fieldAndFormatType) innerHitsType {
	ih["docvalue_fields"] = fieldAndFormat
	return ih
}

// Explain sets the "explain" field for the inner hits configuration.
//
// This method enables or disables the inclusion of explanation details for scoring
// in the inner hits results. When enabled, it provides detailed information about
// how the score was computed for each document.
//
// Example usage:
//
//	innerHits := es.InnerHits().Explain(true)
//	// The inner hits object now includes an "explain" field set to true.
//
// Parameters:
//   - explain: A boolean value indicating whether to include scoring explanations.
//   - true: Include explanations for scoring in the results.
//   - false: Do not include explanations (default behavior).
//
// Returns:
//
//	The updated es.innerHitsType object with the "explain" field set to the specified value.
func (ih innerHitsType) Explain(explain bool) innerHitsType {
	ih["explain"] = explain
	return ih
}

// Fields sets the "fields" field for the inner hits configuration.
//
// This method specifies a list of fields to include in the inner hits results.
// It allows selecting specific fields to be retrieved for the matched documents,
// reducing the amount of data returned.
//
// Example usage:
//
//	innerHits := es.InnerHits().Fields("title", "author", "timestamp")
//	// The inner hits object now includes a "fields" field with the specified field names.
//
// Parameters:
//   - fields: A variadic list of strings representing the field names to include in the results.
//
// Returns:
//
//	The updated es.innerHitsType object with the "fields" field set to the specified values.
func (ih innerHitsType) Fields(fields ...string) innerHitsType {
	ih["fields"] = fields
	return ih
}

// From sets the "from" field for the inner hits configuration.
//
// This method specifies the starting index for the inner hits results, enabling pagination.
// It determines the offset from which to begin returning hits, useful when you want to skip
// a specific number of matches.
//
// Example usage:
//
//	innerHits := es.InnerHits().From(10)
//	// The inner hits object now includes a "from" field set to 10.
//
// Parameters:
//   - from: An integer representing the starting index for the results.
//
// Returns:
//
//	The updated es.innerHitsType object with the "from" field set to the specified value.
func (ih innerHitsType) From(from int) innerHitsType {
	ih["from"] = from
	return ih
}

// IgnoreUnmapped sets the "ignore_unmapped" field for the inner hits configuration.
//
// This method specifies whether to ignore unmapped fields in the inner hits query.
// If set to true, the query will not fail if a field referenced in the inner hits
// is not mapped in the index.
//
// Example usage:
//
//	innerHits := es.InnerHits().IgnoreUnmapped(true)
//	// The inner hits object now includes an "ignore_unmapped" field set to true.
//
// Parameters:
//   - ignoreUnmapped: A boolean value indicating whether to ignore unmapped fields.
//   - true: Ignore unmapped fields and prevent query failures.
//   - false: Do not ignore unmapped fields (default behavior).
//
// Returns:
//
//	The updated es.innerHitsType object with the "ignore_unmapped" field set to the specified value.
func (ih innerHitsType) IgnoreUnmapped(ignoreUnmapped bool) innerHitsType {
	ih["ignore_unmapped"] = ignoreUnmapped
	return ih
}

// Name sets the "name" field for the inner hits configuration.
//
// This method assigns a custom name to the inner hits result. The name can be used
// to identify or label the inner hits in the search response, especially when multiple
// inner hits configurations are used.
//
// Example usage:
//
//	innerHits := es.InnerHits().Name("nested_comments")
//	// The inner hits object now includes a "name" field set to "nested_comments".
//
// Parameters:
//   - name: A string representing the custom name for the inner hits.
//
// Returns:
//
//	The updated es.innerHitsType object with the "name" field set to the specified value.
func (ih innerHitsType) Name(name string) innerHitsType {
	ih["name"] = name
	return ih
}

// ScriptField adds a script field to the inner_hits configuration.
//
// This method allows you to define a dynamically computed field within the inner_hits section
// of a query. The field is generated using the specified script.
//
// Example usage:
//
//	script := es.ScriptSource("doc['price'].value * params.factor", es.ScriptLanguage.Painless)
//	ih := es.InnerHits().ScriptField("discounted_price", es.ScriptField(script))
//	// ih now contains a script field named "discounted_price" that calculates values dynamically.
//
// Parameters:
//   - name: A string representing the name of the script field.
//   - scriptField: An es.scriptFieldType object defining the script to be used.
//
// Returns:
//
//	The updated es.innerHitsType object with the specified script field added.
func (ih innerHitsType) ScriptField(name string, scriptField scriptFieldType) innerHitsType {
	scriptFields, ok := ih["script_fields"].(scriptFieldsType)
	if !ok {
		scriptFields = scriptFieldsType{}
	}
	scriptFields[name] = scriptField
	ih["script_fields"] = scriptFields
	return ih
}

// SeqNoPrimaryTerm sets the "seq_no_primary_term" field for the inner hits configuration.
//
// This method specifies whether to include the sequence number and primary term of
// matched documents in the inner hits results. These values are useful for advanced use
// cases like optimistic concurrency control.
//
// Example usage:
//
//	innerHits := es.InnerHits().SeqNoPrimaryTerm(true)
//	// The inner hits object now includes a "seq_no_primary_term" field set to true.
//
// Parameters:
//   - seqNoPrimaryTerm: A boolean value indicating whether to include the sequence number
//     and primary term in the results.
//   - true: Include the sequence number and primary term.
//   - false: Do not include these values (default behavior).
//
// Returns:
//
//	The updated es.innerHitsType object with the "seq_no_primary_term" field set to the specified value.
func (ih innerHitsType) SeqNoPrimaryTerm(seqNoPrimaryTerm bool) innerHitsType {
	ih["seq_no_primary_term"] = seqNoPrimaryTerm
	return ih
}

// Size sets the "size" field for the inner hits configuration.
//
// This method specifies the maximum number of inner hits to return. It controls the
// number of matched documents included in the inner hits results, allowing for fine-grained
// control over the result set size.
//
// Example usage:
//
//	innerHits := es.InnerHits().Size(5)
//	// The inner hits object now includes a "size" field set to 5.
//
// Parameters:
//   - size: An integer representing the maximum number of inner hits to retrieve.
//
// Returns:
//
//	The updated es.innerHitsType object with the "size" field set to the specified value.
func (ih innerHitsType) Size(size int) innerHitsType {
	ih["size"] = size
	return ih
}

// Sort sets the "sort" field for the inner hits configuration.
//
// This method specifies how to sort the inner hits results. You can provide one or more
// sorting criteria based on fields, values, or custom logic to control the order of
// the matched documents.
//
// Example usage:
//
//	innerHits := es.InnerHits().Sort(es.Sort("timestamp").Order(order.Desc))
//	// The inner hits object now includes a "sort" field with the specified sorting criteria.
//
// Parameters:
//   - sorts: A variadic list of es.sortType objects defining the sorting criteria.
//
// Returns:
//
//	The updated es.innerHitsType object with the "sort" field set to the specified values.
func (ih innerHitsType) Sort(sorts ...sortType) innerHitsType {
	ih["sort"] = sorts
	return ih
}

// SourceFalse sets the "_source" field to false for the inner hits configuration.
//
// This method disables the inclusion of the _source field in the inner hits results.
// The _source field contains the original document source, and setting it to false
// prevents it from being returned, reducing the amount of data in the response.
//
// Example usage:
//
//	innerHits := es.InnerHits().SourceFalse()
//	// The inner hits object now includes "_source" set to false.
//
// Returns:
//
//	The updated es.innerHitsType object with the "_source" field set to false.
func (ih innerHitsType) SourceFalse() innerHitsType {
	ih["_source"] = false
	return ih
}

// SourceIncludes sets the "includes" field for the _source configuration in the inner hits.
//
// This method specifies a list of fields to include in the inner hits results. Only
// the specified fields will be included in the _source field, and all other fields will be excluded.
//
// Example usage:
//
//	innerHits := es.InnerHits().SourceIncludes("title", "author")
//	// The inner hits object now includes only the "title" and "author" fields in the "_source".
//
// Parameters:
//   - fields: A variadic list of field names to include in the _source field.
//
// Returns:
//
//	The updated es.innerHitsType object with the "includes" field in _source set to the specified values.
func (ih innerHitsType) SourceIncludes(fields ...string) innerHitsType {
	if len(fields) == 0 {
		return ih
	}
	source, ok := ih["_source"].(Object)
	if !ok {
		source = Object{}
	}
	includes, ok := source["includes"].(Array)
	if !ok {
		includes = make(Array, 0, len(fields))
	}
	for i := 0; i < len(fields); i++ {
		includes = append(includes, fields[i])
	}
	source["includes"] = includes
	ih["_source"] = source
	return ih
}

// SourceExcludes sets the "excludes" field for the _source configuration in the inner hits.
//
// This method specifies a list of fields to exclude from the inner hits results. The excluded
// fields will not be included in the _source field, allowing you to limit the data returned.
//
// Example usage:
//
//	innerHits := es.InnerHits().SourceExcludes("description", "timestamp")
//	// The inner hits object now excludes the "description" and "timestamp" fields from the "_source".
//
// Parameters:
//   - fields: A variadic list of field names to exclude from the _source field.
//
// Returns:
//
//	The updated es.innerHitsType object with the "excludes" field in _source set to the specified values.
func (ih innerHitsType) SourceExcludes(fields ...string) innerHitsType {
	if len(fields) == 0 {
		return ih
	}
	source, ok := ih["_source"].(Object)
	if !ok {
		source = Object{}
	}
	excludes, exists := source["excludes"].(Array)
	if !exists {
		excludes = make(Array, 0, len(fields))
	}
	for i := 0; i < len(fields); i++ {
		excludes = append(excludes, fields[i])
	}
	source["excludes"] = excludes
	ih["_source"] = source
	return ih
}

// StoredFields sets the "stored_fields" field for the inner hits configuration.
//
// This method specifies which stored fields to retrieve for the matched documents
// in the inner hits results. Only the specified stored fields will be included in
// the response, allowing you to limit the amount of data returned.
//
// Example usage:
//
//	innerHits := es.InnerHits().StoredFields("title", "author")
//	// The inner hits object now includes a "stored_fields" field with the specified field names.
//
// Parameters:
//   - fields: A variadic list of field names to retrieve as stored fields.
//
// Returns:
//
//	The updated es.innerHitsType object with the "stored_fields" field set to the specified values.
func (ih innerHitsType) StoredFields(fields ...string) innerHitsType {
	ih["stored_fields"] = fields
	return ih
}

// TrackScores sets the "track_scores" field for the inner hits configuration.
//
// This method specifies whether to track the relevance scores for the inner hits results.
// When set to true, the scores of the inner hits will be included in the response, which
// can be useful for sorting or filtering based on relevance.
//
// Example usage:
//
//	innerHits := es.InnerHits().TrackScores(true)
//	// The inner hits object now includes a "track_scores" field set to true.
//
// Parameters:
//   - trackScores: A boolean value indicating whether to track the relevance scores.
//   - true: Track the scores for the inner hits.
//   - false: Do not track the scores (default behavior).
//
// Returns:
//
//	The updated es.innerHitsType object with the "track_scores" field set to the specified value.
func (ih innerHitsType) TrackScores(trackScores bool) innerHitsType {
	ih["track_scores"] = trackScores
	return ih
}

// Version sets the "version" field for the inner hits configuration.
//
// This method specifies whether to include the version number of the matched documents
// in the inner hits results. Enabling this can be useful for cases where you need to
// track document versions in the response.
//
// Example usage:
//
//	innerHits := es.InnerHits().Version(true)
//	// The inner hits object now includes a "version" field set to true.
//
// Parameters:
//   - version: A boolean value indicating whether to include the document version.
//   - true: Include the version number of the matched documents.
//   - false: Do not include the version (default behavior).
//
// Returns:
//
//	The updated es.innerHitsType object with the "version" field set to the specified value.
func (ih innerHitsType) Version(version bool) innerHitsType {
	ih["version"] = version
	return ih
}

// ScriptField creates a new es.scriptFieldType object with the specified script.
//
// This function initializes an es.scriptFieldType object, which is used to define a script-based field
// in an Elasticsearch query. The provided script determines the field's value dynamically at query time.
//
// Example usage:
//
//	script := es.ScriptSource("doc['price'].value * params.factor", es.ScriptLanguage.Painless)
//	sf := es.ScriptField(script)
//	// sf now contains an es.scriptFieldType object with the specified script.
//
// Parameters:
//   - script: An es.scriptType object representing the script to be executed.
//
// Returns:
//
//	An es.scriptFieldType object containing the specified script.
func ScriptField(script scriptType) scriptFieldType {
	return scriptFieldType{
		"script": script,
	}
}

// IgnoreFailure sets the "ignore_failure" field in the script field configuration.
//
// This method allows the script field to ignore failures during execution. If set to `true`,
// Elasticsearch will continue processing the query even if the script encounters an error.
//
// Example usage:
//
//	script := es.ScriptSource("doc['price'].value * params.factor", es.ScriptLanguage.Painless)
//	sf := es.ScriptField(script).IgnoreFailure(true)
//	// sf now has an "ignore_failure" field set to true.
//
// Parameters:
//   - ignoreFailure: A boolean indicating whether script execution failures should be ignored.
//
// Returns:
//
//	The updated es.scriptFieldType object with the "ignore_failure" field set.
func (sf scriptFieldType) IgnoreFailure(ignoreFailure bool) scriptFieldType {
	sf["ignore_failure"] = ignoreFailure
	return sf
}

// FieldCollapse creates a field collapse configuration for sorting or grouping inner hits.
//
// This function allows you to collapse the inner hits results based on a specific field.
// It can be useful for grouping documents by field values or limiting the number of results
// per group when working with nested or parent-child queries.
//
// Example usage:
//
//	collapse := es.FieldCollapse("category")
//	// The fieldCollapseType object now specifies that the results should be collapsed based on the "category" field.
//
// Parameters:
//   - field: A string representing the field to collapse the results by.
//
// Returns:
//
//	An es.fieldCollapseType object with the "field" set to the specified value, used for collapsing results.
func FieldCollapse(field string) fieldCollapseType {
	return fieldCollapseType{
		"field": field,
	}
}

// Collapse sets the "collapse" field for the field collapse configuration.
//
// This method allows you to apply further collapse rules to the field collapse configuration,
// such as setting multiple collapse criteria or specifying additional settings for the collapse operation.
//
// Example usage:
//
//		collapse := es.FieldCollapse("category").Collapse(es.FieldCollapse("subcategory"))
//		// The es.fieldCollapseType object now includes a "collapse" field specifying the collapse criteria
//	 // for both "category" and "subcategory".
//
// Parameters:
//   - fieldCollapse: An es.fieldCollapseType object representing the collapse configuration to apply.
//
// Returns:
//
//	The updated es.fieldCollapseType object with the "collapse" field set to the specified value.
func (fc fieldCollapseType) Collapse(fieldCollapse fieldCollapseType) fieldCollapseType {
	fc["collapse"] = fieldCollapse
	return fc
}

// InnerHits sets the "inner_hits" field for the field collapse configuration.
//
// This method allows you to specify inner hits for the collapsed field. Inner hits provide
// the matching documents from the collapsed field, allowing you to include more details about
// the nested or parent-child documents.
//
// Example usage:
//
//	collapse := es.FieldCollapse("category").InnerHits(es.InnerHits().Size(5))
//	// The es.fieldCollapseType object now includes an "inner_hits" field with the specified inner hits configuration.
//
// Parameters:
//   - innerHits: A variadic list of es.innerHitsType objects representing the inner hits configuration.
//
// Returns:
//
//	The updated es.fieldCollapseType object with the "inner_hits" field set to the specified values.
func (fc fieldCollapseType) InnerHits(innerHits ...innerHitsType) fieldCollapseType {
	fc["inner_hits"] = innerHits
	return fc
}

// MaxConcurrentGroupSearches sets the "max_concurrent_group_searches" field for the field collapse configuration.
//
// This method allows you to specify the maximum number of concurrent searches for each group of collapsed results.
// Limiting the number of concurrent searches can help optimize performance, especially when dealing with large datasets.
//
// Example usage:
//
//	collapse := es.FieldCollapse("category").MaxConcurrentGroupSearches(10)
//	// The es.fieldCollapseType object now includes a "max_concurrent_group_searches" field with the specified value.
//
// Parameters:
//   - maxConcurrentGroupSearches: An integer representing the maximum number of concurrent searches to allow per group.
//
// Returns:
//
//	The updated es.fieldCollapseType object with the "max_concurrent_group_searches" field set to the specified value.
func (fc fieldCollapseType) MaxConcurrentGroupSearches(maxConcurrentGroupSearches int) fieldCollapseType {
	fc["max_concurrent_group_searches"] = maxConcurrentGroupSearches
	return fc
}

// FieldAndFormat creates a field and format configuration for retrieving field values.
//
// This function allows you to specify a field along with its format in the search results.
// It can be useful for controlling how fields like dates or numbers are returned in the query response.
//
// Example usage:
//
//	fieldFormat := es.FieldAndFormat("date", "yyyy-MM-dd")
//	// The es.fieldAndFormatType object now specifies that the "date" field should be returned in the "yyyy-MM-dd" format.
//
// Parameters:
//   - field: A string representing the field to retrieve.
//   - format: A string representing the format to apply to the field value (optional).
//
// Returns:
//
//	An es.fieldAndFormatType object with the "field" set to the specified value and an optional "format".
func FieldAndFormat(field string) fieldAndFormatType {
	return fieldAndFormatType{
		"field": field,
	}
}

// Format sets the "format" field for the field and format configuration.
//
// This method allows you to specify the format for a field, such as date or number formatting,
// to control how the field's value will be returned in the query response.
//
// Example usage:
//
//	fieldFormat := es.FieldAndFormat("date").Format("yyyy-MM-dd")
//	// The es.fieldAndFormatType object now includes a "format" field with the specified value.
//
// Parameters:
//   - format: A string representing the format to apply to the field value.
//
// Returns:
//
//	The updated es.fieldAndFormatType object with the "format" field set to the specified value.
func (fnf fieldAndFormatType) Format(format string) fieldAndFormatType {
	fnf["format"] = format
	return fnf
}

// IncludeUnmapped sets the "include_unmapped" field for the field and format configuration.
//
// This method allows you to specify whether to include unmapped fields in the response.
// When set to true, fields that are not mapped in the index will still be included in the
// query response as null values.
//
// Example usage:
//
//	fieldFormat := es.FieldAndFormat("price").IncludeUnmapped(true)
//	// The es.fieldAndFormatType object now includes an "include_unmapped" field set to true.
//
// Parameters:
//   - includeUnmapped: A boolean value indicating whether to include unmapped fields.
//   - true: Include unmapped fields in the response.
//   - false: Do not include unmapped fields (default behavior).
//
// Returns:
//
//	The updated es.fieldAndFormatType object with the "include_unmapped" field set to the specified value.
func (fnf fieldAndFormatType) IncludeUnmapped(includeUnmapped bool) fieldAndFormatType {
	fnf["include_unmapped"] = includeUnmapped
	return fnf
}
