package main

import (
	"fmt"
	"strings"
	"time"
)

func solution(num int) string {

	result := ""
	for i := 0; i < num; i++ {
		result += fmt.Sprintln(strings.Repeat("*", i+1))
	}

	return result
}

func report(num int, desireResult string) {

	input := fmt.Sprintf("Input:[%v]", num)

	start := time.Now()

	result := solution(num)

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
}
