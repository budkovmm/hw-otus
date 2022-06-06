package main

import (
	"fmt"
	"golang.org/x/example/stringutil"
)

// ReverseString function return reversed string
func ReverseString(s string) string {
	return stringutil.Reverse(s)
}

func main() {
	reversedString := ReverseString("Hello, OTUS!")
	fmt.Println(reversedString)
}
