package main

import (
	"fmt"
)

func main() {
	s := "how do you do"
	fmt.Println(lengthOfTheLastWorld(s))
}

func lengthOfTheLastWorld(s string) int {
	length := 0
	x := true
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == ' ' && x == true {
			continue
		}
		if s[i] == ' ' {
			return length
		}
		length++
		x = false
	}
	return length
}
