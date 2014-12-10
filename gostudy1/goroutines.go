package main

import (
	"fmt"
	"time"
)

// start block OMIT
func hoge() {
	fmt.Println("hoge: start")
	time.Sleep(2 * time.Second)
	fmt.Println("hoge: end")
}

func fuga() {
	fmt.Println("fuga: start")
	time.Sleep(1 * time.Second)
	fmt.Println("fuga: end")
}

func main() {
	go hoge()
	go fuga()
	time.Sleep(3 * time.Second)
}

// end block OMIT
