package main

import (
	"fmt"
	"time"
)

func solution(x int, n int) []int64 {

	a := []int64{}
	for i := int64(x); true; i += int64(x) {
		if len(a) >= n {
			break
		}
		a = append(a, i)
	}

	return a
}

func report(x int, n int, desireResult []int64) {

	input := fmt.Sprintf("Input:[%v %v]", x, n)

	start := time.Now()

	result := solution(x, n)

	duration := time.Since(start)

	isSame := true
	if len(desireResult) == len(result) {
		for i, _ := range desireResult {
			if desireResult[i] != result[i] {
				isSame = false
				break
			}
		}
	} else {
		isSame = false
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

	report(2, 5, []int64{2, 4, 6, 8, 10})
	report(4, 3, []int64{4, 8, 12})
	report(-4, 2, []int64{-4, -8})
}
