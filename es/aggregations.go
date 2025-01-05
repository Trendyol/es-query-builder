package es

import Order "github.com/Trendyol/es-query-builder/es/enums/sort/order"

type aggsType Object

type aggTermType Object

// AggTerm creates a new aggregation term with the specified field.
//
// This function initializes an aggregation term with the given field name.
// It can be used to specify a field for aggregation operations in queries.
//
// Example usage:
//
//	termAgg := AggTerm("fieldName")
//	// termAgg now has the "field" set to "fieldName".
//
// Parameters:
//   - field: The name of the field to aggregate on.
//
// Returns:
//
//	An es.aggTermType object with the "field" set to the provided value.
func AggTerm(field string) aggTermType {
	return aggTermType{
		"field": field,
	}
}

// Missing sets the "missing" value for an aggregation term.
//
// This method specifies a value to be used when the field is missing in documents.
// It updates the es.aggTermType object to handle missing values in the aggregation.
//
// Example usage:
//
//	termAgg := AggTerm("fieldName").Missing("N/A")
//	// termAgg now has the "missing" field set to "N/A".
//
// Parameters:
//   - missing: The value to use when the field is missing.
//
// Returns:
//
//	The updated es.aggTermType object with the "missing" field set to the specified value.
func (aggTerm aggTermType) Missing(missing string) aggTermType {
	aggTerm["missing"] = missing
	return aggTerm
}

// AggTerms creates a new "terms" aggregation.
//
// This function initializes an aggregation for terms. It can be used to perform
// aggregation based on the unique terms of a field.
//
// Example usage:
//
//	termsAgg := AggTerms()
//	// termsAgg now has the "terms" field initialized.
//
// Returns:
//
//	An es.aggsType object with the "terms" field initialized.
func AggTerms() aggsType {
	return aggsType{
		"terms": Object{},
	}
}

// AggMultiTerms creates a new "multi_terms" aggregation.
//
// This function initializes an aggregation for multiple terms. It can be used
// to perform aggregation based on multiple fields or term combinations.
//
// Example usage:
//
//	multiTermsAgg := AggMultiTerms()
//	// multiTermsAgg now has the "multi_terms" field initialized.
//
// Returns:
//
//	An es.aggsType object with the "multi_terms" field initialized.
func AggMultiTerms() aggsType {
	return aggsType{
		"multi_terms": Object{},
	}
}

// AggNested creates a new "nested" aggregation.
//
// This function initializes an aggregation for nested fields. It can be used to
// perform aggregations on fields that are within a nested object.
//
// Example usage:
//
//	nestedAgg := AggNested()
//	// nestedAgg now has the "nested" field initialized.
//
// Returns:
//
//	An es.aggsType object with the "nested" field initialized.
func AggNested() aggsType {
	return aggsType{
		"nested": Object{},
	}
}

// AggMax creates a new "max" aggregation.
//
// This function initializes an aggregation to calculate the maximum value of a field.
//
// Example usage:
//
//	maxAgg := AggMax()
//	// maxAgg now has the "max" field initialized.
//
// Returns:
//
//	An es.aggsType object with the "max" field initialized.
func AggMax() aggsType {
	return aggsType{
		"max": Object{},
	}
}

// AggMin creates a new "min" aggregation.
//
// This function initializes an aggregation to calculate the minimum value of a field.
//
// Example usage:
//
//	minAgg := AggMin()
//	// minAgg now has the "min" field initialized.
//
// Returns:
//
//	An es.aggsType object with the "min" field initialized.
func AggMin() aggsType {
	return aggsType{
		"min": Object{},
	}
}

// AggAvg creates a new "avg" aggregation.
//
// This function initializes an aggregation to calculate the average value of a field.
//
// Example usage:
//
//	avgAgg := AggAvg()
//	// avgAgg now has the "avg" field initialized.
//
// Returns:
//
//	An es.aggsType object with the "avg" field initialized.
func AggAvg() aggsType {
	return aggsType{
		"avg": Object{},
	}
}

// AggCustom creates a custom aggregation with the provided aggregation object.
//
// This function initializes an aggregation based on the given custom aggregation definition.
//
// Example usage:
//
//	customAgg := AggCustom(Object{"custom": "value"})
//	// customAgg now has the custom aggregation specified.
//
// Parameters:
//   - agg: An es.Object representing a custom aggregation definition.
//
// Returns:
//
//	An es.aggsType object initialized with the provided custom aggregation.
func AggCustom(agg Object) aggsType {
	return aggsType(agg)
}

// Aggs adds a nested aggregation to the es.aggsType object.
//
// This method adds a nested aggregation under the "aggs" field with the given name.
//
// Example usage:
//
//	nestedAgg := AggTerms().Size(5)
//	agg := AggTerms().Aggs("nested", nestedAgg)
//	// agg now has a nested aggregation named "nested" with the specified aggregation.
//
// Parameters:
//   - name: The name of the nested aggregation.
//   - nestedAgg: The nested aggregation to add.
//
// Returns:
//
//	The updated es.aggsType object with the nested aggregation added.
func (agg aggsType) Aggs(name string, nestedAgg aggsType) aggsType {
	aggs, ok := agg["aggs"].(Object)
	if !ok {
		aggs = Object{}
	}
	aggs[name] = nestedAgg
	agg["aggs"] = aggs
	return agg
}

// Field sets the "field" value in the es.aggsType object.
//
// This method specifies the field to aggregate on in the es.aggsType object.
//
// Example usage:
//
//	agg := AggTerms().Field("fieldName")
//	// agg now has the "field" set to "fieldName".
//
// Parameters:
//   - field: The name of the field to aggregate on.
//
// Returns:
//
//	The updated es.aggsType object with the "field" set to the specified value.
func (agg aggsType) Field(field string) aggsType {
	return agg.putInTheField("field", field)
}

// Path sets the "path" value in the es.aggsType object.
//
// This method specifies the nested path for the aggregation in the es.aggsType object.
//
// Example usage:
//
//	agg := AggNested().Path("nestedField.path")
//	// agg now has the "path" set to "nestedField.path".
//
// Parameters:
//   - path: The nested path to use for the aggregation.
//
// Returns:
//
//	The updated es.aggsType object with the "path" set to the specified value.
func (agg aggsType) Path(path string) aggsType {
	return agg.putInTheField("path", path)
}

// Size sets the "size" value in the es.aggsType object.
//
// This method specifies the number of terms to return in the aggregation result.
//
// Example usage:
//
//	agg := AggTerms().Size(10)
//	// agg now has the "size" field set to 10.
//
// Parameters:
//   - size: The number of terms to return.
//
// Returns:
//
//	The updated es.aggsType object with the "size" field set to the specified value.
func (agg aggsType) Size(size int) aggsType {
	return agg.putInTheField("size", size)
}

// Order sets the "order" field in the es.aggsType object.
//
// This method specifies the sorting order for the aggregation results.
//
// Example usage:
//
//	agg := AggTerms().Order("fieldName", Order.Desc)
//	// agg now has the "order" field set to "desc" for "fieldName".
//
// Parameters:
//   - field: The name of the field to sort by.
//   - order: The Order value specifying the sorting direction (e.g., Asc or Desc).
//
// Returns:
//
//	The updated es.aggsType object with the "order" field set to the specified value.
func (agg aggsType) Order(field string, order Order.Order) aggsType {
	return agg.putInTheField("order",
		Object{
			field: order,
		},
	)
}

// Include sets the "include" field in the es.aggsType object.
//
// This method specifies a pattern to include in the aggregation results.
//
// Example usage:
//
//	agg := AggTerms().Include("pattern*")
//	// agg now has the "include" field set to "pattern*".
//
// Parameters:
//   - include: The pattern to include in the aggregation results.
//
// Returns:
//
//	The updated es.aggsType object with the "include" field set to the specified value.
func (agg aggsType) Include(include string) aggsType {
	return agg.putInTheField("include", include)
}

// Exclude sets the "exclude" field in the es.aggsType object.
//
// This method specifies a pattern to exclude from the aggregation results.
//
// Example usage:
//
//	agg := AggTerms().Exclude("pattern*")
//	// agg now has the "exclude" field set to "pattern*".
//
// Parameters:
//   - exclude: The pattern to exclude from the aggregation results.
//
// Returns:
//
//	The updated es.aggsType object with the "exclude" field set to the specified value.
func (agg aggsType) Exclude(exclude string) aggsType {
	return agg.putInTheField("exclude", exclude)
}

// Terms sets the "terms" field in the aggsType object.
//
// This method adds a list of aggregation terms to the "terms" field of the es.aggsType object.
// It allows specifying multiple term aggregations for the aggregation query.
//
// Example usage:
//
//	agg := AggTerms().
//		Terms(
//			AggTerm("field1"),
//			AggTerm("field2"),
//		)
//	// agg now has the "terms" field containing the provided term aggregations.
//
// Parameters:
//   - terms: A variadic list of es.aggTermType objects representing the term aggregations.
//
// Returns:
//
//	The updated es.aggsType object with the "terms" field set to the provided term aggregations.
func (agg aggsType) Terms(terms ...aggTermType) aggsType {
	return agg.putInTheField("terms", terms)
}

// Aggs adds a named aggregation to the "aggs" field of the es.Object.
//
// This method allows adding a nested aggregation under the "aggs" field in the es.Object.
// It associates the given name with the specified aggregation, enabling complex aggregation queries.
//
// Example usage:
//
//	termAgg := AggTerms().Field("fieldName")
//	query := es.NewQuery().Aggs("myAgg", termAgg)
//	// query now has an "aggs" field with a nested aggregation named "myAgg".
//
// Parameters:
//   - name: The name to associate with the nested aggregation.
//   - agg: The es.aggsType object representing the nested aggregation.
//
// Returns:
//
//	The updated Object with the "aggs" field containing the new named aggregation.
func (o Object) Aggs(name string, agg aggsType) Object {
	aggs, ok := o["aggs"].(Object)
	if !ok {
		aggs = Object{}
	}
	aggs[name] = agg
	o["aggs"] = aggs
	return o
}

func (agg aggsType) putInTheField(key string, value any) aggsType {
	for _, fieldObj := range agg {
		if fieldObject, ok := fieldObj.(Object); ok {
			fieldObject[key] = value
			break
		}
	}
	return agg
}
