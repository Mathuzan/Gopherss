//Package main begins
package main

//Import the packages
import (
	"fmt"
	"math/rand"
)

//Function main begins
func main() {

	fmt.Println("Enter a week no 1-7")

	switch Weekno := rand.Intn(7); Weekno {
	case 1:
		fmt.Println("Monday")
	case 2: 
	    fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	}

	var userName string

	userName = "Tom"
	fmt.Println(userName)

}