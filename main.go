package main

import "fmt"

func main() {

	// variable declarations with datatype determined by the value assigned
	var conferenceName = "Coffee Conference"
	remainingTickets := 50 // a shorthand for version above
	const conferenceTickets = 50

	// printing strings using placeholders
	fmt.Printf("Welcom to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v ticket and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// declaring variables without specify the values (you have to define the type)
	var firstName string
	var lastName string
	var email string
	var userTickets int

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

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

}
