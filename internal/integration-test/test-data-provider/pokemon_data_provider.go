package testdataprovider

import (
	"integration-tests/jsonx"
	"integration-tests/model"
)

const pokemonTestDataRelativePath = "./data/poke.fa"

var cachedTestData model.Pokemons = nil

func PokemonTestData() (model.Pokemons, error) {
	if cachedTestData != nil {
		return cachedTestData.Copy(), nil
	}
	pokemonTestData, err := jsonx.CastSlice[model.Pokemon](DecompressGz(ResolvePath(pokemonTestDataRelativePath)))
	if err != nil {
		return nil, err
	}
	cachedTestData = pokemonTestData
	return cachedTestData.Copy(), nil
}
