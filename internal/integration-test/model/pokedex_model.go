package model

import (
	"fmt"
	"github.com/Trendyol/es-query-builder/es"
)

type Pokemons []Pokemon

func (pokes Pokemons) Copy() Pokemons {
	copiedPokemons := make(Pokemons, len(pokes), len(pokes))
	for i := 0; i < len(pokes); i++ {
		copiedPokemons[i] = pokes[i].Copy()
	}
	return copiedPokemons
}

type Pokemon struct {
	Name           string        `json:"name"`
	Abilities      []Ability     `json:"abilities"`
	Moves          []Move        `json:"moves"`
	Types          []PokemonType `json:"types"`
	Stats          []Stat        `json:"stats"`
	Id             uint16        `json:"id"`
	Height         uint16        `json:"height"`
	Weight         uint16        `json:"weight"`
	BaseExperience uint16        `json:"baseExperience"`
	Order          uint16        `json:"order"`
	IsDefault      bool          `json:"isDefault"`
}

func (poke *Pokemon) GetDocumentID() string {
	return fmt.Sprintf("%d_%d", poke.Id, poke.Order)
}

func (poke *Pokemon) Copy() Pokemon {
	abilities := make([]Ability, len(poke.Abilities), len(poke.Abilities))
	for i := 0; i < len(poke.Abilities); i++ {
		abilities[i] = Ability{
			Name:     poke.Abilities[i].Name,
			Slot:     poke.Abilities[i].Slot,
			IsHidden: poke.Abilities[i].IsHidden,
		}
	}
	types := make([]PokemonType, len(poke.Types), len(poke.Types))
	for i := 0; i < len(poke.Types); i++ {
		types[i] = PokemonType{
			Name: poke.Types[i].Name,
			Slot: poke.Types[i].Slot,
		}
	}
	stats := make([]Stat, len(poke.Stats), len(poke.Stats))
	for i := 0; i < len(poke.Stats); i++ {
		stats[i] = Stat{
			Name:     poke.Stats[i].Name,
			BaseStat: poke.Stats[i].BaseStat,
			Effort:   poke.Stats[i].Effort,
		}
	}
	moves := make([]Move, len(poke.Moves), len(poke.Moves))
	for i := 0; i < len(poke.Moves); i++ {
		versionGroupDetails := make([]VersionGroupDetail, len(poke.Moves[i].VersionGroupDetails), len(poke.Moves[i].VersionGroupDetails))
		for j := 0; j < len(poke.Moves[i].VersionGroupDetails); j++ {
			versionGroupDetails[j] = VersionGroupDetail{
				MoveLearnMethodName: poke.Moves[i].VersionGroupDetails[j].MoveLearnMethodName,
				VersionGroupName:    poke.Moves[i].VersionGroupDetails[j].VersionGroupName,
				LevelLearnedAt:      poke.Moves[i].VersionGroupDetails[j].LevelLearnedAt,
			}
		}
		moves[i] = Move{
			Name:                poke.Moves[i].Name,
			VersionGroupDetails: versionGroupDetails,
		}
	}
	return Pokemon{
		Id:             poke.Id,
		Name:           poke.Name,
		Height:         poke.Height,
		Weight:         poke.Weight,
		BaseExperience: poke.BaseExperience,
		Order:          poke.Order,
		IsDefault:      poke.IsDefault,
		Abilities:      abilities,
		Types:          types,
		Stats:          stats,
		Moves:          moves,
	}
}

func (poke *Pokemon) GetMappings() string {
	return `{
  "mappings": {
    "properties": {
      "name": {
        "type": "text",
        "analyzer": "pokemon_name_analyzer",
        "search_analyzer": "standard",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "abilities": {
        "type": "nested",
        "properties": {
          "name": {
            "type": "keyword"
          },
          "slot": {
            "type": "short"
          },
          "isHidden": {
            "type": "boolean"
          }
        }
      },
      "moves": {
        "type": "nested",
        "properties": {
          "name": {
            "type": "keyword"
          },
          "versionGroupDetails": {
            "type": "nested",
            "properties": {
              "moveLearnMethodName": {
                "type": "keyword"
              },
              "versionGroupName": {
                "type": "keyword"
              },
              "levelLearnedAt": {
                "type": "short"
              }
            }
          }
        }
      },
      "types": {
        "type": "nested",
        "properties": {
          "name": {
            "type": "keyword"
          },
          "slot": {
            "type": "short"
          }
        }
      },
      "stats": {
        "type": "nested",
        "properties": {
          "name": {
            "type": "keyword"
          },
          "baseStat": {
            "type": "short"
          },
          "effort": {
            "type": "short"
          }
        }
      },
      "id": {
        "type": "short"
      },
      "height": {
        "type": "short"
      },
      "weight": {
        "type": "short"
      },
      "baseExperience": {
        "type": "short"
      },
      "order": {
        "type": "short"
      },
      "isDefault": {
        "type": "boolean"
      }
    }
  }
}`
}

func (poke *Pokemon) GetSettings() es.Object {
	return es.Object{
		"settings": es.Object{
			"index": es.Object{
				"refresh_interval":   "1s",
				"number_of_shards":   1,
				"number_of_replicas": 1,
				"max_result_window":  10_000,
				"max_terms_count":    1024,
			},
			"analysis": es.Object{
				"analyzer": es.Object{
					"pokemon_name_analyzer": es.Object{
						"type":      "custom",
						"tokenizer": "pokemon_name_tokenizer",
						"filter": es.Array{
							"lowercase",
							"asciifolding",
						},
					},
				},
				"tokenizer": es.Object{
					"pokemon_name_tokenizer": es.Object{
						"type":     "edge_ngram",
						"min_gram": 2,
						"max_gram": 20,
						"token_chars": es.Array{
							"letter",
							"digit",
						},
					},
				},
			},
		},
	}
}

type Ability struct {
	Name     string `json:"name"`
	Slot     uint16 `json:"slot"`
	IsHidden bool   `json:"isHidden"`
}

type Stat struct {
	Name     string `json:"name"`
	BaseStat uint16 `json:"baseStat"`
	Effort   uint16 `json:"effort"`
}

type PokemonType struct {
	Name string `json:"name"`
	Slot uint16 `json:"slot"`
}

type Move struct {
	Name                string               `json:"name"`
	VersionGroupDetails []VersionGroupDetail `json:"versionGroupDetails"`
}

type VersionGroupDetail struct {
	MoveLearnMethodName string `json:"moveLearnMethodName"`
	VersionGroupName    string `json:"versionGroupName"`
	LevelLearnedAt      uint16 `json:"levelLearnedAt"`
}
