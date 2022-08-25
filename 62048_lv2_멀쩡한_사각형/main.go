package main

import (
	"fmt"
	"time"
)

func solution(w int, h int) int64 {

	if w > h {
		w, h = h, w
	}

	// 최대 공약수
	a := 1
	for i := w; i > 0; i-- {
		if w%i == 0 {
			if h%i == 0 {
				a = i
				break
			}
		}
	}

	return int64(w*h) - int64(w+h-a)
}

func report(w int, h int, desireResult int64) {

	input := fmt.Sprintf("Input:[%v, %v]", w, h)

	start := time.Now()

	result := solution(w, h)

	duration := time.Since(start)

	isSame := true
	if result != desireResult {
		isSame = false
	}

	if isSame == true {
		fmt.Printf("[O] %s result:[%v] duration:%v\n", input, result, duration)
	} else {
		fmt.Printf("[X] %s result:[%v]->[%v] duration:%v\n", input, result, desireResult, duration)
	}
}

func main() {
	report(8, 12, 80)
	report(8, 8, 56)
	report(2, 3, 2)
	report(2, 4, 4)
	report(3, 3, 6)
	report(1, 1, 0)
	report(1, 2, 0)
}
