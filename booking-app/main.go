package main

import (
	"fmt"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = uint(conferenceTickets)

	fmt.Printf("conferenceTickets %T remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firlstName, lastName, email string
	var userTickets uint
	// ask user for their name

	fmt.Println("Enter your email : ")
	fmt.Scan(&email)

	fmt.Println("Enter your first name : ")
	fmt.Scan(&firlstName)

	fmt.Println("Enter your last name : ")
	fmt.Scan(&lastName)

	fmt.Println("Enter numger of tickets : ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firlstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v",remainingTickets,conferenceName)
}
