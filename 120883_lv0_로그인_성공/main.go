package main

import (
	"fmt"
	"time"
)

func solution(id_pw []string, db [][]string) string {

	for _, el := range db {
		if id_pw[0] == el[0] {
			if id_pw[1] == el[1] {
				return "login"
			} else {
				return "wrong pw"
			}
		}
	}

	return "fail"
}
func report(id_pw []string, db [][]string, desireResult string) {

	input := fmt.Sprintf("Input:[%v,%v]", id_pw, db)

	start := time.Now()

	result := solution(id_pw, db)

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
	report([]string{"meosseugi", "1234"}, [][]string{{"rardss", "123"}, {"yyoom", "1234"}, {"meosseugi", "1234"}}, "login")
	report([]string{"programmer01", "15789"}, [][]string{{"programmer02", "111111"}, {"programmer00", "134"}, {"programmer01", "1145"}}, "wrong pw")
	report([]string{"rabbit04", "98761"}, [][]string{{"jaja11", "98761"}, {"krong0313", "29440"}, {"rabbit00", "111333"}}, "fail")
}
