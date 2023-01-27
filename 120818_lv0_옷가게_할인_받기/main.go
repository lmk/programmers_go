package main

import (
	"fmt"
	"time"
)

func solution(price int) int {
	if price >= 500000 {
		price = int(float64(price) * 0.8)
	} else if price >= 300000 {
		price = int(float64(price) * 0.9)
	} else if price >= 100000 {
		price = int(float64(price) * 0.95)
	}

	return price
}

func report(price int, desireResult int) {

	input := fmt.Sprintf("Input:[%v]", price)

	start := time.Now()

	result := solution(price)

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
	report(150000, 142500)
	report(580000, 464000)
}
