package birdazzone

import (
	"testing"
	"time"
)

func TestGhihliottinaTracker(t *testing.T) {
	res := GetBirdazzoneTracker()
	if res.Game != birdazzoneTracker.Game {
		t.Fatal("Wrong game")
	}
	if res.Query != birdazzoneTracker.Query {
		t.Fatal("Wrong query")
	}
}

func TestSolution(t *testing.T) {
	// checking last solution
	sol, err := lastSolution()
	if err != nil {
		t.Fatal(err.Error())
	}
	if sol.Key == "" {
		t.Fatal("Empty solution #1")
	}
	now := time.Now()
	// checking tomorrow's solution
	tm := now.AddDate(0, 0, 1)
	_, err = givenSolution(tm)
	if err == nil {
		t.Fatal("Didn't get expected error")
	}
}
