package es

import (
	ValidationMethod "github.com/Trendyol/es-query-builder/es/enums/validation-method"
)

type geoBoundingBoxType Object

// GeoBoundingBox creates a new es.geoBoundingBoxType object that matches documents
// with geo points within a bounding box defined by top-left and bottom-right corners.
//
// Example usage:
//
//	g := es.GeoBoundingBox("pin.location", 40.73, -74.1, 40.01, -71.12)
//
// Parameters:
//   - field: The geo_point field name.
//   - topLeftLat: Latitude of the top-left corner.
//   - topLeftLon: Longitude of the top-left corner.
//   - bottomRightLat: Latitude of the bottom-right corner.
//   - bottomRightLon: Longitude of the bottom-right corner.
//
// Returns:
//
//	An es.geoBoundingBoxType object containing the specified geo_bounding_box query.
func GeoBoundingBox(field string, topLeftLat, topLeftLon, bottomRightLat, bottomRightLon float64) geoBoundingBoxType {
	return geoBoundingBoxType{
		"geo_bounding_box": Object{
			field: Object{
				"top_left": Object{
					"lat": topLeftLat,
					"lon": topLeftLon,
				},
				"bottom_right": Object{
					"lat": bottomRightLat,
					"lon": bottomRightLon,
				},
			},
		},
	}
}

// ValidationMethod sets the "validation_method" parameter in an es.geoBoundingBoxType query.
//
// Example usage:
//
//	g := es.GeoBoundingBox("pin.location", 40.73, -74.1, 40.01, -71.12).
//	    ValidationMethod(ValidationMethod.Strict)
//
// Parameters:
//   - validationMethod: How geo coordinates are validated.
//
// Returns:
//
//	The updated es.geoBoundingBoxType object with the "validation_method" parameter set.
func (g geoBoundingBoxType) ValidationMethod(validationMethod ValidationMethod.ValidationMethod) geoBoundingBoxType {
	return g.putInTheField("validation_method", validationMethod)
}

// IgnoreUnmapped sets the "ignore_unmapped" parameter in an es.geoBoundingBoxType query.
//
// Example usage:
//
//	g := es.GeoBoundingBox("pin.location", 40.73, -74.1, 40.01, -71.12).IgnoreUnmapped(true)
//
// Parameters:
//   - ignoreUnmapped: Whether to ignore unmapped fields.
//
// Returns:
//
//	The updated es.geoBoundingBoxType object with the "ignore_unmapped" parameter set.
func (g geoBoundingBoxType) IgnoreUnmapped(ignoreUnmapped bool) geoBoundingBoxType {
	return g.putInTheField("ignore_unmapped", ignoreUnmapped)
}

// Boost sets the "boost" parameter in an es.geoBoundingBoxType query.
//
// Example usage:
//
//	g := es.GeoBoundingBox("pin.location", 40.73, -74.1, 40.01, -71.12).Boost(1.5)
//
// Parameters:
//   - boost: A float64 value representing the boost factor.
//
// Returns:
//
//	The updated es.geoBoundingBoxType object with the "boost" parameter set.
func (g geoBoundingBoxType) Boost(boost float64) geoBoundingBoxType {
	return g.putInTheField("boost", boost)
}

// Name sets the "_name" parameter in an es.geoBoundingBoxType query.
//
// Example usage:
//
//	g := es.GeoBoundingBox("pin.location", 40.73, -74.1, 40.01, -71.12).Name("map_viewport")
//
// Parameters:
//   - name: A custom name for the query.
//
// Returns:
//
//	The updated es.geoBoundingBoxType object with the "_name" parameter set.
func (g geoBoundingBoxType) Name(name string) geoBoundingBoxType {
	return g.putInTheField("_name", name)
}

func (g geoBoundingBoxType) putInTheField(key string, value any) geoBoundingBoxType {
	return genericPutInTheField(g, "geo_bounding_box", key, value)
}
