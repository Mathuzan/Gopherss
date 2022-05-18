// The switch Statement: Concise-switch.go
package main

//Import the packages
import (
	"fmt"
)

func main() {
	fmt.Println("There is a cavern entrance here and a path to the east!")
	var command = "go inside"

	switch command {
	case "go east" :
		fmt.Println("You head further up the mountain.")
	case "enter cave", "go inside":
		fmt.Println("You find your self in a dimly cavern")
	case "read sign":
		fmt.Println("The sign reads 'No Minors'.")
		fallthrough
	default:
		fmt.Println("Didn't quite get that.")
		fmt.Println("And follow the order on!")
	}

}