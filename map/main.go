package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":    "#ff0000",
		"green":  "#4b745",
		"white":  "#ffffff",
		"purple": "#6a0dad",
	}
	colors["blue"] = "#0000ff"
	delete(colors, "green")
	printMap(colors)

	var ages map[string]int
	fmt.Println(ages)

	var attendance = make(map[string]int)
	attendance["Elton John In Ottawa"] = 1_000_000
	fmt.Println(attendance)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println(key, value)
	}
}
