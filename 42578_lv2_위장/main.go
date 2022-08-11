package main

import (
	"fmt"
	"time"
)

func solution(clothes [][]string) int {
	m := map[string]int{} // 종류: 의상 개수
	for _, v := range clothes {
		m[v[1]] += 1
	}

	// // 한 종류인 경우
	// if len(m) == 1 {
	// 	return len(clothes)
	// }

	result := 1
	for _, v := range m {
		result *= (v + 1) // 종류가 빠지는 경우
	}

	return result - 1 // 안입는 경우 제외
}

func report(numbers [][]string, desireResult int) {

	input := fmt.Sprintf("Input:[%v]", numbers)

	start := time.Now()

	result := solution(numbers)

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

	report([][]string{{"yellow_hat", "headgear"}, {"blue_sunglasses", "eyewear"}, {"green_turban", "headgear"}}, 5)
	report([][]string{{"crow_mask", "face"}, {"blue_sunglasses", "face"}, {"smoky_makeup", "face"}}, 3)
	report([][]string{{"a", "a"}, {"b", "a"}, {"c", "b"}, {"d", "b"}, {"e", "c"}, {"f", "c"}}, 26)
}
