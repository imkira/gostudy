package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// START OMIT
func runWorker(wg *sync.WaitGroup) {
	n := rand.Intn(1000)
	time.Sleep(time.Duration(n) * time.Millisecond)
	fmt.Println(n)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go runWorker(&wg)
	}
	wg.Wait()
}

// END OMIT
