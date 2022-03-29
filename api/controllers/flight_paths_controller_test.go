package controllers

import (
	"fmt"
	"testing"
)

func TestStartAndDestination(t *testing.T) {
	result := StartAndDestination(FlightPaths{{"A", "B"}, {"B", "C"}, {"C", "D"}})
	if result[0] != "A" || result[1] != "D" {
		t.Errorf("Expected %s, got %v", "[A,D]", result)
	} else {
		fmt.Printf("Got %v\n", result)
	}
}