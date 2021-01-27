package main

import (
	"fmt"
	"math/rand"
	"time"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func loadingAnimation(msg string, done chan struct{}) {
	fmt.Print(msg)

	for {
		select {
		case <-done:
			return
		case <-time.Tick(1 * time.Second):
			fmt.Print(".")
		}
	}
}
