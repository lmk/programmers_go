package main

import (
	"fmt"
	"strconv"
	"time"
)

func solution(bin1 string, bin2 string) string {

	b1, _ := strconv.ParseInt(bin1, 2, 64)
	b2, _ := strconv.ParseInt(bin2, 2, 64)

	return strconv.FormatInt(b1+b2, 2)
}

func report(bin1 string, bin2 string, desireResult string) {

	input := fmt.Sprintf("Input:[%v,%v]", bin1, bin2)

	start := time.Now()

	result := solution(bin1, bin2)

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
	report("10", "11", "101")
	report("1001", "1111", "11000")
}
