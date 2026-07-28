package suggestmode_test

import (
	"testing"

	SuggestMode "github.com/Trendyol/es-query-builder/es/enums/suggest-mode"

	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_SuggestModeString(t *testing.T) {
	tests := []struct {
		suggestMode SuggestMode.SuggestMode
		result      string
	}{
		{SuggestMode.Missing, "missing"},
		{SuggestMode.Popular, "popular"},
		{SuggestMode.Always, "always"},
	}

	for _, test := range tests {
		t.Run(test.result, func(t *testing.T) {
			assert.Equal(t, test.result, test.suggestMode.String())
		})
	}
}
