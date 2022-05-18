// Scoping rules.
package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main() {
	var count = 0
	year := 2018

	for count < 10 {
		var num = rand.Intn(10) + 1
		fmt.Println(num)
		count++
	}

	fmt.Println("Variable 'num' is now Out of Scope!")

	for count := 10; count < 0; count-- {
		fmt.Println(count)
	}

	switch num := rand.Intn(10); num {
	case 0:
		fmt.Println("Space Adventures!")
	case 1:
		fmt.Println("SpaceX")
	case 2:
		fmt.Println("Virgin Galactic")
	default: 
	fmt.Println("Random Spaceline #", num)
	}

	switch month := rand.Intn(12) + 1; month {
	case 2:
		day := rand.Intn(28) + 1
		fmt.Println(era, year, month, day)
	case 4, 6, 9, 11:
		day := rand.Intn(30) + 1
		fmt.Println(era, year, month, day)
	default:
		day := rand.Intn(31) + 1
		fmt.Println(era, year, month, day)
	}
}