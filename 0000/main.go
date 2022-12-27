package main

import (
	"fmt"
	"strings"
	"time"
)

func solution(n int) string {
	answer := ""
	for i := 0; i < n; i++ {
		answer += fmt.Sprintln(strings.Repeat("*", i+1))
	}

	return answer
}

func report(n int, desireResult string) {

	input := fmt.Sprintf("Input:[%v]", n)

	start := time.Now()

	result := solution(n)

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
	report(3, "*\n**\n***\n")
	report(5, "*\n**\n***\n****\n*****\n")
}
