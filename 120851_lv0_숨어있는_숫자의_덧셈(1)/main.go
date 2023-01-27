package main

import (
	"fmt"
	"time"
)

func solution(my_string string) int {
	sum := 0
	for _, r := range my_string {
		if r >= '0' && r <= '9' {
			sum += int(r - '0')
		}
	}
	return sum
}

func report(my_string string, desireResult int) {

	input := fmt.Sprintf("Input:[%v]", my_string)

	start := time.Now()

	result := solution(my_string)

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
	report("aAb1B2cC34oOp", 10)
	report("1a2b3c4d123", 16)
}
