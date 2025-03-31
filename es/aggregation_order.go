package es

import Order "github.com/Trendyol/es-query-builder/es/enums/sort/order"

type aggOrder Object

// AggOrder creates an aggregation order specification.
//
// Aggregation orders determine the sorting of buckets based on a specific key
// (e.g., "_count" or "_key") and the desired order (ascending or descending).
//
// Example usage:
//
//	order := es.AggOrder("_count", es.Order.Desc)
//	// Sorts the aggregation buckets by count in descending order.
//
// Parameters:
//   - key: The aggregation key to sort by (e.g., "_count", "_key", or a metric sub-aggregation name).
//   - order: The sorting order (ascending or descending), defined in the `Order` package.
//
// Returns:
//
//	An es.aggOrder object representing the order definition.
func AggOrder(key string, order Order.Order) aggOrder {
	return aggOrder{
		key: order,
	}
}
