package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0) //empty list of map

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

//how can we save mixed data types of an entities
//we can use Struct which stands for "Structure" that can hold mixed data types

// type UserData struct {
// firstName string
// lastName string
// email string
// numberOfTickets uint
// }

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(firstName, lastName, userTickets, email)

			sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}
		} else if userTickets == remainingTickets {

			fmt.Printf("You have bought all the tickets. Thank you\n")
			fmt.Printf("All the tickets sold out. Please come back later\n")
			break

		} else {

			if !isValidName {

				fmt.Println("First name or Last name you entered is invalid")
			}

			if !isValidEmail {

				fmt.Println("Email address you entered is invalid")
			}

			if !isValidTicketNumber {

				fmt.Println("Number of tickets you entered is invalid")
			}

		}

	}

}

func greetUsers() {

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {

	firstNames := []string{}

	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames

}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets you want to book: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, userTickets uint, email string) {

	remainingTickets = remainingTickets - userTickets

	//Create a struct instead of map
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {

	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("################################")
	fmt.Printf("Sending %v tickets to email address %v\n", ticket, email)
	fmt.Println("################################")

}
