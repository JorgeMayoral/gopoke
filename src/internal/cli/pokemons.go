package cli

import (
	"fmt"
	"strconv"

	gopoke "github.com/JorgeMayoral/gopoke/src/internal"
	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

const idFlag = "id"
const limitFlag = "limit"
const offsetFlag = "offset"

// InitPokemonsCmd initialize pokemons command
func InitPokemonsCmd(repository gopoke.PokemonRepo) *cobra.Command {
	pokemonsCmd := &cobra.Command{
		Use:   "pokemons",
		Short: "Print data about pokemons",
		Run:   runPokemonsFn(repository),
	}

	pokemonsCmd.Flags().StringP(idFlag, "i", "", "id of the pokemon")
	pokemonsCmd.Flags().String(limitFlag, "20", "number of pokemon to fetch")
	pokemonsCmd.Flags().String(offsetFlag, "0", "number of pokemon to skip")

	return pokemonsCmd
}

func runPokemonsFn(repository gopoke.PokemonRepo) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString(idFlag)
		limit, _ := cmd.Flags().GetString(limitFlag)
		offset, _ := cmd.Flags().GetString(offsetFlag)

		if id != "" {
			i, _ := strconv.Atoi(id)
			pokemon, _ := repository.GetPokemonById(i)
			fmt.Printf("Pokemon ID: %v\nName: %v\nBaseExperience: %v\nHeight: %v\nIsDefault: %v\nOrder: %v\nWeight: %v\n", pokemon.PokemonID, pokemon.Name, pokemon.BaseExperience, pokemon.Height, pokemon.IsDefault, pokemon.Order, pokemon.Weight)
			return
		} else {
			l, _ := strconv.Atoi(limit)
			o, _ := strconv.Atoi(offset)
			pokemons, _ := repository.GetPokemons(l, o)
			for _, p := range pokemons {
				fmt.Printf("Name: %v\nUrl: %v\n\n", p.Name, p.Url)
			}
		}
	}
}
