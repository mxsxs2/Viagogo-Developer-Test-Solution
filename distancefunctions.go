//Viagogo Developer Test Solution By Krisztian Nagy - mxsxs2@gmail.com
package main

import "math"

//The distance between an event and a give point
type EventDistanceToPoint struct {
	eventIndex int     //The index of the event in the event pool. This is not the same as event id
	distance   float64 //The distance to this event
}

//Function used to merge the event details and the sorted event distances. Returns the sorted events with distances
func mergeEventAndDistance(eventPool []Event, sortedDistancePool []EventDistanceToPoint) []Event {
	//Create a new pool for the events
	newEventPool := []Event{}

	//Loop the distances to pull the event details and sort the tickets in the events
	for _, distance := range sortedDistancePool {
		//Get the event by its index
		e := eventPool[distance.eventIndex]
		//Fill the distance
		e.currentDistance = distance.distance
		//Add the event to the new pool
		newEventPool = append(newEventPool, e)
	}
	//Return the new merged event pool
	return newEventPool
}

//Function used to get the Manhattan distance between two points
func getManhattanDistance(c1 Coordinates, c2 Coordinates) float64 {
	//Calculate and return distance
	return math.Abs(c2.x-c1.x) + math.Abs(c2.y-c1.y)
}

//Get the distance between a point and all of the events
func calculateDistancesBetweenEventsAndPoint(point Coordinates, eventPool []Event) []EventDistanceToPoint {
	//Create a new pool of distances
	newDistancePool := []EventDistanceToPoint{}
	//Get the event pool size
	sliceSize := len(eventPool)
	//Number of runs
	numOfRuns := int(math.Floor(float64(sliceSize) / 2.0))
	//Calculate number of runs
	if sliceSize%2 != 0 {
		numOfRuns = sliceSize/2 + 1
	}

	//Loop the events
	for i := 0; i < numOfRuns; i++ {
		//Get two event at a time. One from the begining and one from the end
		d1 := EventDistanceToPoint{
			//Copy the event index
			eventIndex: i,
			//Calculate the distances
			distance: getManhattanDistance(point, eventPool[i].location),
		}
		//Add the distance to the pool
		newDistancePool = append(newDistancePool, d1)

		d2 := EventDistanceToPoint{
			//Copy the event index
			eventIndex: sliceSize - i - 1,
			//Calculate the distances
			distance: getManhattanDistance(point, eventPool[sliceSize-i-1].location),
		}
		//Add the distance to the pool
		newDistancePool = append(newDistancePool, d2)

	}

	//Remove the last distance if the number of events is odd as this is a duplicate
	if sliceSize%2 != 0 {
		//Remove last distance
		newDistancePool = newDistancePool[:len(newDistancePool)-1]
	}

	//Return the distance pool
	return newDistancePool
}

//Sort the distances in ascending order. This is a modified merge sort (Merge sort is one of the most efficient sorting algorithm)
func sortEventDistances(distancePool []EventDistanceToPoint) []EventDistanceToPoint {
	//If the pool is shorter then one then return the pool
	if len(distancePool) <= 1 {
		return distancePool
	}
	//Cut the pool in half
	n := len(distancePool) / 2
	//Sort the first half
	l := sortEventDistances(distancePool[:n])
	//Sort the second half
	r := sortEventDistances(distancePool[n:])
	//Inner function for mergint the two array back together
	mergeEventDistances := func(l1 []EventDistanceToPoint, l2 []EventDistanceToPoint) []EventDistanceToPoint {
		//Initialize the new slice from the first slice
		var newlist []EventDistanceToPoint
		//If the size of the l1 or l2 are higher than 0
		for len(l1) > 0 || len(l2) > 0 {
			//If both are higher than 0
			if len(l1) > 0 && len(l2) > 0 {
				//If the first distance of the first array is lower
				if l1[0].distance <= l2[0].distance {
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
	return mergeEventDistances(l, r)
}
