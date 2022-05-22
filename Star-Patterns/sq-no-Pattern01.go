//Package main begins
package main

//Import packages
import(
	"fmt"
)

//Function main begins
func main() {

	fmt.Println("Square no Pattern")

	fmt.Println("Enter the no of rows: ")
	
	var i int
	fmt.Scan(&i)

	fmt.Println("Enter the no of Columns: ")

	var j int
	fmt.Scan(&j)

	for x := 1; x <= 5; x++ {		
		for y := 1; y <= j; y++ {
			fmt.Print("1")
		}
		fmt.Println()
	}
}