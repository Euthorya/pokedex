package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	//var reader io.Reader
	scanner := bufio.NewScanner(os.Stdin)
	commands := commands()
	config := Config{}
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
		data.callback(&config)
	}

}

func cleanInput(text string) []string {
	return strings.Fields(strings.Trim(strings.ToLower(text), " "))

}

func commandExit(config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *Config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, c := range commands() {
		fmt.Printf("%v: %v\n", c.name, c.description)
	}
	return nil
}

func commandMap(config *Config) error {
	var location LocationResponse
	url := config.Next
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("There was an error while trying to fetch a map")
		return err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	decoder.Decode(&location)
	config.Next = location.Next
	config.Previous = location.Previous
	for _, loc := range location.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(config *Config) error {
	if config.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	var location LocationResponse
	url := config.Previous
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("There was an error while trying to fetch a map")
		return err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	decoder.Decode(&location)
	config.Next = location.Next
	config.Previous = location.Previous
	for _, loc := range location.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
