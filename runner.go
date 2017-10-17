//Viagogo Developer Test Solution By Krisztian Nagy - mxsxs2@gmail.com
package main

import (
	"fmt"
)

func main() {
	//Generate the events
	events := generateSeedData()

	//Get the user coordinates
	usercoordinates := getCoordinatesFromUser()
	fmt.Printf("Closest Events to (%g,%g):\n", usercoordinates.x, usercoordinates.y)

	//Calculate the distances between the point and all of the events
	distancePool := calculateDistancesBetweenEventsAndPoint(usercoordinates, events)
	//Sort the distances and get only the required amount of them
	sortedDistancePool := sortEventDistances(distancePool)[:NUMBEROFRESULTS]

	//Get the sorted events with the distances
	closestEvents := mergeEventAndDistance(events, sortedDistancePool)

	//Print the result
	printEvents(closestEvents)

	//if the -printdata argument is present then print the generated data
	if isFlagOn("-printdata") {
		//Print the seed data
		printGeneratedData(events)
	}

}
