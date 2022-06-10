package main

import(
	"fmt"
)

func main() {

	var conferenceName = "Go Conference"
	const ConferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to", conferenceName, "our Conference booking application")
	fmt.Println("Get your tickets here to attend!")
	fmt.Printf("We have total of %v tickets and %v are still avaialble!", ConferenceTickets, remainingTickets)

}