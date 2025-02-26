package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
	cleanInput("ojo IIII VAFGA   ")

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
