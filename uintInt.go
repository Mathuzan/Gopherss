//Package main begins
package main

//Import packages
import(
	"fmt"
)

func main() {

	// year := 2018
	// var year = 2018
	var year int64 = 2018 

	fmt.Println(year)
	fmt.Printf("%T\n", year)

	var name string
	fmt.Scan(&name)

	var green uint = 3
	fmt.Printf("%08b\n", green)

	green++
	fmt.Printf("%08b\n", green)
	

}