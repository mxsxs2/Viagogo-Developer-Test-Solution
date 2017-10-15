//Viagogo Developer Test Solution By Krisztian Nagy - mxsxs2@gmail.com
package main

import (
	"bufio"
	"fmt"
	"os"
)

//Function used to print the sorted event
func printEvents(eventPool []Event) {
	//Loop the events
	for _, event := range eventPool {
		//Print the details of the event
		printSingleEvent(event)
	}
}

/*Function used to print a single event with the cheapest ticket and with the distance
The event id gets 2 leading zeros
The price gets the currency symbol and the precision is set to 2
The distance is round and casted to int so it is a whole number instead of a float*/
func printSingleEvent(event Event) {

	//Pad the event number to have the leading zeros
	eid := fmt.Sprintf("%03d", event.eventID)

	//Format price
	price := fmt.Sprintf("%s%05.2f", BASECURRENCY, event.tickets[0].price)

	//Format the distance
	distance := fmt.Sprintf("%05.2f", event.currentDistance)

	//Print the event details
	fmt.Printf("Event %s - %s, Distance %s\n", eid, price, distance)
}

//Function used to get coordinates from user
func getCoordinatesFromUser() Coordinates {
	fmt.Println("Please Input Coordinates:")
	//Read and return the coordinates
	return readCoordinates()
}

//Function used to read in a pair of coordinates with input validation
func readCoordinates() Coordinates {
	//Create the holder
	var coordiates Coordinates
	//Specify the reader
	stdin := bufio.NewReader(os.Stdin)
	//Loop until there is a valid input
	for {
		//Read the preformatted coordinates
		_, err := fmt.Fscanf(stdin, "%g,%g\n", &coordiates.x, &coordiates.y)
		//If there was no error and the coordinates are within the world boundaries then break the loop
		if err == nil && int(coordiates.x) >= COORDMIN && int(coordiates.x) <= COORDMAX && int(coordiates.y) >= COORDMIN && int(coordiates.y) <= COORDMAX {
			break
		}
		//Read until the end of the line
		stdin.ReadString('\n')
		//If there was an error then let the user know about it
		fmt.Println("Invalid input. Please enter the coordinates again again:")
	}
	//Return the coordinates
	return coordiates
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
