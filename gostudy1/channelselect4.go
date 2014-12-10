package main

// start block OMIT
type lock struct {
	c chan int
}

func newLock() *lock {
	return &lock{c: make(chan int, 1)}
}

func (l *lock) lock() {
	l.c <- 1
}

func (l *lock) unlock() {
	<-l.c
}

// end block OMIT

func main() {
}
