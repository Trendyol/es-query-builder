package collectmode_test

import (
	"testing"

	CollectMode "github.com/Trendyol/es-query-builder/es/enums/collect-mode"

	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_CollectModeString(t *testing.T) {
	tests := []struct {
		collectMode CollectMode.CollectMode
		result      string
	}{
		{CollectMode.BreadthFirst, "breadth_first"},
		{CollectMode.DepthFirst, "depth_first"},
	}

	for _, test := range tests {
		t.Run(test.result, func(t *testing.T) {
			assert.Equal(t, test.result, test.collectMode.String())
		})
	}
}
