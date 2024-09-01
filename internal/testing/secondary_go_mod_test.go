package testing_test

import (
	"testing"

	"github.com/GokselKUCUKSAHIN/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func Test_2nd_go_mod(t *testing.T) {
	query := es.NewQuery(nil)
	assert.NotNil(t, query)
}
