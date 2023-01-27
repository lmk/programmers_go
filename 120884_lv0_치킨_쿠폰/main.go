package main

import (
	"fmt"
	"time"
)

func solution(chicken int) int {
	sc := 0

	for chicken >= 10 {
		sc += chicken / 10
		chicken = (chicken % 10) + (chicken / 10)
	}
	return sc
}

func report(chicken int, desireResult int) {

	input := fmt.Sprintf("Input:[%v]", chicken)

	start := time.Now()

	result := solution(chicken)

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
	report(100, 11)
	report(1081, 120)
}
