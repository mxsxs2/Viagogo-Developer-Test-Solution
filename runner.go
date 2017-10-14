//Viagogo Developer Test Solution By Krisztian Nagy - mxsxs2@gmail.com
package main

import (
	"fmt"
)

func main() {
	//Generate the events
	events := generateSeedData()
	//Print the seed data
	printGeneratedData(events)

}

//Function used to print the generated data
func printGeneratedData(events []Event) {
	fmt.Println("Number of events:", len(events))
	fmt.Println("")
	//Loop each event
	for _, event := range events {
		fmt.Println("Event id:", event.eventID)
		fmt.Println("Event location:(", event.location.x, ";", event.location.y, ")")
		fmt.Println("Number of tickets in this event:", len(event.tickets))
		//Loop the tickets in the events
		for _, ticket := range event.tickets {
			fmt.Println("\tTicket id:", ticket.ticketID)
			fmt.Println("\tTicket price:", ticket.price)
		}
		fmt.Println("")
	}
}
