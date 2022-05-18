// Working in Package main.
package main

// Import fmt & String package
import (
	"fmt"
	"strings"
)

func main() {

	//Blinding midday Sun
	fmt.Println("Blinding mdiday Sun")

	var command = "midday Sun"
	var wearShades = strings.Contains(command, "midday Sun")

	fmt.Println("Wear the shades: ", wearShades)

	fmt.Println(strings.Contains("cave entrance sign", "read"))

	//Comparision Operators
	fmt.Println("There is a sign near the entrance that reads 'No Minors'.")

	var age = 40
	var minor = age < 20

	fmt.Printf("At the age %v, am I a minor? %v\n", age, minor)

	var call = 2
	var attend = call >= 2

	fmt.Printf("Show notification to the %vnd because it is %v\n", call, attend)

	fmt.Println("apple" > "Banana")
}
