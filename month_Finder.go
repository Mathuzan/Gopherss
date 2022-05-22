//Package main begins
package main

//import packages
import(
	"fmt"
)

//Function main begins
func main() {
	fmt.Printf("Enter a value from the keyboard :  ")

	var i uint
	fmt.Scan(&i)

	switch i {
	case 1:
		fmt.Println("January") 
	case 2:
		fmt.Println("February")
	case 3:
		fmt.Println("March")
	case 4:
		fmt.Println("April")
	case 5:
		fmt.Println("May")
	case 6:
		fmt.Println("June")
	case 7:
		fmt.Println("July")
	case 8:
		fmt.Println("Auguest")
	case 9:
		fmt.Println("September")
	case 10:
		fmt.Println("October")
	case 11:
		fmt.Println("November")
	case 12:
		fmt.Println("December")
	default:
		fmt.Println("Enter a valid no!")
	}
}