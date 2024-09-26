package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type CLI struct {
	reader *bufio.Reader
}

func New(userInput io.Reader) *CLI {
	c := &CLI{
		reader: bufio.NewReader(userInput),
	}
	return c
}

func (c *CLI) Run() {
	c.Ask()
}

func (c *CLI) Ask() {
	for {
		fmt.Println("pokedex >")
		userInput, _, err := c.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "The CLI failed to read your input: %s\n", err)
			continue
		}
		result := splitToLowerCase(string(userInput))
		c.ProcessInput(string(result))
	}
}

func (c *CLI) ProcessInput(userInput string) {
	switch userInput {
	case "help":
		fmt.Printf("Displays a help message.\n")
	case "exit":
		os.Exit(0)
	}
}

func splitToLowerCase(input string) []rune {
	return []rune(strings.ToLower(input))
}
