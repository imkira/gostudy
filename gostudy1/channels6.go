package main

func main() {
	// start block OMIT
	c := make(chan int, 3)
	close(c)
	close(c)
	// end block OMIT
}
