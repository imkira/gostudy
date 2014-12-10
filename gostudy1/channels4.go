package main

import (
	"fmt"
	"time"
)

func writer(c chan int) {
	for i := 1; i <= 5; i++ {
		time.Sleep(1 * time.Second)
		c <- i
	}
	close(c)
}

// start block OMIT

func reader(id int, c chan int, done chan int) {
	for x := range c {
		fmt.Printf("worker:%d, read:%d\n", id, x)
		time.Sleep(1 * time.Second)
	}
	done <- 1
}

func main() {
	c := make(chan int)
	go writer(c)

	done1 := make(chan int)
	go reader(1, c, done1)
	done2 := make(chan int)
	go reader(2, c, done2)
	done3 := make(chan int)
	go reader(3, c, done3)
	// wait
	<-done1
	<-done2
	<-done3
}

// end block OMIT
