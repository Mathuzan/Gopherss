package main

import "fmt"

func main () {
	fmt.Print("Enter your name!\n")
	fmt.Println("Enter your fucking data")

	//Getting user inputs from users.
	var userName string
	var userTickets int

	//Ask user for their name
	fmt.Scan(&userName)
	fmt.Scan(&userTickets)

	fmt.Printf("User %v Booked %v tickets\n", userName, userTickets )

	fmt.Println(&userName)
	fmt.Println(&userTickets)
}