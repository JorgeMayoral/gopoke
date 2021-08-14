package main

import (
	_ "flag"

	"github.com/JorgeMayoral/gopoke/src/internal/storage/pokeapi"

	gopoke "github.com/JorgeMayoral/gopoke/src/internal"
	"github.com/JorgeMayoral/gopoke/src/internal/cli"
	"github.com/spf13/cobra"
)

func main() {
	var repo gopoke.PokemonRepo = pokeapi.NewPokeapiRepository()

	rootCmd := &cobra.Command{Use: "gopoke"}
	rootCmd.AddCommand(cli.InitPokemonsCmd(repo))
	rootCmd.Execute()
}
