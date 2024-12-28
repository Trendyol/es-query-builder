// Package rangerelation provides utilities for working with range relations.
//
// RangeRelation is a string type used to specify the spatial or logical relationship
// between two ranges, such as whether one range is entirely within another,
// contains another, or intersects with another.
//
// Example usage:
//
//	var relation RangeRelation = Within
//
//	// Use relation in a query configuration
//
// Constants:
//   - Within: Indicates that one range is entirely within another.
//   - Contains: Indicates that one range entirely contains another.
//   - Intersects: Indicates that the ranges overlap or intersect at least partially.
package rangerelation

// RangeRelation represents the relationship between two ranges.
//
// RangeRelation is a string type that specifies the spatial or logical relationship
// between ranges, which is commonly used in geospatial queries or similar scenarios.
type RangeRelation string

const (
	// Within indicates that one range is entirely within another.
	Within RangeRelation = "within"

	// Contains indicates that one range entirely contains another.
	Contains RangeRelation = "contains"

	// Intersects indicates that the ranges overlap or intersect at least partially.
	Intersects RangeRelation = "intersects"
)

func (rangeRelation RangeRelation) String() string {
	return string(rangeRelation)
}
