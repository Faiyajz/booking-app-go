package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = make([]map[string]string, 0) //empty list of map

// Instead of this kind of value "Faiyaj Zaman"

// we want this kind of data block for each user
// firstName: "Faiyaj"
// lastName: "Zaman"
// email: "faiyajz007@gmail.com"
// tickets: 3

//we havea to define a map for the above situation
// map is a collection of key-value pairs
//create a map for a user
// var userData = map[string]string

// first of all, create a variable of an empty map with a keyword map
// then the data type of the key and data type of the value respectively

//To create an empty map we have a built in function called "make"
// var userData = make(map[string]string)
// All keys have the same data type
// All values have the same data type
//if the values have different data type we will have to convert it
// if the values have string and int mixed data types
// we will have to convert the int to string
// strconv.FormatUint(uint64(userTickets))
//userData["firstName"] = firstName // also be declared as like this userData["a"] = firstName, key can be any name

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(firstName, lastName, userTickets, email)

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

		firstNames = append(firstNames, booking["firstName"])
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

	var userData = make(map[string]string)
	userData["firstName"] = firstName // also be declared as like this userData["a"] = firstName, key can be any name
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
