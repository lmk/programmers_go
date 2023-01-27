package main

import (
	"fmt"
	"time"
)

func solution(keyinput []string, board []int) []int {

	lx, ly := board[0]/2, board[1]/2
	x, y := 0, 0

	for _, k := range keyinput {
		switch k {
		case "right":
			x++
			if x > lx {
				x = lx
			}
		case "left":
			x--
			if x < (lx * -1) {
				x = (lx * -1)
			}
		case "up":
			y++
			if y > ly {
				y = ly
			}
		case "down":
			y--
			if y < (ly * -1) {
				y = (ly * -1)
			}
		}
	}

	return []int{x, y}
}

func report(keyinput []string, board []int, desireResult []int) {

	input := fmt.Sprintf("Input:[%v,%v]", keyinput, board)

	start := time.Now()

	result := solution(keyinput, board)

	duration := time.Since(start)

	isSame := true
	if len(result) != len(desireResult) {
		isSame = false
	} else {
		for i := range desireResult {
			if result[i] != desireResult[i] {
				isSame = false
				break
			}
		}
	}

	if isSame {
		fmt.Printf("[O] %s result:[%v] duration:%v\n", input, result, duration)
	} else {
		fmt.Printf("[X] %s result:[%v]->[%v] duration:%v\n", input, result, desireResult, duration)
	}
}

func main() {
	report([]string{"left", "right", "up", "right", "right"}, []int{11, 11}, []int{2, 1})
	report([]string{"down", "down", "down", "down", "down"}, []int{7, 9}, []int{0, -4})
}
