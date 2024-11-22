package main

import (
	"fmt"
	"strings"
)

func main() {
	//declaring main variables
	conferenceName := "Go conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	booking := []string{}

	//output
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d are still availaible\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		//User Inputs
		var fname string
		var lname string
		var email string
		var userTickets uint
		fmt.Printf("Enter firstname :- ")
		fmt.Scan(&fname)
		fmt.Printf("Enter lastname :- ")
		fmt.Scan(&lname)
		fmt.Printf("Enter email :- ")
		fmt.Scan(&email)
		fmt.Printf("Enter No. of tickets :- ")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			//Ticket booking logic
			remainingTickets = remainingTickets - userTickets
			booking = append(booking, fname+" "+lname)

			fmt.Printf("\nUser %v %v has booked %d tickets sent to there email %v", fname, lname, userTickets, email)
			fmt.Printf("\n%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range booking {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first name of bookings %v\n", firstNames)
			if remainingTickets == 0 {
				fmt.Println("All tickets are sold")
				break
			}
		} else {
			fmt.Printf("Entered tickets not availaible we have only %v availaible\n", remainingTickets)

		}
	}
}
