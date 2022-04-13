package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
	"sync"
	"time"
)

const conferenceName = "Go Conference"
const conferenceTickets int = 50

var remainingTickets uint = uint(conferenceTickets)
var bokings = make([]UserData, 0)
var firstName, lastName, email string
var userTickets uint

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var waitGroup = sync.WaitGroup{}

func main() {

	helper.GreetUsers(conferenceName, conferenceTickets, remainingTickets)

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

		waitGroup.Add(2)
		go sendTicket(firstName, lastName, userTickets)	 
		go taskDone()

		var userData = UserData{
			firstName:       firstName,
			lastName:        lastName,
			email:           email,
			numberOfTickets: userTickets,
		}

		bokings = append(bokings, userData)
		fmt.Printf("\nbookings map %v\n", bokings)
		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation at %v \n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		if remainingTickets == 0 {
			// end program
			fmt.Printf("Our conference is booked out. Come back next year!\n")
			// break
		}
	} else {
		fmt.Printf("We onley have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)

	}

	waitGroup.Wait()

}

func sendTicket(
	firstName string, lastName string, userTickets uint) {
	time.Sleep(time.Second * 5)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println("--------------------------------------------------------------------------------")
	waitGroup.Done()
}
func taskDone() {
	time.Sleep(time.Second * 7)
	fmt.Println(" Task done --------------------------------------------------------------------------------")
	waitGroup.Done()
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
