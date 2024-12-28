package range_relation_test

import (
	"testing"

	RangeRelation "github.com/Trendyol/es-query-builder/es/enums/range-relation"

	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_RangeRelationString(t *testing.T) {
	tests := []struct {
		relation RangeRelation.RangeRelation
		result   string
	}{
		{RangeRelation.Within, "within"},
		{RangeRelation.Contains, "contains"},
		{RangeRelation.Intersects, "intersects"},
	}

	for _, test := range tests {
		t.Run(test.result, func(t *testing.T) {
			assert.Equal(t, test.result, test.relation.String())
		})
	}
}
