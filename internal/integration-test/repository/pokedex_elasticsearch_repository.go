package repository

import (
	"encoding/json"

	"integration-tests/constants"
	"integration-tests/model"
	"integration-tests/model_repository"

	"github.com/elastic/go-elasticsearch/v8"
)

type pokedexElasticsearchRepository struct {
	BaseGenericRepository[string, model.Pokemon]
}

func NewPokedexElasticsearchRepository(client *elasticsearch.Client) BaseGenericRepository[string, model.Pokemon] {
	return &pokedexElasticsearchRepository{
		NewBaseGenericRepository(client, constants.PokemonIndex, mapToPokemon, model_repository.MapToId, mapToPokemonId),
	}
}

func mapToPokemonId(pokemon model.Pokemon) string {
	return pokemon.GetDocumentID()
}

func mapToPokemon(docID string, searchHit model_repository.SearchHit) (string, model.Pokemon, error) {
	var pokemon model.Pokemon
	if err := json.Unmarshal(searchHit.Source, &pokemon); err != nil {
		return "", model.Pokemon{}, err
	}
	return docID, pokemon, nil
}
