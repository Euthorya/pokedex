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

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)
		fmt.Printf("Your command was: %v\n", words[0])
	}

}

func cleanInput(text string) []string {
	//fmt.Println(strings.Split(strings.Trim(strings.ToLower(text), " "), " "))
	// var result []string
	// for s := range strings.SplitSeq(strings.TrimSpace(strings.ToLower(text)), " ") {
	// 	result = append(result, strings.TrimSpace(s))
	// }
	// return result
	return strings.Fields(strings.Trim(strings.ToLower(text), " "))

}
