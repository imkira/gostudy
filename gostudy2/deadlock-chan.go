package main

import (
	"fmt"
	"sync"
)

// START OMIT
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	c1 := make(chan int)
	c2 := make(chan int)
	go func() {
		defer wg.Done()
		for i := range c1 {
			fmt.Println("c1:", i)
			c2 <- i + 1
		}
	}()
	go func() {
		defer wg.Done()
		for i := range c2 {
			fmt.Println("c2:", i)
			c1 <- i + 1
		}
	}()
	c1 <- 1
	c1 <- 1
	wg.Wait()
}

// END OMIT
