package day13_test

import (
	"advent-of-code-22/day13"
	"fmt"
	"testing"
)

func TestCompareIntegersLowerFirst(t *testing.T) {
	result, reason := day13.Compare(1, 2)
	if !result {
		t.Errorf("Compare(1, 2) = %t; want true", result)
	}

	if reason != "Left side is smaller, so inputs are in the right order" {
		t.Errorf("Compare(1, 2) reason = %s; want 'Left side is smaller, so inputs are in the right order'", reason)
	}
}

func TestCompareIntegersAreTheSame(t *testing.T) {
	result, reason := day13.Compare(1, 1)
	if !result {
		t.Errorf("Compare(1, 1) = %t; want true", result)
	}

	if reason != "Integers are the same" {
		t.Errorf("Compare(1, 1) reason = %s; want 'Integers are the same'", reason)
	}
}

func TestCompareIntegersHigherFirst(t *testing.T) {
	result, reason := day13.Compare(2, 1)
	if result {
		t.Errorf("Compare(2, 1) = %t; want false", result)
	}

	if reason != "Left side is larger, so inputs are not in the right order" {
		t.Errorf("Compare(2, 1) reason = %s; want 'Left side is larger, so inputs are not in the right order'", reason)
	}
}

func TestTokenizeStringSimple(t *testing.T) {
	result := day13.Tokenize("[1,1,3,1,1]")

	want := day13.PacketList{
		float64(1),
		float64(1),
		float64(3),
		float64(1),
		float64(1),
	}

	for i, token := range want {
		if result[i] != token {
			t.Errorf("Tokenize(\"[1,1,3,1,1]\") = %v; want PacketList{1,1,3,1,1}", result)
		}
	}
}

func TestTokenizeStringComplex(t *testing.T) {
	result := day13.Tokenize("[[1],[2,3,4]]")

	want := day13.PacketList{
		day13.PacketList{float64(1)},
		day13.PacketList{
			float64(2),
			float64(3),
			float64(4),
		},
	}

	if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", want) {
		t.Errorf("Tokenize(\"[[1],[2,3,4]]\") = %v; want PacketList{PacketList{1}, PacketList{2, 3, 4}}", result)
	}
}

func TestTokenizeStringDoubleDigit(t *testing.T) {
	result := day13.Tokenize("[10]")

	want := day13.PacketList{float64(10)}

	if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", want) {
		t.Errorf("Tokenize(\"[10]\") = %v; want PacketList{10}", result)
	}
}
