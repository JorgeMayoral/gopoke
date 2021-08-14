package gopoke

// Pokemon representation of pokemon into data struct
type Pokemon struct {
	PokemonID      int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
}

type PokemonSimple struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonRepo interface {
	GetPokemons(limit, offset int) ([]PokemonSimple, error)
	GetPokemonById(pokemonID int) (Pokemon, error)
}

// NewPokemon initialize struct pokemon
func NewPokemon(pokemonId int, name string, baseExperience, height int, isDefault bool, order, weight int) (p Pokemon) {
	p = Pokemon{
		PokemonID:      pokemonId,
		Name:           name,
		BaseExperience: baseExperience,
		Height:         height,
		IsDefault:      isDefault,
		Order:          order,
		Weight:         weight,
	}
	return
}

func NewPokemonSimple(name, url string) (p PokemonSimple) {
	p = PokemonSimple{
		Name: name,
		Url:  url,
	}

	return
}
