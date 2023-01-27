package main

import (
	"fmt"
	"math/big"
	"time"
)

func calcFactorial(factorialnum int) *big.Int {
	one := big.NewInt(1)
	factorial := big.NewInt(1)
	max := big.NewInt(int64(factorialnum))
	for i := big.NewInt(1); i.Cmp(max) != 1; i.Add(i, one) {
		factorial.Mul(factorial, i)
	}
	return factorial
}

func solution(balls int, share int) int {

	b := calcFactorial(balls)
	bs := calcFactorial(balls - share)
	s := calcFactorial(share)
	b.Div(b, bs.Mul(bs, s))

	return int(b.Int64())
}

func report(balls int, share int, desireResult int) {

	input := fmt.Sprintf("Input:[%v, %v]", balls, share)

	start := time.Now()

	result := solution(balls, share)

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
	report(3, 2, 3)
	report(5, 3, 10)
}
