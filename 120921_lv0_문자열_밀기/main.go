package main

import (
	"fmt"
	"time"
)

func solution(A string, B string) int {

	count := 0

	for A != B {
		A = A[len(A)-1:len(A)] + A[0:len(A)-1]
		count++
		if count > len(A) {
			return -1
		}
	}

	return count
}

func report(a string, b string, desireResult int) {

	input := fmt.Sprintf("Input:[%v, %v]", a, b)

	start := time.Now()

	result := solution(a, b)

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
	report("hello", "ohell", 1)
	report("apple", "elppa", -1)
	report("atat", "tata", 1)
	report("abc", "abc", 0)
}
