package day13

import (
	"fmt"
)

func Run() {
	fmt.Println("[1,1,3,1,1]", Tokenize("[1,1,3,1,1]"))
	fmt.Println("[[1],[2,3,4]]", Tokenize("[[1],[2,3,4]]"))
}

func Compare(left, right int) (bool, string) {
	fmt.Printf("- Compare %d vs %d\n", left, right)
	if left < right {
		return true, "Left side is smaller, so inputs are in the right order"
	}

	if left == right {
		return true, "Integers are the same"
	}

	return false, "Left side is larger, so inputs are not in the right order"
}

func Tokenize(str string) []string {
	tokens := []string{}
	// [1,1,3,1,1]
	// [[1],[2,3,4]]

	var token string
	for _, s := range str {
		if s == '[' {
			tokens = append(tokens, string(s))
		} else if s == ',' {
			if len(token) > 0 {
				tokens = append(tokens, token)
				token = ""
			}
		} else if s == ']' {
			if len(token) > 0 {
				tokens = append(tokens, token)
				token = ""
			}

			tokens = append(tokens, string(s))
		} else {
			token += string(s)
		}
	}

	return tokens
}
