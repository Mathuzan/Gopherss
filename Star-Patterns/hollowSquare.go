//Package main begins
package main

//Import packages
import(
	"fmt"
)

//Function main begins
func main() {

	//Enter the no of rows & Colums
	fmt.Println("Enter the no of Rows : ")

	var i int
	fmt.Scan(&i)

	for x := 1; x <= i; x++ {
		for y := 1; y <= i; y++ {
			if (x == 1 || x == i || y == 1 || y == i || x == y || (y== (i - x + 1))) {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}