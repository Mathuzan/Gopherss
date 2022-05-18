// My weight loss program 
package main

import "fmt"

//main is the function where it all begins.
func main() {
	fmt.Print("My weight on the surface of Mars is ")
	fmt.Print(149.0 * 0.3783)
	fmt.Print(" lbs, and I would be ")
	fmt.Print(41 * 365 / 687)
	fmt.Println(" Years old.")
	fmt.Println("My weight on the surface of Mars is", 149.0*0.3783, "lbs, and I would be", 41*365.2425/678, "Years old.")
	fmt.Printf("My weight on the surface of %v is %v lbs.\n", "Earth", 149.0)
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galatic", 100)

}