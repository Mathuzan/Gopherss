//Package main begins
package main

//Import packages
import(
	"fmt"
)

//Function main begins
func main() {

	//Enter the no of rows
	fmt.Println("Enter the no of rows: ")
	
	var i int
	fmt.Scan(&i)

	//Enter the no of colums
	fmt.Println("Enter the no of Colums: ")

	var j int
	fmt.Scan(&j)

	for x := 1; x <= i; x++ {
		for y := 1; y <= j; y++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}