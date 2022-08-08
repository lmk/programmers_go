package main

import (
	"fmt"
	"time"
)

func solution(array []int, commands [][]int) []int {
	result := make([]int, len(commands))

	for r, cmd := range commands {
		i := cmd[0]
		j := cmd[1]
		k := cmd[2]

		ar := make([]int, j-i+1)
		copy(ar, array[i-1:j])

		// 정렬
		for a := 0; a < len(ar); a++ {
			for b := a + 1; b < len(ar); b++ {
				if ar[a] > ar[b] {
					ar[a], ar[b] = ar[b], ar[a]
				}
			}
		}

		result[r] = ar[k-1]
	}

	return result
}

func report(array []int, commands [][]int, desireResult []int) {

	start := time.Now()

	result := solution(array, commands)

	duration := time.Since(start)

	isSame := true
	if len(result) != len(desireResult) {
		isSame = false
	} else {
		for i, _ := range desireResult {
			if result[i] != desireResult[i] {
				isSame = false
				break
			}
		}
	}

	if isSame {
		fmt.Printf("[O] ")
	} else {
		fmt.Printf("[X] ")
	}

	fmt.Printf("Input:[%v, %v], result:[%v], duration:%v\n", array, commands, result, duration)
}

func main() {

	report([]int{1, 5, 2, 6, 3, 7, 4}, [][]int{{2, 5, 3}, {4, 4, 1}, {1, 7, 3}}, []int{5, 6, 3})
	report([]int{1}, [][]int{{1, 1, 1}}, []int{1})
	report([]int{1, 5, 2, 6, 3, 7, 4, 1, 5, 2, 6, 3, 7, 4, 1, 5, 2, 6, 3, 7, 4}, [][]int{{2, 5, 3}, {4, 4, 1}, {1, 7, 3}, {2, 5, 3}, {4, 4, 1}, {1, 7, 3}}, []int{5, 6, 3, 5, 6, 3})
	report([]int{1, 5, 2, 6, 3, 7, 4}, [][]int{{5, 2, 3}, {4, 4, 1}, {7, 1, 3}}, []int{5, 6, 3})
}
