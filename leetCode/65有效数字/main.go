package main

import "regexp"

var reg = regexp.MustCompile(`^[+-]?(?:\d+\.?\d*|\.\d+)(?:[eE][+-]?\d+)?$`)

func isNumber(s string) bool {
	return reg.MatchString(s)
}
