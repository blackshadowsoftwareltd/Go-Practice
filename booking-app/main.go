package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
	"strings"
)

const conferenceName = "Go Conference"
const conferenceTickets int = 50

var remainingTickets uint = uint(conferenceTickets)
var bokings = make([]map[string]string, 0)
var firstName, lastName, email string
var userTickets uint

func main() {

	helper.GreetUsers(conferenceName, conferenceTickets, remainingTickets)

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

			var userData = make(map[string]string)
			userData["firstName"] = firstName
			userData["lastName"] = lastName
			userData["email"] = email
			userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

			bokings = append(bokings, userData)
			fmt.Printf("\nbookings map %v\n", bokings)
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
