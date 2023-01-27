package main

import (
	"fmt"
	"time"
)

func solution(dots [][]int) int {
	maxX, minX := dots[0][0], dots[0][0]
	maxY, minY := dots[0][1], dots[0][1]

	for _, dot := range dots {
		if maxX < dot[0] {
			maxX = dot[0]
		}
		if minX > dot[0] {
			minX = dot[0]
		}
		if maxY < dot[1] {
			maxY = dot[1]
		}
		if minY > dot[1] {
			minY = dot[1]
		}
	}
	return (maxX - minX) * (maxY - minY)
}

func report(dots [][]int, desireResult int) {

	input := fmt.Sprintf("Input:[%v]", dots)

	start := time.Now()

	result := solution(dots)

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
	report([][]int{{1, 1}, {2, 1}, {2, 2}, {1, 2}}, 1)
	report([][]int{{-1, -1}, {1, 1}, {1, -1}, {-1, 1}}, 4)
}
