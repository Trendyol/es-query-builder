package boundaryscanner_test

import (
	"testing"

	BoundaryScanner "github.com/Trendyol/es-query-builder/es/enums/boundary-scanner"

	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_BoundaryScannerString(t *testing.T) {
	tests := []struct {
		boundaryScanner BoundaryScanner.BoundaryScanner
		result          string
	}{
		{BoundaryScanner.Chars, "chars"},
		{BoundaryScanner.Sentence, "sentence"},
		{BoundaryScanner.Word, "word"},
	}

	for _, test := range tests {
		t.Run(test.result, func(t *testing.T) {
			assert.Equal(t, test.result, test.boundaryScanner.String())
		})
	}
}
