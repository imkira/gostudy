package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// START OMIT
func foo(total *int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		atomic.AddInt64(total, 1)
	}
}

func main() {
	var wg sync.WaitGroup
	var val int64
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go foo(&val, &wg)
	}
	wg.Wait()
	fmt.Printf("TOTAL: %d\n", val)
}

// END OMIT
