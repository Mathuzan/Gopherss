//Package main begins
package main

// Import packages
import (
	"fmt"
	"math/rand"
)

//Main function begins
func main() {

	third := 1.0 / 3
	fmt.Println(third)
	fmt.Printf("%v\n", third)
	fmt.Printf("%f\n", third)
	fmt.Printf("%3f\n", third)
	fmt.Printf("%4.2f\n", third)

	piggyBank := 0

	for piggyBank < 2000 {

		switch rand.Intn(3) {
		case 0:
			piggyBank += 5
		case 1:
			piggyBank += 10
		case 2:
			piggyBank += 25
		}
	
		dollars := piggyBank / 100
		cents := piggyBank % 100
		fmt.Printf("$%d.%02d\n", dollars, cents)

	}
}