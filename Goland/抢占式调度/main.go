package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	done := make(chan bool, 1)
	var mu sync.Mutex
	var g1 int

	// goroutine 1
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				mu.Lock()
				g1++
				time.Sleep(100 * time.Microsecond)
				mu.Unlock()
			}
		}
	}()

	// goroutine 2
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Microsecond)
		mu.Lock()
		mu.Unlock()
	}
	done <- true

	fmt.Println(g1)
}
