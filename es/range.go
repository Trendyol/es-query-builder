package es

import RangeRelation "github.com/Trendyol/es-query-builder/es/enums/range-relation"

type rangeType Object

// Range creates a new es.rangeType object with the specified field.
//
// This function initializes an es.rangeType object for specifying range queries. The key represents
// the field name, and the es.rangeType object is used to define the range conditions for that field.
//
// Example usage:
//
//	r := es.Range("age")
//	// r now contains an es.rangeType object with the specified field "age" for range queries.
//
// Parameters:
//   - key: A string representing the field name for the range query.
//
// Returns:
//
//	An es.rangeType object with the specified field ready for defining range conditions.
func Range(key string) rangeType {
	return rangeType{
		"range": Object{
			key: Object{},
		},
	}
}

// LessThan sets the "lt" (less than) field for the range query.
//
// This method specifies that the range query should match values that are less than
// the provided value. It removes any existing "lte" (less than or equal to) field to ensure
// that only one type of range condition is applied.
//
// Example usage:
//
//	r := es.Range("age").LessThan(20)
//	// r now has an "lt" field set to 20 in the range query for the "age" field.
//
// Parameters:
//   - lt: The value that the field should be less than.
//
// Returns:
//
//	The updated es.rangeType object with the "lt" field set to the specified value.
func (r rangeType) LessThan(lt any) rangeType {
	return r.putInTheField("lt", lt).delete("lte")
}

// LessThanOrEqual sets the "lte" (less than or equal to) field for the range query.
//
// This method specifies that the range query should match values that are less than or equal
// to the provided value. It removes any existing "lt" (less than) field to ensure that only
// one type of range condition is applied.
//
// Example usage:
//
//	r := es.Range("age").LessThanOrEqual(20)
//	// r now has an "lte" field set to 20 in the range query for the "age" field.
//
// Parameters:
//   - lte: The value that the field should be less than or equal to.
//
// Returns:
//
//	The updated es.rangeType object with the "lte" field set to the specified value.
func (r rangeType) LessThanOrEqual(lte any) rangeType {
	return r.putInTheField("lte", lte).delete("lt")
}

// GreaterThan sets the "gt" (greater than) field for the range query.
//
// This method specifies that the range query should match values that are greater than
// the provided value. It removes any existing "gte" (greater than or equal to) field
// to ensure that only one type of range condition is applied.
//
// Example usage:
//
//	r := es.Range("age").GreaterThan(50)
//	// r now has a "gt" field set to 50 in the range query for the "age" field.
//
// Parameters:
//   - gt: The value that the field should be greater than.
//
// Returns:
//
//	The updated es.rangeType object with the "gt" field set to the specified value.
func (r rangeType) GreaterThan(gt any) rangeType {
	return r.putInTheField("gt", gt).delete("gte")
}

// GreaterThanOrEqual sets the "gte" (greater than or equal to) field for the range query.
//
// This method specifies that the range query should match values that are greater than or equal
// to the provided value. It removes any existing "gt" (greater than) field to ensure that only
// one type of range condition is applied.
//
// Example usage:
//
//	r := es.Range("age").GreaterThanOrEqual(50)
//	// r now has a "gte" field set to 50 in the range query for the "age" field.
//
// Parameters:
//   - gte: The value that the field should be greater than or equal to.
//
// Returns:
//
//	The updated es.rangeType object with the "gte" field set to the specified value.
func (r rangeType) GreaterThanOrEqual(gte any) rangeType {
	return r.putInTheField("gte", gte).delete("gt")
}

// Format sets the "format" field for the range query.
//
// This method specifies a format for the range query values, which can be useful for
// controlling how date or numeric values are interpreted. It applies the format to
// all fields in the range query object.
//
// Example usage:
//
//	r := es.Range("date").Format("yyyy-MM-dd")
//	// r now has a "format" field set to "yyyy-MM-dd" in the range query for the "date" field.
//
// Parameters:
//   - format: A string representing the format to be applied to the range query values.
//
// Returns:
//
//	The updated es.rangeType object with the "format" field set to the specified value.
func (r rangeType) Format(format string) rangeType {
	return r.putInTheField("format", format)
}

// Boost sets the "boost" field for the range query.
//
// This method applies a boost factor to the range query, influencing the relevance scoring
// of documents that match the query. It applies the boost to all fields in the range query object.
//
// Example usage:
//
//	r := es.Range("age").Boost(1.5)
//	// r now has a "boost" field set to 1.5 in the range query for the "age" field.
//
// Parameters:
//   - boost: A float64 value representing the boost factor to be applied to the range query.
//
// Returns:
//
//	The updated es.rangeType object with the "boost" field set to the specified value.
func (r rangeType) Boost(boost float64) rangeType {
	return r.putInTheField("boost", boost)
}

// From sets the "from" field for the range query.
//
// This method specifies the lower bound for the range query. The "from" field defines
// the start of the range, and documents with values greater than or equal to this
// value will be considered a match.
//
// Example usage:
//
//	r := es.Range("age").From(18)
//	// r now has a "from" field set to 18 in the range query for the "age" field.
//
// Parameters:
//   - from: The value representing the lower bound of the range. It can be of any type
//     that is valid for the field (e.g., integer, string, date).
//
// Returns:
//
//	The updated es.rangeType object with the "from" field set to the specified value.
func (r rangeType) From(from any) rangeType {
	return r.putInTheField("from", from)
}

// To sets the "to" field for the range query.
//
// This method specifies the upper bound for the range query. The "to" field defines
// the end of the range, and documents with values less than or equal to this
// value will be considered a match.
//
// Example usage:
//
//	r := es.Range("age").To(65)
//	// r now has a "to" field set to 65 in the range query for the "age" field.
//
// Parameters:
//   - to: The value representing the upper bound of the range. It can be of any type
//     that is valid for the field (e.g., integer, string, date).
//
// Returns:
//
//	The updated es.rangeType object with the "to" field set to the specified value.
func (r rangeType) To(to any) rangeType {
	return r.putInTheField("to", to)
}

// Relation sets the "relation" field for the range query.
//
// This method specifies the relationship between the ranges in the query. It allows
// you to define how the "from" and "to" values are related, such as whether one
// range is within another, contains it, or intersects with it.
//
// Example usage:
//
//	r := es.Range("age").Relation(es.RangeRelation.Within)
//	// r now has a "relation" field set to "within" in the range query for the "age" field.
//
// Parameters:
//   - relation: The RangeRelation value representing the relationship between the ranges.
//     It can be one of the following values:
//   - Within
//   - Contains
//   - Intersects
//
// Returns:
//
//	The updated es.rangeType object with the "relation" field set to the specified value.
func (r rangeType) Relation(relation RangeRelation.RangeRelation) rangeType {
	return r.putInTheField("relation", relation)
}

func (r rangeType) putInTheField(key string, value any) rangeType {
	return genericPutInTheFieldOfFirstChild(r, "range", key, value)
}

func (r rangeType) delete(key string) rangeType {
	if rang, ok := r["range"].(Object); ok {
		for field := range rang {
			if fieldObject, foOk := rang[field].(Object); foOk {
				delete(fieldObject, key)
			}
		}
	}
	return r
}
