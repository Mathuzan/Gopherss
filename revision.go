package main

import(
	"fmt"
)

func main() {

	var conferenceName string = "Go Conference"
	const ConferenceTickets int = 50
	var remainingTickets int = 50

	var UserName string
	UserName = "Tom"

	fmt.Println("Welcome to", conferenceName, "our Conference booking application")
	fmt.Println("Get your tickets here to attend!")
	fmt.Printf("We have total of %v tickets and %v are still avaialble!", ConferenceTickets, remainingTickets)
	fmt.Println(UserName)
}