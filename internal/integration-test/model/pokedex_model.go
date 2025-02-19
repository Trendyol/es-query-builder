package model

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
