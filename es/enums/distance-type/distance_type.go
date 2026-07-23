package distancetype

// DistanceType represents how distances are calculated for geo queries.
//
// Example usage:
//
//	es.GeoDistance("pin.location", 40.0, -70.0, "12km").DistanceType(DistanceType.Arc)
type DistanceType string

const (
	// Arc indicates that arc distance calculation should be used (default, more accurate).
	Arc DistanceType = "arc"

	// Plane indicates that plane distance calculation should be used (faster, less accurate).
	Plane DistanceType = "plane"
)

func (distanceType DistanceType) String() string {
	return string(distanceType)
}
