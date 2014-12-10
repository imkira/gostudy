package main

import (
	"fmt"
	"time"
)

func writer(base int, c chan int, done chan int) {
	for i := base; i <= base+3; i++ {
		time.Sleep(1 * time.Second)
		c <- i
	}
	close(done)
}

func main() {
	// start block OMIT
	c := make(chan int)

	done1 := make(chan int)
	go writer(5, c, done1)
	done2 := make(chan int)
	go writer(10, c, done2)

	for done1 != nil || done2 != nil {
		select {
		case <-done1:
			done1 = nil
		case <-done2:
			done2 = nil
		case x := <-c:
			fmt.Printf("read:%d\n", x)
		}
	}
	// end block OMIT
}
