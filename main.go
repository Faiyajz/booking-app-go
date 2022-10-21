package main

import (
	"fmt"
	"strings"
)

func main() {
	// var conferenceName string = "Go Conference"

	conferenceName := "Go Conference" //we can't declare const variable with this alternative syntax

	const conferenceTickets = 50
	// const conferenceTickets int = 50

	var remainingTickets uint = 50
	var bookings []string //Slice Declaration
	// var bookings = []string{} //alternative way of declaring slice
	// bookings := []string{} // Another alternative way of declaring slice

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for remainingTickets > 0 && len(bookings) < 50 {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// fmt.Println(remainingTickets)
		// fmt.Println(&remainingTickets) pointer is referring the memory address using a hash(Pointer)

		fmt.Println("Please enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Please enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Please enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets you want to book: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets < remainingTickets

		// isValidCity := city == "Singapore" || city == "London"
		// isInValidCity := city != "Singapore" || city != "London"
		// !isValidCity := city == "Singapore" || city == "London"

		if isValidName && isValidEmail && isValidTicketNumber {

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("The whole slice: %v\n", bookings)
			fmt.Printf("The first value: %v\n", bookings[0])
			fmt.Printf("Slice type: %T\n", bookings)
			fmt.Printf("Slice Size: %v\n", len(bookings))

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{} //another way of declaring-- var firstNames []string
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			// books := []string{"Maria Zaman","Hermione Ginger","Afrose Meghla"}

			// bookFirstNames := []string{}
			// for _, book := range(books){
			// 	var bookNames = strings.Fields(book)
			// 	bookFirstNames = append(bookFirstNames, bookNames[0])
			// }

			// fmt.Printf("Book First Names are: %v\n", bookFirstNames)

			// var noTicketsRemaining bool = remainingTickets == 0
			// // noTicketsRemaining := remainingTickets == 0 //another way of the declaration
			// if noTicketsRemaining {
			// 	fmt.Println("Our conference is booked out. Come back next year")
			// 	break
			// }

			//Since we are using this variable "noTicketsRemaining" only once
			//there's actually no need to save this expression in a separate variable

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

	// Switch Case Statement

	city := "London"

	switch city {

	case "New York":
		//Execute code for booking New York conference tickets

	case "Singapore", "Hong Kong":
		//Execute code for booking Singapore and Hong Kong conference tickets

	case "London", "Berlin":
		//Execute code for booking London and Berlin conference tickets

	case "Mexico City":
		//Execute code for booking Mexico City conference tickets

	default:
		//execute when none of the above is true
		fmt.Println("No valid city is selected")
	}
}
