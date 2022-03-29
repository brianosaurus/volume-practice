package controllers

import (
	"fmt"
	"testing"
)

func TestStartAndDestination(t *testing.T) {
	result := [2]string{"", ""}

	result = StartAndDestination(FlightPaths{{"A", "B"}})
	if result[0] != "A" || result[1] != "B" {
		t.Errorf("Just one element: Expected %s, got %v", "[A,B]", result)
	} else {
		fmt.Printf("Just one element: Got %v\n", result)
	}

	result = StartAndDestination(FlightPaths{{"A", "B"}, {"B", "C"}})
	if result[0] != "A" || result[1] != "C" {
		t.Errorf("Two elements in order: Expected %s, got %v", "[A,C]", result)
	} else {
		fmt.Printf("Two elements in order: Got %v\n", result)
	}

	result = StartAndDestination(FlightPaths{{"B", "C"}, {"A", "B"}})
	if result[0] != "A" || result[1] != "C" {
		t.Errorf("Two elements out of order: Expected %s, got %v", "[A,C]", result)
	} else {
		fmt.Printf("Two elements out of order: Got %v\n", result)
	}

	result = StartAndDestination(FlightPaths{{"A", "B"}, {"E", "F"}})
	if result[0] != "" || result[1] != "" {
		t.Errorf("Two disjoint elements: Expected %s, got %v", "['','']", result)
	} else {
		fmt.Printf("Two disjoint elements: Got %v\n", result)
	}

	result = StartAndDestination(FlightPaths{{"A", "B"}, {"B", "C"}, {"C", "D"}})
	if result[0] != "A" || result[1] != "D" {
		t.Errorf("Three elements in order: Expected %s, got %v", "[A,D]", result)
	} else {
		fmt.Printf("Three elements in order: Got %v\n", result)
	}

	result = StartAndDestination(FlightPaths{{"A", "B"}, {"C", "D"}, {"B", "C"}})
	if result[0] != "A" || result[1] != "D" {
		t.Errorf("Three elements out of order: Expected %s, got %v", "[A,D]", result)
	} else {
		fmt.Printf("Three elements out of order: Got %v\n", result)
	}

	result = StartAndDestination(FlightPaths{{"A", "B"}, {"B", "A"}, {"B", "C"}, {"C", "D"}, {"D", "A"}})
	if result[0] != "" || result[1] != "" {
		t.Errorf("Multiple elements with loop: Expected %s, got %v", "['','']", result)
	} else {
		fmt.Printf("Multiple elements with loop: Got %v\n", result)
	}

	result = StartAndDestination(FlightPaths{{"A", "B"}, {"B", "C"}, {"C", "D"}, {"D", "A"}})
	if result[0] != "A" || result[1] != "A" {
		t.Errorf("Multiple elements, beginning and ending same airport: Expected %s, got %v", "[A,A]", result)
	} else {
		fmt.Printf("Multiple elements, beginning and ending same airport: Got %v\n", result)
	}

	result = StartAndDestination(FlightPaths{{"A", "B"}, {"B", "C"}, {"C", "D"}, {"D", "A"}, {"J", "K"}})
	if result[0] != "" || result[1] != "" {
		t.Errorf("Multiple airports, disjoint test 1: Expected %s, got %v", "['','']", result)
	} else {
		fmt.Printf("Multiple airports, disjoint test 1: Got %v\n", result)
	}

	result = StartAndDestination(FlightPaths{{"A", "B"}, {"B", "Z"}, {"B", "C"}, {"C", "D"}, {"D", "A"}, {"J", "K"}})
	if result[0] != "" || result[1] != "" {
		t.Errorf("Multiple airports, disjoint test 2: Expected %s, got %v", "['','']", result)
	} else {
		fmt.Printf("Multiple airports, disjoint test 2: Got %v\n", result)
	}

	result = StartAndDestination(FlightPaths{{"B", "C"}, {"C", "D"}, {"A", "B"}, {"D", "E"}})
	if result[0] != "A" || result[1] != "E" {
		t.Errorf("Start is in the middle: Expected %s, got %v", "[A,E]", result)
	} else {
		fmt.Printf("Start is in the middle: Got %v\n", result)
	}

	result = StartAndDestination(FlightPaths{{"B", "C"}, {"C", "D"}, {"D", "E"}, {"A", "B"}})
	if result[0] != "A" || result[1] != "E" {
		t.Errorf("Start is on the end: Expected %s, got %v", "[A,E]", result)
	} else {
		fmt.Printf("Start is on the end: Got %v\n", result)
	}

	result = StartAndDestination(FlightPaths{{"B", "C"}, {"C", "D"}, {"D", "E"}, {"A", "B"}})
	if result[0] != "A" || result[1] != "E" {
		t.Errorf("Start and end swapped: Expected %s, got %v", "[A,E]", result)
	} else {
		fmt.Printf("Start and end swapped: Got %v\n", result)
	}
}