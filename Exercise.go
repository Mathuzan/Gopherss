// Go Practice - https://golangr.com/exercises/
package main

// Import packages for the function
import (
	"fmt"
)

// Begin main function
func main() {

	inputNo1 := 10
	inputNo2 := 20

	if (inputNo1 > inputNo2) {
		fmt.Println("Maximum value is : ", inputNo1)
	} else {
		fmt.Println("Maximum value is : ", inputNo2)
	}

	fmt.Println("Exercise 1 is done!")
	fmt.Println("Code changes aren't working!")

	inputNo03 := 20
	inputNo04 := 30
	inputNo05 := 40

	if (inputNo03 > inputNo04) {

		if (inputNo03 > inputNo05) {
			fmt.Println("inputNo03 is the greatest integer")
		} else {
			fmt.Printf("%v is the largest no!", inputNo05)
		} 
	} else {
		if (inputNo03 > inputNo05) {
			fmt.Printf("%v is the greatest no", inputNo04)
		} else {
			fmt.Printf("%v is the greatest no and its %v", inputNo05, 40)
		}
	}
} 