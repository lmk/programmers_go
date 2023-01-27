package main

import (
	"fmt"
	"time"
)

func solution(cipher string, code int) string {

	answer := ""
	for i := code; i <= len(cipher); i += code {
		answer += string(cipher[i-1])
	}
	return answer
}

func report(cipher string, code int, desireResult string) {

	input := fmt.Sprintf("Input:[%v, %v]", cipher, code)

	start := time.Now()

	result := solution(cipher, code)

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
	report("dfjardstddetckdaccccdegk", 4, "attack")
	report("dfjardstddetckdaccccdegk", 4, "attack")
}
