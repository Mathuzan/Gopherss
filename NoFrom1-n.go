//Package main begins
package main

//Import main function
import(
	"fmt"
)

//Function main begins
func main() {

	fmt.Printf("Enter a value from the keyboard: ")

	var i int
	fmt.Scan(&i)

	for j := 0; j <= i; j++ {
		fmt.Printf("%v\t", j)	
	}

	//Print all the even no from 1-10
	fmt.Println("\nPrint all the even no from 1-10")

	for e := 0; e <= 10; e += 2 {
		fmt.Printf("%v\t", e)
	}

	for e := 0; e <= 10; e++ {
		if (e % 2 == 0) {
			fmt.Printf("%v\t", e)
		} else {
			fmt.Printf("%v\n", e)
		}
	}

}