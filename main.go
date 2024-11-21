package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d are still availaible\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}
