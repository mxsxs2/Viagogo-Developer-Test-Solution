# Viagogo Developer Test App
The application is made for Viagogo's Developer Test. It is an event finder app which finds the closest events and their cheapest ticket to given coordinates (x;y)

### How to install and run GO

To install, simply go to GO's website and download the installer and run it: https://golang.org/

When the installation is done, the "go" command is going to be avaialable in terminal.(You have to restart an opened terminal)

### How to use this repository

To build the application, go clone this repository and in the folder run: 
```
go build
```

The previous command will compile the go file into a runnable.

Once the runnable is created then it can be run in terminal, for example on windows: 
```
./Viagogo-Developer-Test-Solution.exe
```

## How to use the application
Once tha application is loaded it will ask for the x and y coordinates. The coordinates should be given as "x,y" for example: 4,2.
If valid coordinates are present the app file find the closest events and show the first 5 of them with the cheapest ticket for that event.
### How to show generated data
The randomly generated data can be written onto the command line console with the -printdata argument to examine if the application works correctly:
```
./Viagogo-Developer-Test-Solution.exe -printdata
```

## Extendability
### How might you change your program if you needed to support multiple events at the same location?
The application can be changed to support multiple events at the same location. The ```eventgenerator.go``` file has a function call at line 53:
```
//Generate the unique location
location: generateRandomCoordinates(true, eventPool),
```

If the true is changed to false, the generator will not check for unique coordinates,so multiple events could occour at the same location.
The rest of the application is ready for multiple events at the same location.
### How would you change your program if you were working with a much larger world size?
In order to support larger world size, the ```config.go``` file has to be changed at line 8 and 11:
```
//The highest coordinate for both X and Y
const COORDMAX = 10

//The lowest coordinate for both X and Y
const COORDMIN = -10
```
The application will be able to handle the larger world without any errors, thanks to the efficient sorting algorithms. However if there are more than a few thousands of events are available, the app should be modified to store the events in a database, rather than in memory.


## Assumptions
* As the documentation did not clarify how many events and tickects should I have, I decided to generate random amount of them, so every time the app runs it generates completely different amunt of events and tickets
* The maximum and miminum were given but was not specified if the have to be whole numbers or floating point numbers. I used floating point numbers as in a real world scenarion, therefore when the distance is shown it is shown as a floating point number with a precision of two.
* I decided to give padding three for the event names and padding two to the distance and the price, this way it looks nicer an more organizer. The padding only added at the print out stage, it is not stored anywhere.