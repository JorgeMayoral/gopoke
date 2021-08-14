package cli

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	gopoke "github.com/JorgeMayoral/gopoke/src/internal"
	"github.com/JorgeMayoral/gopoke/src/internal/errors"
	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

const idFlag = "id"
const outputFlag = "output"
const filenameFlag = "filename"
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
	pokemonsCmd.Flags().BoolP(outputFlag, "o", false, "get output as csv file")
	pokemonsCmd.Flags().String(filenameFlag, "output", "custom output filename")
	pokemonsCmd.Flags().String(limitFlag, "20", "number of pokemon to fetch")
	pokemonsCmd.Flags().String(offsetFlag, "0", "number of pokemon to skip")

	return pokemonsCmd
}

func runPokemonsFn(repository gopoke.PokemonRepo) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString(idFlag)
		output, _ := cmd.Flags().GetBool(outputFlag)
		filename, _ := cmd.Flags().GetString(filenameFlag)
		limit, _ := cmd.Flags().GetString(limitFlag)
		offset, _ := cmd.Flags().GetString(offsetFlag)

		if id != "" {
			i, _ := strconv.Atoi(id)
			pokemon, err := repository.GetPokemonById(i)

			if errors.IsDataUnreacheable(err) {
				log.Fatal(err)
			}

			fmt.Printf("Pokemon ID: %v\nName: %v\nBaseExperience: %v\nHeight: %v\nIsDefault: %v\nOrder: %v\nWeight: %v\n", pokemon.PokemonID, pokemon.Name, pokemon.BaseExperience, pokemon.Height, pokemon.IsDefault, pokemon.Order, pokemon.Weight)

			if output {
				filename = fmt.Sprintf("%v.csv", filename)
				file, _ := os.Create(filename)
				w := csv.NewWriter(file)
				headers := pokemon.GetHeaders()
				values := pokemon.ToSlice()

				if err := w.Write(headers); err != nil {
					return
				}

				if err := w.Write(values); err != nil {
					return
				}

				w.Flush()
				file.Close()
			}

			return
		} else {
			l, _ := strconv.Atoi(limit)
			o, _ := strconv.Atoi(offset)
			pokemons, err := repository.GetPokemons(l, o)

			if errors.IsDataUnreacheable(err) {
				log.Fatal(err)
			}

			for _, p := range pokemons {
				fmt.Printf("Name: %v\nUrl: %v\n\n", p.Name, p.Url)
			}
		}
	}
}
