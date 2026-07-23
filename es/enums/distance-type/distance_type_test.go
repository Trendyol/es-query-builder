package distancetype_test

import (
	"testing"

	DistanceType "github.com/Trendyol/es-query-builder/es/enums/distance-type"

	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_DistanceTypeString(t *testing.T) {
	tests := []struct {
		distanceType DistanceType.DistanceType
		result       string
	}{
		{DistanceType.Arc, "arc"},
		{DistanceType.Plane, "plane"},
	}

	for _, test := range tests {
		t.Run(test.result, func(t *testing.T) {
			assert.Equal(t, test.result, test.distanceType.String())
		})
	}
}
