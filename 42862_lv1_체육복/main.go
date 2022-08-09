package main

import (
	"fmt"
	"sort"
	"time"
)

func removeByIndex(arr []int, i int) []int {
	return append(arr[:i], arr[i+1:]...)
}

func remove(arr []int, n int) []int {
	for i, a := range arr {
		if a == n {
			return removeByIndex(arr, i)
		}
	}

	return arr
}

// contains 존재하면 index를 반환, 없으면 -1
func contains(arr []int, n int) int {
	for i, a := range arr {
		if a == n {
			return i
		}
	}

	return -1
}

func solution(n int, lost []int, reserve []int) int {

	// 여벌이 있는데, 한벌을 도난당한 경우 lost, reserve 모두 삭제
	delete := []int{}
	for _, lo := range lost {
		for _, re := range reserve {
			if lo == re {
				delete = append(delete, lo)
			}
		}
	}

	for _, l := range delete {
		lost = remove(lost, l)
		reserve = remove(reserve, l)
	}

	sort.Ints(lost)
	sort.Ints(reserve)

	lostCount := len(lost)
	for _, l := range lost {
		i := contains(reserve, l-1)
		if i != -1 {
			reserve = removeByIndex(reserve, i)
			lostCount--
		} else {
			i = contains(reserve, l+1)
			if i != -1 {
				reserve = removeByIndex(reserve, i)
				lostCount--
			}
		}
	}

	return n - lostCount
}

func report(n int, lost []int, reserve []int, desireResult int) {

	input := fmt.Sprintf("Input:[%v, %v, %v]", n, lost, reserve)

	start := time.Now()

	result := solution(n, lost, reserve)

	duration := time.Since(start)

	ox := ""
	if result == desireResult {
		ox = fmt.Sprintf("[O]")
	} else {
		ox = fmt.Sprintf("[X]")
	}

	fmt.Printf("%s %s result:[%v] duration:%v\n", ox, input, result, duration)
}

func main() {
	report(5, []int{1, 3, 5}, []int{1, 3, 5}, 5)
	report(5, []int{2, 4}, []int{1, 3, 5}, 5)
	report(5, []int{2, 4}, []int{3}, 4)
	report(3, []int{3}, []int{1}, 2)
}
