package main

import "fmt"

func main() {

	s := []int{1, 2, 3, 4}
	fmt.Println(len(s), cap(s))
	s = s[:1]
	fmt.Println(len(s), cap(s))
	s = s[:0]
	fmt.Println(len(s), cap(s))
}
