package main

import "fmt"

func main() {

	/*second option
	 */
	// var colors map[string]string

	/*third option
	 */
	// colors := make(map[int]string)

	// colors[10] = "#fffff"
	// delete(colors, 10)

	/*first option
	 */
	colors := map[string]string{ // all the keys are of type string and all the values are of type string
		"red":   "#FF0000",
		"green": "#008000",
		"white": "#FFFFFF",
	}

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
