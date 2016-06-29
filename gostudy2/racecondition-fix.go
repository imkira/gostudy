package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// START OMIT
func main() {
	var val int64
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		days := []string{"mon", "tue", "wed", "thu", "fri", "sat", "sun"}
		for i := 0; i < 100000; i++ {
			if val2 := atomic.LoadInt64(&val); val2 == 0 {
				fmt.Println(days[val2])
			}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			atomic.StoreInt64(&val, int64(i%7))
		}
	}()
	wg.Wait()
}

// END OMIT
