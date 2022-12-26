package main

import (
	"fmt"
	"strconv"
	"time"
)

func solution(t string, p string) int {
	size := len(p)
	n, _ := strconv.Atoi(p)

	result := 0

	for i := 0; i <= len(t)-size; i++ {
		v, _ := strconv.Atoi(t[i : i+size])
		if n >= v {
			result++
		}
	}

	return result
}

func report(t string, p string, desireResult int) {

	input := fmt.Sprintf("Input:[%v, %v]", t, p)

	start := time.Now()

	result := solution(t, p)

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
	report("3141592", "271", 2)
	report("500220839878", "7", 8)
	report("10203", "15", 3)
}
