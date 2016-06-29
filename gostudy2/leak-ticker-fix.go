package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT
func runTicker() {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	// END OMIT
	for now := range ticker.C {
		fmt.Println(now)
		if rand.Intn(100) < 25 {
			fmt.Println("DONE")
			return
		}
	}
}

func main() {
	runTicker()
}
