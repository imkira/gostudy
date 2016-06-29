package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT
func runWorker(done chan struct{}) {
	n := rand.Intn(1000)
	time.Sleep(time.Duration(n) * time.Millisecond)
	fmt.Println(n)
	done <- struct{}{}
}

func main() {
	n := 10
	done := make(chan struct{}, n)
	for i := 0; i < n; i++ {
		go runWorker(done)
	}
	for i := 0; i < n; i++ {
		<-done
	}
}

// END OMIT
