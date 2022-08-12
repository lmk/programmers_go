package main

import (
	"fmt"
	"time"
)

func limit(lo, hi, n int) int {
	if n > hi {
		return hi
	}

	if n < lo {
		return lo
	}

	return n
}

func solution(lottos []int, win_nums []int) []int {

	// 0의 개수를 세고
	zero := 0
	for _, n := range lottos {
		if n == 0 {
			zero++
		}
	}

	// 남은 숫자의 빙고 개수를 세고
	bingo := 0
	for _, l := range lottos {
		for _, w := range win_nums {
			if l == w {
				bingo++
			}
		}
	}

	// hi = 빙고수 + 0수
	hi := limit(1, 6, bingo+zero)

	// lo = 빙고수
	lo := bingo

	//등수 계산
	hi = limit(1, 6, 6-hi+1)
	lo = limit(1, 6, 6-lo+1)

	return []int{hi, lo}
}

func report(lottos []int, win_nums []int, desireResult []int) {

	input := fmt.Sprintf("Input:[%v %v]", lottos, win_nums)

	start := time.Now()

	result := solution(lottos, win_nums)

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

	report([]int{44, 1, 0, 0, 31, 25}, []int{31, 10, 45, 1, 6, 19}, []int{3, 5})
	report([]int{0, 0, 0, 0, 0, 0}, []int{38, 19, 20, 40, 15, 25}, []int{1, 6})
	report([]int{45, 4, 35, 20, 3, 9}, []int{20, 9, 3, 45, 4, 35}, []int{1, 1})
	report([]int{1, 2, 3, 4, 5, 6}, []int{7, 8, 9, 10, 11, 12}, []int{6, 6})
}
