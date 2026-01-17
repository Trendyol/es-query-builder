package es

// genericPutInTheField safely inserts a key-value pair into a nested es.Object
// within the root map. It returns the root unchanged if parentKey doesn't exist
// or is not an es.Object.
func genericPutInTheField[T ~map[string]any](root T, parentKey, key string, value any) T {
	if container, ok := root[parentKey].(Object); ok {
		container[key] = value
	}
	return root
}

// genericPutInTheFieldOfFirstChild inserts a key-value pair into the first child
// es.Object found within a parent container. It traverses the parent's immediate children
// and modifies the first es.Object it encounters. This is useful for nested query structures
// where parameters need to be added to the first nested element (e.g., adding fields to
// the first query clause in a bool filter array).
func genericPutInTheFieldOfFirstChild[T ~map[string]any](root T, parentKey, key string, value any) T {
	if container, ok := root[parentKey].(Object); ok {
		for _, item := range container {
			if child, chOk := item.(Object); chOk {
				child[key] = value
				return root
			}
		}
	}
	return root
}

// genericPutInTheFieldOfFirstObject inserts a key-value pair into the first es.Object
// found in the root map. It iterates through the root's values and modifies the first
// es.Object it encounters. Useful for query types that don't have a specific parent key
// wrapper (e.g., query_string, simple_query_string, sort).
func genericPutInTheFieldOfFirstObject[T ~map[string]any](root T, key string, value any) T {
	for _, fieldObj := range root {
		if fieldObject, ok := fieldObj.(Object); ok {
			fieldObject[key] = value
			break
		}
	}
	return root
}
