package main

import (
	"fmt"
	"sync"
)

// START OMIT
func main() {
	var val int
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		days := []string{"mon", "tue", "wed", "thu", "fri", "sat", "sun"}
		for i := 0; i < 100000; i++ {
			if val == 0 {
				fmt.Println(days[val])
			}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			val = i % 7
		}
	}()
	wg.Wait()
}

// END OMIT
