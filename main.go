package main

import (
	"booking-app/helper"
	"fmt"
)

// Package Level Variables
// They can be accessed inside any of the functions and all files which are in the same package

// variable declarations with datatype determined by the value assigned
const conferenceTickets = 50

var remainingTickets uint = 50
var conferenceName = "Coffee Conference"

// [EXAMPLE] conferenceName := "Coffee Conference" // a shorthand for variable declaration with a datapyte but its not allowed in the Package Level Variables

// [EXAMPLE] var bookings [50]string // array with 50 positions and no initial value
// [REPLACED] var bookings []string // slice: dynamic array (no lenght defined)
// [REPLACED] var bookings = make([]map[string]string, 0) // empty slice of a maps
var bookings = make([]UserData, 0)

type UserData struct { // struct can be compared to "class" in other languages
	firstName       string
	lastName        string
	email           string
	userTickets 	uint
}

func main() {

	greetUsers()

	for remainingTickets > 0 && len(bookings) < 50 {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The firsts names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}

		} else {

			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Your email address is invalid")
			}
			if !isValidTicketNumber {
				fmt.Println("The number of ticker you entered is invalid")
			}

		}

	}

}

func greetUsers() {
	// printing strings using placeholders
	fmt.Printf("Welcom to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v ticket and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings { // user _ to especify unused variables or those ones you don't want to use

		// range expression allow us iterate over the elements and data structures

		//[DELETED] var names = strings.Fields(booking) // splits the string with white space as separator and returns a slice with the slip elements
		// [REPLACED] firstNames = append(firstNames, booking["firstName"])
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	// declaring variables without specify the values (you have to define the type)
	var firstName string
	var lastName string
	var email string
	var userTickets uint // uint is a datatype that only allows positive numbers

	// alternative syntax to create a slice:
	//   var bookings = []string{}
	// or
	//   bookings := []string{}

	// Asking user for inputs

	fmt.Println("Enter your firstname:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your lastname:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	// & represents a **pointer** (a variable that points to the memory address of another variable that references the actual value). Pointers in Go are also called special variables

	// If you use the scan without the pointer the method dont wait for the user input because its working with a copy from the variable and can´t change the original value

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {

	// 	[REPLACED] 
	//var userData = make(map[string]string) // map[keyDataType]valueDataType is used to declare a map (like a dictonary or hash) and the "make" is used to declare a empty map
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	var userData = UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		userTickets: userTickets
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

	remainingTickets = remainingTickets - userTickets
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
