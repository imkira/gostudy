package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT
func runWorker() {
	n := rand.Intn(1000)
	time.Sleep(time.Duration(n) * time.Millisecond)
	fmt.Println(n)
}

func main() {
	for i := 0; i < 10; i++ {
		go runWorker()
	}
}

// END OMIT
