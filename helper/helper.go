package helper

import "strings"

// function with multiples returns
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) { // capitalize the first letter from a function will "export" (public) and make this acessible in others packages
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	// returning multiples values
	return isValidName, isValidEmail, isValidTicketNumber
}
