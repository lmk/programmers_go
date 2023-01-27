package main

import (
	"fmt"
	"strings"
	"time"
)

func solution(before string, after string) int {
	a := strings.Split(after, "")

	for _, s := range a {
		tobe := strings.Replace(before, s, "", 1)
		if tobe == before {
			return 0
		}
		before = tobe
	}

	return 1
}

func report(before string, after string, desireResult int) {

	input := fmt.Sprintf("Input:[%v,%v]", before, after)

	start := time.Now()

	result := solution(before, after)

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
	report("olleh", "hello", 1)
	report("allpe", "apple", 0)
}
