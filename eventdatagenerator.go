//Viagogo Developer Test Solution By Krisztian Nagy - mxsxs2@gmail.com
package main

import (
	"math/rand"
	"time"
)

//Structure for a single pair of coordinates
type Coordinates struct {
	//The coordinates set to be floats as in realworld they are floats as well
	x float64
	y float64
}

//Structure for a single ticket
type Ticket struct {
	ticketID int     //I assume every ticket should have an identifier
	price    float64 //The non zero price
}

//Structure for a single event
type Event struct {
	eventID         int         //Unique event identifier
	tickets         []Ticket    //The set of tickets
	location        Coordinates //The event location
	currentDistance float64     //The current distance from a given point. This is populated every time a user needs the distance
}

//Function used to generate random numbers
func random(min, max int) float64 {
	//Create the random number in boundaries.
	return rand.Float64()*float64((max-min)) + float64(min)
}

//Function used to generate random events
func generateSeedData() []Event {
	//Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	//The empty holder for the events
	eventPool := []Event{}
	//The number of generated events. It is down casted to int for the loop
	numofEvents := int(random(MINSEEDDATA, MAXSEEDDATA))

	//Do a loop to generate the events
	for i := 1; i <= numofEvents; i++ {
		//Create an event
		event := Event{
			//Set the event id, which is unique as it is the iteration number
			eventID: i,
			//Generate the unique location
			location: generateRandomCoordinates(true, eventPool),
			//Generate a random number of tickets
			tickets: generateRandomTickets(),
		}

		//Add the event to the pool
		eventPool = append(eventPool, event)
	}
	//Return the events
	return eventPool
}

/*Function used to generate random coordinates
  The coordinates can be either unique or not
  If the the coordinates has to be uniqe, the function should receive an event slice to compare the coordinates to, otherwise just an empty event slice is enough*/
func generateRandomCoordinates(unique bool, eventPool []Event) Coordinates {
	//Create new coordinates holder
	coordinates := Coordinates{
		//Generate the x coordinate
		x: random(COORDMIN, COORDMAX),
		//Generate the y coordinate
		y: random(COORDMIN, COORDMAX),
	}

	//If the unique flag is on then loop the event pool to check if it is unique or not
	if unique {
		//Loop the events
		for _, event := range eventPool {
			//If the coordinates exists
			if event.location == coordinates {
				//Call the random coordinates generator again and return the value of it
				return generateRandomCoordinates(unique, eventPool)
			}
		}
	}

	//If the coordinates does not have to be uniqe then just return the generated ones
	return coordinates
}

/*Function used to generate random amount of tickets
  It may generate the same ticket price once or more as there was no requirement specified
  The ticket id maybe the same for one or more events, therefore when we are talking about a uniqe ticket is should be in the format of:
  Eventid:TicketID
*/
func generateRandomTickets() []Ticket {
	//The empty holder for the tickets
	ticketPool := []Ticket{}
	//The number of generated tickets. It is down casted to int for the loop
	numofTickets := int(random(MINNUMBEROFTICKETS, MAXNUMBEROFTICKETS))
	//Do a loop to generate the tickests
	for i := 1; i <= numofTickets; i++ {
		ticket := Ticket{
			//Set the event id, which is unique for the current event as it is the iteration number
			ticketID: i,
			//Randomly genertate a non-zero ticket price
			price: random(1, MAXTICKETPRICE),
		}
		//Add the ticket to the ticket pool
		ticketPool = append(ticketPool, ticket)
	}

	//Sort the tickets by price in ascending order
	//I sor the tickets here so it does not have to be sorted every time a user checks them
	ticketPool = sortTickets(ticketPool)

	//Return the tickest
	return ticketPool
}

//Sort the distances in ascending order. This is a modified merge sort (Merge sort is one of the most efficient sorting algorithm)
func sortTickets(ticketPool []Ticket) []Ticket {
	//If the pool is shorter then one then return the pool
	if len(ticketPool) <= 1 {
		return ticketPool
	}
	//Cut the pool in half
	n := len(ticketPool) / 2
	//Sort the first half
	l := sortTickets(ticketPool[:n])
	//Sort the second half
	r := sortTickets(ticketPool[n:])
	//Inner function for mergint the two array back together
	mergeTickets := func(l1 []Ticket, l2 []Ticket) []Ticket {
		//Initialize the new slice from the first slice
		var newlist []Ticket
		//If the size of the l1 or l2 are higher than 0
		for len(l1) > 0 || len(l2) > 0 {
			//If both are higher than 0
			if len(l1) > 0 && len(l2) > 0 {
				//If the first price of the first array is lower
				if l1[0].price <= l2[0].price {
					//Add to the new list
					newlist = append(newlist, l1[0])
					//Resize the slice
					l1 = l1[1:len(l1)]
				} else {
					//Add to the new list
					newlist = append(newlist, l2[0])
					//Resize the slice
					l2 = l2[1:len(l2)]
				}
				//If just the first array is higher then 0
			} else if len(l1) > 0 {
				newlist = append(newlist, l1[0])
				l1 = l1[1:len(l1)]
				//If just the second array is higher then 0
			} else if len(l2) > 0 {
				newlist = append(newlist, l2[0])
				l2 = l2[1:len(l2)]
			}
		}

		//Retrun the new sorted array
		return newlist
	}
	//Recursively call the merhe function then return the sorted array
	return mergeTickets(l, r)
}
