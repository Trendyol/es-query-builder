package es

import (
	DistanceType "github.com/Trendyol/es-query-builder/es/enums/distance-type"
	ValidationMethod "github.com/Trendyol/es-query-builder/es/enums/validation-method"
)

type geoDistanceType Object

// GeoDistance creates a new es.geoDistanceType object that matches documents
// within a given distance from a geo point.
//
// Example usage:
//
//	g := es.GeoDistance("pin.location", 40.0, -70.0, "12km")
//	// g now contains an es.geoDistanceType matching docs within 12km of the point.
//
// Parameters:
//   - field: The geo_point field name.
//   - lat: Latitude of the center point.
//   - lon: Longitude of the center point.
//   - distance: Distance string (e.g. "12km", "5mi").
//
// Returns:
//
//	An es.geoDistanceType object containing the specified geo_distance query.
func GeoDistance(field string, lat float64, lon float64, distance string) geoDistanceType {
	return geoDistanceType{
		"geo_distance": Object{
			"distance": distance,
			field: Object{
				"lat": lat,
				"lon": lon,
			},
		},
	}
}

// DistanceType sets the "distance_type" parameter in an es.geoDistanceType query.
//
// Example usage:
//
//	g := es.GeoDistance("pin.location", 40.0, -70.0, "12km").
//	    DistanceType(DistanceType.Arc)
//
// Parameters:
//   - distanceType: How distances are calculated (arc or plane).
//
// Returns:
//
//	The updated es.geoDistanceType object with the "distance_type" parameter set.
func (g geoDistanceType) DistanceType(distanceType DistanceType.DistanceType) geoDistanceType {
	return g.putInTheField("distance_type", distanceType)
}

// ValidationMethod sets the "validation_method" parameter in an es.geoDistanceType query.
//
// Example usage:
//
//	g := es.GeoDistance("pin.location", 40.0, -70.0, "12km").
//	    ValidationMethod(ValidationMethod.Strict)
//
// Parameters:
//   - validationMethod: How geo coordinates are validated.
//
// Returns:
//
//	The updated es.geoDistanceType object with the "validation_method" parameter set.
func (g geoDistanceType) ValidationMethod(validationMethod ValidationMethod.ValidationMethod) geoDistanceType {
	return g.putInTheField("validation_method", validationMethod)
}

// IgnoreUnmapped sets the "ignore_unmapped" parameter in an es.geoDistanceType query.
//
// Example usage:
//
//	g := es.GeoDistance("pin.location", 40.0, -70.0, "12km").IgnoreUnmapped(true)
//
// Parameters:
//   - ignoreUnmapped: Whether to ignore unmapped fields.
//
// Returns:
//
//	The updated es.geoDistanceType object with the "ignore_unmapped" parameter set.
func (g geoDistanceType) IgnoreUnmapped(ignoreUnmapped bool) geoDistanceType {
	return g.putInTheField("ignore_unmapped", ignoreUnmapped)
}

// Boost sets the "boost" parameter in an es.geoDistanceType query.
//
// Example usage:
//
//	g := es.GeoDistance("pin.location", 40.0, -70.0, "12km").Boost(1.5)
//
// Parameters:
//   - boost: A float64 value representing the boost factor.
//
// Returns:
//
//	The updated es.geoDistanceType object with the "boost" parameter set.
func (g geoDistanceType) Boost(boost float64) geoDistanceType {
	return g.putInTheField("boost", boost)
}

// Name sets the "_name" parameter in an es.geoDistanceType query.
//
// Example usage:
//
//	g := es.GeoDistance("pin.location", 40.0, -70.0, "12km").Name("nearby")
//
// Parameters:
//   - name: A custom name for the query (useful for debugging).
//
// Returns:
//
//	The updated es.geoDistanceType object with the "_name" parameter set.
func (g geoDistanceType) Name(name string) geoDistanceType {
	return g.putInTheField("_name", name)
}

func (g geoDistanceType) putInTheField(key string, value any) geoDistanceType {
	return genericPutInTheField(g, "geo_distance", key, value)
}
