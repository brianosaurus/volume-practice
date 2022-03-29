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

func StartAndDestination(flightPaths FlightPaths) []string {
	var beginAndEnd []string

	for _, flightPath := range flightPaths {
		beginAndEnd = append(beginAndEnd, flightPath[0])
		beginAndEnd = append(beginAndEnd, flightPath[1])
	}

	return beginAndEnd
}
