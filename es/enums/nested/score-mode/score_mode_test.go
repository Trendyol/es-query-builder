package scoremode_test

import (
	"testing"

	scoremode "github.com/Trendyol/es-query-builder/es/enums/nested/score-mode"
	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_ScoreModeString(t *testing.T) {
	tests := []struct {
		mode   scoremode.ScoreMode
		result string
	}{
		{scoremode.Avg, "avg"},
		{scoremode.Max, "max"},
		{scoremode.Min, "min"},
		{scoremode.None, "none"},
		{scoremode.Sum, "sum"},
	}

	for _, test := range tests {
		t.Run(test.result, func(t *testing.T) {
			assert.Equal(t, test.result, test.mode.String())
		})
	}
}
