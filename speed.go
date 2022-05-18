//How long it take to get to Mars
package main
 
//Import fmt package
import "fmt"

func main() {
	const lightSpeed = 299792 //km/s
	var distance = 560000000  // km 

	var (
		weight = 1000 //Weight of Man
		height = 20   // Height of man
	)
	
	fmt.Println(distance/lightSpeed, "seconds")

	distance = 401000000
	fmt.Println(distance/lightSpeed, "seconds")

	fmt.Printf("BMI of the man is %v", weight*height)
}