package main

import (
	"fmt"
	"strings"
)

func main() {
	var paird, valued string
	var str string
	fmt.Scanln(&paird, &valued, &str)

	pairs := strings.Split(str, paird)
	count := 0
	kvs := [][2]string{}
	for _, pair := range pairs {
		kv := strings.Split(pair, valued)
		if len(kv) == 2 {
			count++
			kvs = append(kvs, [2]string{kv[0], kv[1]})
		}
	}
	fmt.Println(count)
	for _, kv := range kvs {
		fmt.Println(kv[0], kv[1])
	}
}
