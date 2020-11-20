package main

import (
	"fmt"
	"sync"
)

func main() {
	fb := newFooBar(4)
	fb.wg.Add(2)
	go fb.foo()
	go fb.bar()
	fb.wg.Wait()
}

// FooBar struct
type FooBar struct {
	n         int
	fooDoneCh chan struct{}
	barDoneCh chan struct{}
	wg        sync.WaitGroup
}

func newFooBar(n int) *FooBar {
	return &FooBar{
		n:         n,
		fooDoneCh: make(chan struct{}, 1),
		barDoneCh: make(chan struct{}, 1),
		wg:        sync.WaitGroup{},
	}
}

func (fb *FooBar) foo() {
	fmt.Print("foo")
	fb.fooDoneCh <- struct{}{}
	for i := 1; i < fb.n; i++ {
		<-fb.barDoneCh
		fmt.Print("foo")
		fb.fooDoneCh <- struct{}{}
	}
	fb.wg.Done()
}

func (fb *FooBar) bar() {
	for i := 0; i < fb.n; i++ {
		<-fb.fooDoneCh
		fmt.Print("bar")
		fb.barDoneCh <- struct{}{}
	}
	fb.wg.Done()
}
