//Package main begins
package main

//Import packages
import(
	"fmt"
)

//Function main begins
func main() {

	//Initializing variable
	var(
		x int
		y int
	)

	//Input rows
	fmt.Println("Enter no of rows: ")
	
	var i int
	fmt.Scan(&i)

	//Input Columns
	fmt.Println("Enter no of Columns: ")

	var j int
	fmt.Scan(&j)

	for x = 1; x <= i; x++ {
		for y = 1; y <= j; y++ {
			if (x % 2 == 0) { 
				fmt.Printf("0")	
			} else {
				fmt.Print("1")
			}
		}

		fmt.Println()

	}
}