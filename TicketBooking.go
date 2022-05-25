//Package main begins
package main

//Import the packages
import (
	"fmt"
)

//Function main begins
func main() {
	conferenceName := "Go Conference"
	const ConferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", ConferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")

	var booking [50]string
	booking[0] = "Nana"
	booking[1] = "Nicole"

	//Print the values of array
	fmt.Printf("The whole array: %v\n", booking)
	fmt.Printf("The first value: %v\n", booking[0])
	fmt.Printf("Array type: %T\n", booking)
	fmt.Printf("Array length: %v\n", len(booking))


	var firstName string
	var lastName string
	var email string
	var userTickets uint
 
	//ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter your Ticket count: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, ConferenceTickets)


}