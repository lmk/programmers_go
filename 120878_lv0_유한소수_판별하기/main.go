package main

import (
	"fmt"
	"time"
)

func gcd(a int, b int) int {
	for a > 0 {
		a, b = b%a, a
	}

	return b
}

func fac(n int, i int) int {
	pre := n
	for n > 1 {
		if n%i != 0 {
			return n
		}
		pre = n
		n = n / i
	}

	return pre
}

func solution(a int, b int) int {

	// 기약분수
	g := gcd(a, b)
	a, b = a/g, b/g

	fmt.Printf("%v %v %v ", g, a, b)

	// 정수
	if a%b == 0 {
		fmt.Printf("I ")
		return 1
	}

	// 기약분수의 분모
	if b == 2 || b == 5 {
		fmt.Printf("A ")
		return 1
	}

	b = fac(b, 2)
	if b == 0 {
		fmt.Printf("B ")
		return 1
	}

	bt := fac(b, 3)
	if b != bt {
		fmt.Printf("C ")
		return 2
	}
	b = bt

	bt = fac(b, 4)
	if b != bt {
		fmt.Printf("D ")
		return 2
	}
	b = bt

	b = fac(b, 5)
	if b == 0 {
		fmt.Printf("E ")
		return 1
	}

	if b == 2 || b == 5 {
		fmt.Printf("F %d", b)
		return 1
	}

	fmt.Printf("G %d", b)
	return 2
}

func report(a int, b int, desireResult int) {

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
	// report(7, 20, 1)
	// report(11, 22, 1)
	// report(12, 21, 2)
	// report(3500, 500, 1)
	// report(25, 30, 2)
	// report(12, 36, 2)
	// report(1, 16, 1)
	// report(1, 32, 1)
	// report(1, 512, 1)
	// report(1, 11, 2)
	// report(1, 17, 2)
	// report(1, 30, 1)
	// report(42, 6, 1)
	// report(1, 1, 1)

	report(1, 1, 1)
	report(1, 2, 1)
	report(1, 3, 2)
	report(1, 4, 1)
	report(1, 5, 1)
	report(1, 6, 2)
	report(1, 7, 2)
	report(1, 8, 1)
	report(1, 9, 2)
	report(1, 10, 1)
	report(1, 11, 2)
	report(1, 12, 2)
	report(1, 13, 2)
	report(1, 14, 2)
	report(1, 15, 2)
	report(1, 16, 1)
	report(1, 17, 2)
	report(1, 18, 2)
	report(1, 19, 2)
	report(1, 20, 1)
	report(1, 21, 2)
	report(1, 22, 2)
	report(1, 23, 2)
	report(1, 24, 2)
	report(1, 25, 1)
	report(1, 26, 2)
	report(1, 27, 2)
	report(1, 28, 2)
	report(1, 29, 2)
	report(1, 30, 2)
	report(1, 31, 2)
	report(1, 32, 1)
	report(1, 33, 2)
	report(1, 34, 2)
	report(1, 35, 2)
	report(1, 36, 2)
	report(1, 37, 2)
	report(1, 38, 2)
	report(1, 39, 2)
	report(1, 40, 1)
	report(1, 41, 2)
	report(1, 42, 2)
	report(1, 43, 2)
	report(1, 44, 2)
	report(1, 45, 2)
	report(1, 46, 2)
	report(1, 47, 2)
	report(1, 48, 2)
	report(1, 49, 2)
	report(1, 50, 1)
	report(1, 51, 2)
	report(1, 52, 2)
	report(1, 53, 2)
	report(1, 54, 2)
	report(1, 55, 2)
	report(1, 56, 2)
	report(1, 57, 2)
	report(1, 58, 2)
	report(1, 59, 2)
	report(1, 60, 2)
	report(1, 61, 2)
	report(1, 62, 2)
	report(1, 63, 2)
	report(1, 64, 1)
	report(1, 65, 2)
	report(1, 66, 2)
	report(1, 67, 2)
	report(1, 68, 2)
	report(1, 69, 2)
	report(1, 70, 2)
	report(1, 71, 2)
	report(1, 72, 2)
	report(1, 73, 2)
	report(1, 74, 2)
	report(1, 75, 2)
	report(1, 76, 2)
	report(1, 77, 2)
	report(1, 78, 2)
	report(1, 79, 2)
	report(1, 80, 1)
	report(1, 81, 2)
	report(1, 82, 2)
	report(1, 83, 2)
	report(1, 84, 2)
	report(1, 85, 2)
	report(1, 86, 2)
	report(1, 87, 2)
	report(1, 88, 2)
	report(1, 89, 2)
	report(1, 90, 2)
	report(1, 91, 2)
	report(1, 92, 2)
	report(1, 93, 2)
	report(1, 94, 2)
	report(1, 95, 2)
	report(1, 96, 2)
	report(1, 97, 2)
	report(1, 98, 2)
	report(1, 99, 2)
}
