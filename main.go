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
	i := len(s) - 1
	for i >= 0 && s[i] == ' ' {
		i--
	}

	for i >= 0 && s[i] != ' ' {
		length++
		i--
	}
	return length
}
