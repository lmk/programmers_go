package main

import (
	"fmt"
	"strings"
	"time"
)

func solution(id_list []string, report []string, k int) []int {
	// 동일한 유저 신고 회수는 1회로 처리
	mr := map[string]int{}
	for _, v := range report {
		mr[v] = 1
	}
	report = []string{}
	for k, _ := range mr {
		report = append(report, k)
	}

	// 신고 당한 id 회수 계산
	mr = map[string]int{} // 신고당한ID: 회수
	for _, v := range report {
		r := strings.Split(v, " ")
		_, exists := mr[r[1]]
		if !exists {
			mr[r[1]] = 1
		} else {
			mr[r[1]] = mr[r[1]] + 1
		}
	}

	// 메일 발송 계산
	mail := map[string]int{} // 사용자ID: 메일발송건수
	for _, v := range report {
		r := strings.Split(v, " ")

		for id, n := range mr {

			if n < k {
				continue
			}

			if r[1] == id {
				_, exists := mail[r[0]]
				if !exists {
					mail[r[0]] = 1
				} else {
					mail[r[0]] = mail[r[0]] + 1
				}
			}
		}
	}

	result := []int{}
	for _, id := range id_list {
		result = append(result, mail[id])
	}

	return result
}

func report(id_list []string, report []string, k int, desireResult []int) {

	input := fmt.Sprintf("Input:[%v, %v, %v]", id_list, report, k)

	start := time.Now()

	result := solution(id_list, report, k)

	duration := time.Since(start)

	isSame := true
	if len(desireResult) != len(result) {
		isSame = false
	} else {
		for i, _ := range desireResult {
			if desireResult[i] != result[i] {
				isSame = false
				break
			}

		}
	}
	ox := ""
	if isSame == true {
		ox = fmt.Sprintf("[O]")
	} else {
		ox = fmt.Sprintf("[X]")
	}

	fmt.Printf("%s %s result:[%v] duration:%v\n", ox, input, result, duration)
}

func main() {
	report([]string{"muzi", "frodo", "apeach", "neo"}, []string{"muzi frodo", "apeach frodo", "frodo neo", "muzi neo", "apeach muzi"}, 2, []int{2, 1, 1, 0})
	report([]string{"con", "ryan"}, []string{"ryan con", "ryan con", "ryan con", "ryan con"}, 3, []int{0, 0})
}
