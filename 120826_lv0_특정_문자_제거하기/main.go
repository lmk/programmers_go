package main

import (
	"fmt"
	"strings"
	"time"
)

func solution(my_string string, letter string) string {
	return strings.ReplaceAll(my_string, letter, "")
}

func report(my_string string, letter string, desireResult string) {

	input := fmt.Sprintf("Input:[%v, %v]", my_string, letter)

	start := time.Now()

	result := solution(my_string, letter)

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
	report("abcdef", "f", "abcde")
	report("BCBdbe", "B", "Cdbe")
}
