package main

import (
	"fmt"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = uint(conferenceTickets)
	var bokings []string

	fmt.Printf("conferenceTickets %T remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	for {

		var firstName, lastName string
		var userTickets uint
		// ask user for their name

		fmt.Println("Enter your first name : ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name : ")
		fmt.Scan(&lastName)

		fmt.Println("Enter numger of tickets : ")
		fmt.Scan(&userTickets)
		if userTickets <= remainingTickets {

			remainingTickets = remainingTickets - userTickets
			bokings = append(bokings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email \n", firstName, lastName, userTickets)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			if remainingTickets == 0 {
				// end program
				fmt.Printf("Our conference is booked out. Come back next year!\n")
				break
			}
		} else {
			fmt.Printf("We onley have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)

		}

	}
}
