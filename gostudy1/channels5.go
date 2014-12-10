package main

import (
	"fmt"
	"time"
)

func reader(id int, c chan int, done chan int) {
	for x := range c {
		fmt.Printf("reader:%d, read:%d\n", id, x)
		time.Sleep(1 * time.Second)
	}
	done <- 1
}

func writer(base int, c chan int, done chan int) {
	for i := base; i <= base+3; i++ {
		time.Sleep(1 * time.Second)
		c <- i
	}
	done <- 1
}

// start block OMIT

func main() {
	c := make(chan int)

	done1 := make(chan int)
	go writer(5, c, done1)
	done2 := make(chan int)
	go writer(10, c, done2)

	done3 := make(chan int)
	go reader(1, c, done3)
	done4 := make(chan int)
	go reader(2, c, done4)

	// wait
	<-done1
	<-done2
	close(c)
	<-done3
	<-done4
}

// end block OMIT
