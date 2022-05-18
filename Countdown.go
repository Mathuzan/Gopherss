// A countdown loop
package main

import (
	"fmt"
	"time"
	"math/rand"
)

// Begin main Function
func main() {
	var count = 10

	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		count--
	}

	fmt.Println("Liftoff!")

	var degrees = 0

	for {
	fmt.Println(degrees)

	degrees++
		if degrees >= 360 {
		degrees = 0
			if rand.Intn(2) == 0 {
			break
			}
		}
	}
}