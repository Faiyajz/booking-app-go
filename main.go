package main

import "fmt"

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

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("The whole slice: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Slice type: %T\n", bookings)
	fmt.Printf("Slice Size: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	fmt.Printf("These are all our bookings: %v\n",bookings)
}
