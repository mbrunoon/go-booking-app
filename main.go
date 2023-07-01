package main

import (
	"fmt"
	"strings"
)

func main() {

	// variable declarations with datatype determined by the value assigned
	const conferenceTickets = 50
	var remainingTickets uint = 50
	conferenceName := "Coffee Conference" // a shorthand for variable declaration with a datapyte

	// var bookings [50]string // array with 50 positions and no initial value
	var bookings []string // slice: dynamic array (no lenght defined)

	// printing strings using placeholders
	fmt.Printf("Welcom to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v ticket and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for remainingTickets > 0 && len(bookings) < 50 {

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

		if userTickets <= remainingTickets {

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings { // user _ to especify unused variables or those ones you don't want to use
				var names = strings.Fields(booking) // splits the string with white space as separator and returns a slice with the slip elements
				firstNames = append(firstNames, names[0])
			}
			// range expression allow us iterate over the elements and data structures

			fmt.Printf("The firsts names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}

		} else {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickers\n", remainingTickets, userTickets)
		}

	}

}
