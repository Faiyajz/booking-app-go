package helper

import "strings"

var MyVar = "somevalue" // Global variable has to be declared starting with capital letter

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets < remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
