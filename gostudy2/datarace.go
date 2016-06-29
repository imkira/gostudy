package main

import (
	"fmt"
	"sync"
)

// START OMIT
func foo(total *int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		*total++
	}
}

func main() {
	var wg sync.WaitGroup
	var val int
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go foo(&val, &wg)
	}
	wg.Wait()
	fmt.Printf("TOTAL: %d\n", val)
}

// END OMIT
