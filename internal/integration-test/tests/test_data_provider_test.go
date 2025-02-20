package tests_test

import (
	"github.com/stretchr/testify/assert"
	testdataprovider "integration-tests/test-data-provider"
	"testing"
	"unsafe"
)

func Test_TestDataProvider_Should_Provide_Pokemon_Data(t *testing.T) {
	// Given When
	pokedata, err := testdataprovider.PokemonTestData()
	if err != nil {
		assert.Fail(t, err.Error())
	}
	pokedata2, err := testdataprovider.PokemonTestData()
	if err != nil {
		assert.Fail(t, err.Error())
	}

	// Then
	assert.NotNil(t, pokedata)
	assert.NotEmpty(t, pokedata)
	assert.NotNil(t, pokedata2)
	assert.NotEmpty(t, pokedata2)
	assert.EqualValues(t, pokedata, pokedata2)
	assert.False(t, &pokedata == &pokedata2)
	assert.False(t, unsafe.Pointer(&pokedata) == unsafe.Pointer(&pokedata2))
}
