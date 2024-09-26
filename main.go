package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/maBeghtel/pokedexcli/cli"
)

func main() {
	fmt.Printf("Welcome to the Pokedex!\n\n")
	fmt.Fprintf(os.Stdout, "Usage: \n\nhelp: Displays a help message. \nexit: Exits the Pokedex\n")
	p := cli.New(bufio.NewReader(os.Stdin))
	p.Run()
}
