package main

import (
	"fmt"
	"strings"
	"time"
)

func solution(my_string string, n int) string {

	s := ""
	for _, r := range my_string {
		s += strings.Repeat(string(r), n)
	}

	return s
}

func report(my_string string, n int, desireResult string) {

	input := fmt.Sprintf("Input:[%v, %v]", my_string, n)

	start := time.Now()

	result := solution(my_string, n)

	duration := time.Since(start)

	isSame := true
	if result != desireResult {
		isSame = false
	}

	if isSame {
		fmt.Printf("[O] %s result:[%v] duration:%v\n", input, result, duration)
	} else {
		fmt.Printf("[X] %s result:[%v]->[%v] duration:%v\n", input, result, desireResult, duration)
	}
}

func main() {
	report("hello", 3, "hhheeellllllooo")
}
