package day13

import (
	"encoding/json"
	"fmt"
)

type Packet interface{}
type PacketList []Packet

func Run() {
	left := "[1,1,3,1,1]"
	right := "[1,1,5,1,1]"

	fmt.Println(left, Tokenize(left))
	fmt.Println(right, Tokenize(right))
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

func Tokenize(str string) PacketList {
	tokens := PacketList{}
	json.Unmarshal([]byte(str), &tokens)

	return tokens
}
