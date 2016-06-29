package main

import (
	"fmt"
	"sync"
)

type Foo struct {
	l1 sync.Mutex
	d1 int

	l2 sync.Mutex
	d2 int
}

// START OMIT
func (f *Foo) inc1dec2() {
	f.l1.Lock()
	f.l2.Lock()
	f.d1++
	f.d2--
	f.l2.Unlock()
	f.l1.Unlock()
}

func (f *Foo) inc2dec1() {
	f.l1.Lock()
	f.l2.Lock()
	f.d2++
	f.d1--
	f.l2.Unlock()
	f.l1.Unlock()
}

// END OMIT

func main() {
	var f Foo
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			f.inc1dec2()
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			f.inc2dec1()
		}
	}()
	wg.Wait()
	fmt.Println("FINISHED", f.d1, f.d2)
}
