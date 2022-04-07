package main

import (
	"fmt"
	"strings"
)

const conferenceName = "Go Conference"
const conferenceTickets int = 50

var remainingTickets uint = uint(conferenceTickets)
var bokings []string
var firstName, lastName, email string
var userTickets uint

func main() {

	greetUsers()

	for {

		// ask user for their name

		email = pickUserEmail(email)
		firstName, lastName = pickFirstAndLastName()
		userTickets = pickUserTickets()

		isValidName := len(firstName) >= 3 && len(lastName) >= 3
		isValidEmail := strings.Contains(email, "@")
		isValidTicket :=
			userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicket {

			remainingTickets = remainingTickets - userTickets
			bokings = append(bokings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation at %v \n", firstName, lastName, userTickets, email)
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
func pickUserTickets() uint {
	fmt.Println("Enter numger of tickets : ")
	fmt.Scan(&userTickets)
	return userTickets
}

func pickFirstAndLastName() (string, string) {
	fmt.Println("Enter your first name : ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name : ")
	fmt.Scan(&lastName)
	return firstName, lastName
}
func pickUserEmail(email string) string {
	fmt.Println("Enter your Email : ")
	fmt.Scan(&email)
	return email

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
