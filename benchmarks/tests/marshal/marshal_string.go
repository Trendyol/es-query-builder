package marshal

import (
	"encoding/json"
	"testing"

	"github.com/Trendyol/es-query-builder/test/assert"
)

func String(t *testing.T, query map[string]any) string {
	marshal, err := json.Marshal(query)
	assert.Nil(t, err)
	return string(marshal)
}
