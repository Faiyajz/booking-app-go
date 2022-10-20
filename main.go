package main

import "fmt"

func main() {
	// var conferenceName string = "Go Conference"

	conferenceName := "Go Conference" //we can't declare const variable with this alternative syntax

	const conferenceTickets = 50
	// const conferenceTickets int = 50

	var remainingTickets = 50

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int

	userName = "Faiyaj"
	userTickets = 2

	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)

}
