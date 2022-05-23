//Package main
package main

//Import packages
import "fmt"

//Function main begins
func main() {

	//Input the rows
	fmt.Println("Enter the no of rows: ")

	var i int
	fmt.Scan(&i)


	for x := 1; x <= i; x++ {
		for y := 1; y <= i; y++ {	
			if (x == 1 || y == i || y == 1 || x == i){
				fmt.Print("*")
			}	else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}