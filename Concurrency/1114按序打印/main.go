package main

import (
	"fmt"
)

func main() {
	f := NewFoo()
	f.Start()
}

type Foo struct {
	ch1, ch2, ch3 chan struct{}
}

func NewFoo() *Foo {
	return &Foo{
		ch1: make(chan struct{}),
		ch2: make(chan struct{}),
		ch3: make(chan struct{}),
	}
}

func (f *Foo) Start() {
	go f.third()
	go f.second()
	go f.first()
	<-f.ch3
}

func (f *Foo) first() {
	fmt.Print("first")
	f.ch1 <- struct{}{}
}

func (f *Foo) second() {
	<-f.ch1
	fmt.Print("second")
	f.ch2 <- struct{}{}

}

func (f *Foo) third() {
	<-f.ch2
	fmt.Print("third")
	f.ch3 <- struct{}{}
}
