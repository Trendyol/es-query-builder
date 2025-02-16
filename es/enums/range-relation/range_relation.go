package rangerelation

// RangeRelation represents the relationship between two ranges in a query.
//
// RangeRelation is a string type that specifies how a range interacts with another range.
// It is commonly used in range queries, especially in geospatial or numeric data contexts.
//
// Constants:
//   - Within: Specifies that the range is entirely within another range.
//   - Contains: Specifies that the range entirely contains another range.
//   - Intersects: Specifies that the range overlaps or intersects with another range.
//
// Example usage:
//
//	var relation RangeRelation = Within
//
//	// Use relation in a range query configuration.
type RangeRelation string

const (
	// Within specifies that the range is entirely within another range.
	Within RangeRelation = "within"

	// Contains specifies that the range entirely contains another range.
	Contains RangeRelation = "contains"

	// Intersects specifies that the range overlaps or intersects with another range.
	Intersects RangeRelation = "intersects"
)

func (rangeRelation RangeRelation) String() string {
	return string(rangeRelation)
}
