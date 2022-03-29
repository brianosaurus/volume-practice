package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/brianosaurus/volume-practice/api/utils/formaterror"
)

//flightPaths The structure that each listing sits in
type FlightPaths [][2]string


//CreateflightPaths makes a estate
func (server *Server) GetStartAndDestination(c *gin.Context) {
	//clear previous error if any
	errList := map[string]string{}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	flightPaths := FlightPaths{}

	fmt.Println("body: ", string(body))

	err = json.Unmarshal(body, &flightPaths)
	if err != nil {
		errList["Unmarshal_error"] = "Cannot unmarshal body"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	fmt.Printf("FlightPaths are %v\n", flightPaths)

	beginAndEnd := StartAndDestination(flightPaths)

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": beginAndEnd,
	})
}

func StartAndDestination(flightPaths FlightPaths) [2]string {
	startingAirports := make(map[string][2]string)
	endingAirports := make(map[string][2]string)

	// first loop
	for _, flight := range flightPaths {
		// we've seen the starting point before, therefore there is a loop
		if _, ok := startingAirports[flight[0]]; ok {
			return [2]string{"", ""}
		}

		startingAirports[flight[0]] = flight

		if _, ok := endingAirports[flight[1]]; ok {
			return [2]string{"", ""}
		}

		endingAirports[flight[1]] = flight
	}

	sortedFlightPaths := make(FlightPaths, 1)
	sortedFlightPaths[0] = flightPaths[0]

	// first loop going right
	i := 1
	for ; i < len(flightPaths); i++ {
		curr := sortedFlightPaths[len(sortedFlightPaths)-1]
		if flight, ok := startingAirports[curr[1]]; ok && 
			(flight[0] != sortedFlightPaths[0][0] &&  // this check is for loops
			flight[1] != sortedFlightPaths[0][1]) {
			// if there is a flight that starts at the ending point of this flight
			sortedFlightPaths = append(sortedFlightPaths, flight)
		} else {
			break
		}
	}

	// second loop going left
	for ; i < len(flightPaths); i++ {
		curr := sortedFlightPaths[0]
		if flight, ok := endingAirports[curr[0]]; ok &&
			(flight[0] != sortedFlightPaths[len(sortedFlightPaths)-1][0] && // this check is for loops
			flight[1] != sortedFlightPaths[len(sortedFlightPaths)-1][1]) {
			sortedFlightPaths = append(FlightPaths{flight}, sortedFlightPaths...)
		} else {
			break
		}
	}

	// somehow we didn't find a cogent flight path
	if len(flightPaths) != len(sortedFlightPaths) {
		return [2]string{"", ""}
	}

	return [2]string{sortedFlightPaths[0][0], sortedFlightPaths[len(sortedFlightPaths)-1][1]}
}