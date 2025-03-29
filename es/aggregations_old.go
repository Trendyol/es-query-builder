package es

import Order "github.com/Trendyol/es-query-builder/es/enums/sort/order"

type aggsTypeOld Object

type aggTermTypeOld Object

// AggTermOld creates a new aggregation term with the specified field.
//
// This function initializes an aggregation term with the given field name.
// It can be used to specify a field for aggregation operations in queries.
//
// Example usage:
//
//	termAgg := AggTermOld("fieldName")
//	// termAgg now has the "field" set to "fieldName".
//
// Parameters:
//   - field: The name of the field to aggregate on.
//
// Returns:
//
//	An es.aggTermTypeOld object with the "field" set to the provided value.
func AggTermOld(field string) aggTermTypeOld {
	return aggTermTypeOld{
		"field": field,
	}
}

// Missing sets the "missing" value for an aggregation term.
//
// This method specifies a value to be used when the field is missing in documents.
// It updates the es.aggTermTypeOld object to handle missing values in the aggregation.
//
// Example usage:
//
//	termAgg := AggTermOld("fieldName").Missing("N/A")
//	// termAgg now has the "missing" field set to "N/A".
//
// Parameters:
//   - missing: The value to use when the field is missing.
//
// Returns:
//
//	The updated es.aggTermTypeOld object with the "missing" field set to the specified value.
func (aggTerm aggTermTypeOld) Missing(missing string) aggTermTypeOld {
	aggTerm["missing"] = missing
	return aggTerm
}

// MinDocCount sets the "min_doc_count" value for an aggregation term.
//
// This method specifies the minimum number of documents that must match a term
// for it to be included in the aggregation results. It updates the es.aggTermTypeOld
// object to enforce this constraint.
//
// Example usage:
//
//	termAgg := AggTermOld("fieldName").MinDocCount(2)
//	// termAgg now has the "min_doc_count" field set to 2.
//
// Parameters:
//   - minDocCount: The minimum number of documents required for a term to be included.
//
// Returns:
//
//	The updated es.aggTermTypeOld object with the "min_doc_count" field set to the specified value.
func (aggTerm aggTermTypeOld) MinDocCount(minDocCount int) aggTermTypeOld {
	aggTerm["min_doc_count"] = minDocCount
	return aggTerm
}

// AggTermsOld creates a new "terms" aggregation.
//
// This function initializes an aggregation for terms. It can be used to perform
// aggregation based on the unique terms of a field.
//
// Example usage:
//
//	termsAgg := AggTermsOld()
//	// termsAgg now has the "terms" field initialized.
//
// Returns:
//
//	An es.aggsTypeOld object with the "terms" field initialized.
func AggTermsOld() aggsTypeOld {
	return aggsTypeOld{
		"terms": Object{},
	}
}

// AggMultiTermsOld creates a new "multi_terms" aggregation.
//
// This function initializes an aggregation for multiple terms. It can be used
// to perform aggregation based on multiple fields or term combinations.
//
// Example usage:
//
//	multiTermsAgg := AggMultiTermsOld()
//	// multiTermsAgg now has the "multi_terms" field initialized.
//
// Returns:
//
//	An es.aggsTypeOld object with the "multi_terms" field initialized.
func AggMultiTermsOld() aggsTypeOld {
	return aggsTypeOld{
		"multi_terms": Object{},
	}
}

// AggNestedOld creates a new "nested" aggregation.
//
// This function initializes an aggregation for nested fields. It can be used to
// perform aggregations on fields that are within a nested object.
//
// Example usage:
//
//	nestedAgg := AggNestedOld()
//	// nestedAgg now has the "nested" field initialized.
//
// Returns:
//
//	An es.aggsTypeOld object with the "nested" field initialized.
func AggNestedOld() aggsTypeOld {
	return aggsTypeOld{
		"nested": Object{},
	}
}

// AggMaxOld creates a new "max" aggregation.
//
// This function initializes an aggregation to calculate the maximum value of a field.
//
// Example usage:
//
//	maxAgg := AggMaxOld()
//	// maxAgg now has the "max" field initialized.
//
// Returns:
//
//	An es.aggsTypeOld object with the "max" field initialized.
func AggMaxOld() aggsTypeOld {
	return aggsTypeOld{
		"max": Object{},
	}
}

// AggMinOld creates a new "min" aggregation.
//
// This function initializes an aggregation to calculate the minimum value of a field.
//
// Example usage:
//
//	minAgg := AggMinOld()
//	// minAgg now has the "min" field initialized.
//
// Returns:
//
//	An es.aggsTypeOld object with the "min" field initialized.
func AggMinOld() aggsTypeOld {
	return aggsTypeOld{
		"min": Object{},
	}
}

// AggAvgOld creates a new "avg" aggregation.
//
// This function initializes an aggregation to calculate the average value of a field.
//
// Example usage:
//
//	avgAgg := AggAvgOld()
//	// avgAgg now has the "avg" field initialized.
//
// Returns:
//
//	An es.aggsTypeOld object with the "avg" field initialized.
func AggAvgOld() aggsTypeOld {
	return aggsTypeOld{
		"avg": Object{},
	}
}

// AggCustomOld creates a custom aggregation with the provided aggregation object.
//
// This function initializes an aggregation based on the given custom aggregation definition.
//
// Example usage:
//
//	customAgg := AggCustomOld(Object{"custom": "value"})
//	// customAgg now has the custom aggregation specified.
//
// Parameters:
//   - agg: An es.Object representing a custom aggregation definition.
//
// Returns:
//
//	An es.aggsTypeOld object initialized with the provided custom aggregation.
func AggCustomOld(agg Object) aggsTypeOld {
	return aggsTypeOld(agg)
}

// Aggs adds a nested aggregation to the es.aggsTypeOld object.
//
// This method adds a nested aggregation under the "aggs" field with the given name.
//
// Example usage:
//
//	nestedAgg := AggTermsOld().Size(5)
//	agg := AggTermsOld().Aggs("nested", nestedAgg)
//	// agg now has a nested aggregation named "nested" with the specified aggregation.
//
// Parameters:
//   - name: The name of the nested aggregation.
//   - nestedAgg: The nested aggregation to add.
//
// Returns:
//
//	The updated es.aggsTypeOld object with the nested aggregation added.
func (agg aggsTypeOld) Aggs(name string, nestedAgg aggsTypeOld) aggsTypeOld {
	aggs, ok := agg["aggs"].(Object)
	if !ok {
		aggs = Object{}
	}
	aggs[name] = nestedAgg
	agg["aggs"] = aggs
	return agg
}

// Field sets the "field" value in the es.aggsTypeOld object.
//
// This method specifies the field to aggregate on in the es.aggsTypeOld object.
//
// Example usage:
//
//	agg := AggTermsOld().Field("fieldName")
//	// agg now has the "field" set to "fieldName".
//
// Parameters:
//   - field: The name of the field to aggregate on.
//
// Returns:
//
//	The updated es.aggsTypeOld object with the "field" set to the specified value.
func (agg aggsTypeOld) Field(field string) aggsTypeOld {
	return agg.putInTheField("field", field)
}

// Path sets the "path" value in the es.aggsTypeOld object.
//
// This method specifies the nested path for the aggregation in the es.aggsTypeOld object.
//
// Example usage:
//
//	agg := AggNestedOld().Path("nestedField.path")
//	// agg now has the "path" set to "nestedField.path".
//
// Parameters:
//   - path: The nested path to use for the aggregation.
//
// Returns:
//
//	The updated es.aggsTypeOld object with the "path" set to the specified value.
func (agg aggsTypeOld) Path(path string) aggsTypeOld {
	return agg.putInTheField("path", path)
}

// Size sets the "size" value in the es.aggsTypeOld object.
//
// This method specifies the number of terms to return in the aggregation result.
//
// Example usage:
//
//	agg := AggTermsOld().Size(10)
//	// agg now has the "size" field set to 10.
//
// Parameters:
//   - size: The number of terms to return.
//
// Returns:
//
//	The updated es.aggsTypeOld object with the "size" field set to the specified value.
func (agg aggsTypeOld) Size(size int) aggsTypeOld {
	return agg.putInTheField("size", size)
}

// Order sets the "order" field in the es.aggsTypeOld object.
//
// This method specifies the sorting order for the aggregation results.
//
// Example usage:
//
//	agg := AggTermsOld().Order("fieldName", Order.Desc)
//	// agg now has the "order" field set to "desc" for "fieldName".
//
// Parameters:
//   - field: The name of the field to sort by.
//   - order: The Order value specifying the sorting direction (e.g., Asc or Desc).
//
// Returns:
//
//	The updated es.aggsTypeOld object with the "order" field set to the specified value.
func (agg aggsTypeOld) Order(field string, order Order.Order) aggsTypeOld {
	return agg.putInTheField("order",
		Object{
			field: order,
		},
	)
}

// Include sets the "include" field in the es.aggsTypeOld object.
//
// This method specifies a pattern to include in the aggregation results.
//
// Example usage:
//
//	agg := AggTermsOld().Include("pattern*")
//	// agg now has the "include" field set to "pattern*".
//
// Parameters:
//   - include: The pattern to include in the aggregation results.
//
// Returns:
//
//	The updated es.aggsTypeOld object with the "include" field set to the specified value.
func (agg aggsTypeOld) Include(include string) aggsTypeOld {
	return agg.putInTheField("include", include)
}

// Exclude sets the "exclude" field in the es.aggsTypeOld object.
//
// This method specifies a pattern to exclude from the aggregation results.
//
// Example usage:
//
//	agg := AggTermsOld().Exclude("pattern*")
//	// agg now has the "exclude" field set to "pattern*".
//
// Parameters:
//   - exclude: The pattern to exclude from the aggregation results.
//
// Returns:
//
//	The updated es.aggsTypeOld object with the "exclude" field set to the specified value.
func (agg aggsTypeOld) Exclude(exclude string) aggsTypeOld {
	return agg.putInTheField("exclude", exclude)
}

// Terms sets the "terms" field in the aggsTypeOld object.
//
// This method adds a list of aggregation terms to the "terms" field of the es.aggsTypeOld object.
// It allows specifying multiple term aggregations for the aggregation query.
//
// Example usage:
//
//	agg := AggTermsOld().
//		Terms(
//			AggTermOld("field1"),
//			AggTermOld("field2"),
//		)
//	// agg now has the "terms" field containing the provided term aggregations.
//
// Parameters:
//   - terms: A variadic list of es.aggTermTypeOld objects representing the term aggregations.
//
// Returns:
//
//	The updated es.aggsTypeOld object with the "terms" field set to the provided term aggregations.
func (agg aggsTypeOld) Terms(terms ...aggTermTypeOld) aggsTypeOld {
	return agg.putInTheField("terms", terms)
}

// Aggs adds a named aggregation to the "aggs" field of the es.Object.
//
// This method allows adding a nested aggregation under the "aggs" field in the es.Object.
// It associates the given name with the specified aggregation, enabling complex aggregation queries.
//
// Example usage:
//
//	termAgg := AggTermsOld().Field("fieldName")
//	query := es.NewQuery().Aggs("myAgg", termAgg)
//	// query now has an "aggs" field with a nested aggregation named "myAgg".
//
// Parameters:
//   - name: The name to associate with the nested aggregation.
//   - agg: The es.aggsTypeOld object representing the nested aggregation.
//
// Returns:
//
//	The updated Object with the "aggs" field containing the new named aggregation.
func (o Object) Aggs(name string, agg aggsTypeOld) Object {
	aggs, ok := o["aggs"].(Object)
	if !ok {
		aggs = Object{}
	}
	aggs[name] = agg
	o["aggs"] = aggs
	return o
}

func (agg aggsTypeOld) putInTheField(key string, value any) aggsTypeOld {
	for _, fieldObj := range agg {
		if fieldObject, ok := fieldObj.(Object); ok {
			fieldObject[key] = value
			break
		}
	}
	return agg
}
