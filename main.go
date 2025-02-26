package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//var reader io.Reader
	scanner := bufio.NewScanner(os.Stdin)
	commands := commands()
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)
		data, err := commands[words[0]]
		if !err {
			fmt.Printf("Unknown command: %v\n", words[0])
			continue
		}
		data.callback()
	}

}

func cleanInput(text string) []string {
	return strings.Fields(strings.Trim(strings.ToLower(text), " "))

}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, c := range commands() {
		fmt.Printf("%v: %v\n", c.name, c.description)
	}
	return nil
}
