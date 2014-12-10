package main

// start block OMIT
func notifyAll(observers []chan int, val int) {
	for _, o := range observers {
		o <- val
	}
}

// end block OMIT

func main() {
}
