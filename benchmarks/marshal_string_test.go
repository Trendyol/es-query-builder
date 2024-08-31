package benchmarks_test

import (
	"encoding/json"
	"testing"

	"github.com/GokselKUCUKSAHIN/es-query-builder/test/assert"
)

func marshalString(t *testing.T, query map[string]any) string {
	marshal, err := json.Marshal(query)
	assert.Nil(t, err)
	return string(marshal)
}
