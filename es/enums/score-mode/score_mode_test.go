package scoremode_test

import (
	"testing"

	ScoreMode "github.com/Trendyol/es-query-builder/es/enums/score-mode"

	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_ScoreModeString(t *testing.T) {
	tests := []struct {
		mode   ScoreMode.ScoreMode
		result string
	}{
		{ScoreMode.Avg, "avg"},
		{ScoreMode.Max, "max"},
		{ScoreMode.Min, "min"},
		{ScoreMode.None, "none"},
		{ScoreMode.Sum, "sum"},
	}

	for _, test := range tests {
		t.Run(test.result, func(t *testing.T) {
			assert.Equal(t, test.result, test.mode.String())
		})
	}
}
