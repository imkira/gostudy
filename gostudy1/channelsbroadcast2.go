package main

import (
	"fmt"
	"time"
)

// start block OMIT
type state struct {
	val  int
	next *state
	done chan int
}

func (s *state) update(val int) *state {
	s.next = &state{val: val, done: make(chan int)}
	close(s.done)
	return s.next
}

func observer(id int, s *state) {
	for {
		fmt.Printf("observer:%d val:%d\n", id, s.val)
		<-s.done
		s = s.next
	}
}

// end block OMIT

// start main OMIT
func main() {
	s := &state{val: 123, done: make(chan int)}

	go observer(1, s)
	go observer(2, s)
	go observer(3, s)

	time.Sleep(1 * time.Second)
	s = s.update(4)
	time.Sleep(1 * time.Second)
	s = s.update(8)
	time.Sleep(1 * time.Second)
	s = s.update(16)
	time.Sleep(1 * time.Second)
}

// end main OMIT
