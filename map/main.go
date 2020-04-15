package main

import "fmt"

func main() {

	// colors := make(map[string]string)

	// colors["white"] = "#ffffff"

	// delete(colors, "white")

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bg745",
		"white": "#ffffff",
	}
	printMap(colors)
	// fmt.Println(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
