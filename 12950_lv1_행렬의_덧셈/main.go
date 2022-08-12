package main

import (
	"fmt"
	"time"
)

func solution(arr1 [][]int, arr2 [][]int) [][]int {
	a := make([][]int, len(arr1))
	for i, aa := range arr1 {
		a[i] = make([]int, len(aa))
		for j, _ := range aa {
			a[i][j] = arr1[i][j] + arr2[i][j]
		}
	}

	return a
}

func report(arr1 [][]int, arr2 [][]int, desireResult [][]int) {

	input := fmt.Sprintf("Input:[%v %v]", arr1, arr2)

	start := time.Now()

	result := solution(arr1, arr2)

	duration := time.Since(start)

	isSame := true
	if len(desireResult) == len(result) {
		for i, a := range desireResult {
			for j, _ := range a {
				if desireResult[i][j] != result[i][j] {
					isSame = false
					break
				}
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

	report([][]int{{1, 2}, {2, 3}}, [][]int{{3, 4}, {5, 6}}, [][]int{{4, 6}, {7, 9}})
	report([][]int{{1}, {2}}, [][]int{{3}, {4}}, [][]int{{4}, {6}})

}
