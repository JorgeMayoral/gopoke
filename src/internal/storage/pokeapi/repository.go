package pokeapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	gopoke "github.com/JorgeMayoral/gopoke/src/internal"
	"github.com/JorgeMayoral/gopoke/src/internal/errors"
)

const (
	pokemonsEndpoint = "/pokemon"
	pokeapiURL       = "https://pokeapi.co/api/v2"
)

type pokemonRepo struct {
	url string
}

type apiResponse struct {
	Count    int                    `json:"count"`
	Next     string                 `json:"next"`
	Previous string                 `json:"previous"`
	Results  []gopoke.PokemonSimple `json:"results"`
}

// NewPokeapiRepository fetch pokemons data
func NewPokeapiRepository() gopoke.PokemonRepo {
	return &pokemonRepo{url: pokeapiURL}
}

func (p *pokemonRepo) GetPokemons(limit, offset int) (pokemons []gopoke.PokemonSimple, err error) {
	response, err := http.Get(fmt.Sprintf("%v%v?limit=%v&offset=%v", p.url, pokemonsEndpoint, limit, offset))
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "error getting response from %s", pokemonsEndpoint)
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "error reading the response from %s", pokemonsEndpoint)
	}

	var apiResponse apiResponse
	err = json.Unmarshal(contents, &apiResponse)
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "can't parse response into pokemons")
	}

	pokemons = apiResponse.Results

	return
}

func (p *pokemonRepo) GetPokemonById(pokemonID int) (pokemon gopoke.Pokemon, err error) {
	response, err := http.Get(fmt.Sprintf("%v%v/%v", p.url, pokemonsEndpoint, pokemonID))
	if err != nil {
		return pokemon, errors.WrapDataUnreacheable(err, "error getting response from %s", pokemonsEndpoint)
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return pokemon, errors.WrapDataUnreacheable(err, "error reading response from %s", pokemonsEndpoint)
	}

	err = json.Unmarshal(contents, &pokemon)
	if err != nil {
		return pokemon, errors.WrapDataUnreacheable(err, "can't parse response into pokemon")
	}

	return
}
