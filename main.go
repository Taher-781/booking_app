package main

import "fmt"

func main() {
	conferenceName := "Go conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d are still availaible\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

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

	remainingTickets = remainingTickets - userTickets
	fmt.Printf("\nUser %v %v has booked %d tickets sent to there email %v", fname, lname, userTickets, email)
	fmt.Printf("\n%v tickets remaining for %v", remainingTickets, conferenceName)
}
