package main

import (
	//"os"
	//"runtime/pprof"

	gopoke "github.com/JorgeMayoral/gopoke/src/internal"
	"github.com/JorgeMayoral/gopoke/src/internal/cli"
	"github.com/JorgeMayoral/gopoke/src/internal/fetching"
	pokeapi "github.com/JorgeMayoral/gopoke/src/internal/storage/pokeapi"

	"github.com/spf13/cobra"
)

func main() {
	// CPU profiling code starts here
	// f, _ := os.Create("pokemons.cpu.prof")
	// defer f.Close()
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()
	// CPU profiling code ends here

	var repo gopoke.PokemonRepo = pokeapi.NewPokeapiRepository()

	fetchingService := fetching.NewService(repo)

	rootCmd := &cobra.Command{Use: "gopoke"}
	rootCmd.AddCommand(cli.InitPokemonsCmd(fetchingService))
	rootCmd.Execute()

	// Memory profiling code starts here
	// f2, _ := os.Create("pokemons.mem.prof")
	// defer f2.Close()
	// pprof.WriteHeapProfile(f2)
	// Memory profiling code ends here
}
