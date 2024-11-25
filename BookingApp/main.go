package main

import (
	"fmt"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = conferenceTickets

	// Welcome message
	fmt.Println("Welcome to the", conferenceName, "booking application!")
	fmt.Printf("We have a total of %d tickets, and %d are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend!")

	// Slice to store bookings
	var bookings []string

	for {
		// Gather user information
		var firstName, lastName, email string
		var userTickets int

		fmt.Println("Enter your first name:")
		fmt.Scanln(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scanln(&lastName)

		fmt.Println("Enter your email address:")
		fmt.Scanln(&email)

		fmt.Println("Enter the number of tickets you want to book:")
		fmt.Scanln(&userTickets)

		// Check ticket availability
		if userTickets <= remainingTickets && userTickets > 0 {
			// Process booking
			remainingTickets -= userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %s %s for booking %d tickets. You will receive a confirmation email at %s.\n", firstName, lastName, userTickets, email)
			fmt.Printf("%d tickets remaining for %s.\n", remainingTickets, conferenceName)

			// Print all bookings
			fmt.Println("The bookings are:")
			for _, booking := range bookings {
				fmt.Println(" -", booking)
			}

			// Check if tickets are sold out
			if remainingTickets == 0 {
				fmt.Println("Our conference is sold out. Please come back next year!")
				break
			}
		} else {
			fmt.Printf("Invalid number of tickets. We have only %d tickets remaining.\n", remainingTickets)
		}
	}
}
