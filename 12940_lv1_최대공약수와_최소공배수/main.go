package main

import (
	"fmt"
	"time"
)

func solution(n int, m int) []int {
	if n > m {
		n, m = m, n
	}

	// 최대 공약수
	a := 1
	for i := n; i > 0; i-- {
		if n%i == 0 {
			if m%i == 0 {
				a = i
				break
			}
		}
	}

	// 최소 공배수
	b := n * m / a

	return []int{a, b}
}

func report(n int, m int, desireResult []int) {

	input := fmt.Sprintf("Input:[%v %v]", n, m)

	start := time.Now()

	result := solution(n, m)

	duration := time.Since(start)

	isSame := true
	if len(result) != len(desireResult) {
		isSame = false
	} else {
		for i, v := range result {
			if desireResult[i] != v {
				isSame = false
				break
			}
		}
	}

	ox := ""
	if isSame {
		ox = fmt.Sprintf("[O]")
	} else {
		ox = fmt.Sprintf("[X]")
	}

	fmt.Printf("%s %s result:[%v] duration:%v\n", ox, input, result, duration)
}

func main() {
	report(3, 12, []int{3, 12})
	report(2, 5, []int{1, 10})
}
