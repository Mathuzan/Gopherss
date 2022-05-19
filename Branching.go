// Branching with if
package main

// Import the packages
import (
	"fmt"
	"strings"
)

//Main function begins
func main () {
	var command = "rooms"

	if command == "cave" {
		fmt.Printf("You've entered into the %v", "Cave")
	} else if command == "Entrance" {
		fmt.Printf("You've entered into the %v", "Entrance")
	} else if command == "Mountain" {
		fmt.Printf("You've entered into the %v", "Mountain")
	} else {
		fmt.Println("Everything is white!")
	}

	var dataA = "JavaScript"
	var dataB = "Golang"

	var source = "This is a Golang"

	if strings.Contains(source, dataA) {
		fmt.Println(dataA)
	} else {
		fmt.Println(dataB)
	}

}