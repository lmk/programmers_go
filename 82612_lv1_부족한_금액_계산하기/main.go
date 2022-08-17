package main

import (
	"fmt"
	"time"
)

func solution(price int, money int, count int) int64 {

	sum := int64(0)
	for i := 1; i <= count; i++ {
		sum += int64(price * i)
	}

	if int64(money) > sum {
		return 0
	}

	return sum - int64(money)
}

func report(price int, money int, count int, desireResult int64) {

	start := time.Now()

	result := solution(price, money, count)

	duration := time.Since(start)

	if result == desireResult {
		fmt.Printf("[O] ")
	} else {
		fmt.Printf("[X] ")
	}

	fmt.Printf("Input:[%v %v %v], result:[%v], duration:%v\n", price, money, count, result, duration)
}

func main() {

	report(3, 20, 4, 10)
}
