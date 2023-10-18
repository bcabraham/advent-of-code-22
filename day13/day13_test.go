package day13_test

import (
	"advent-of-code-22/day13"
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

	want := []string{"[", "1", "1", "3", "1", "1", "]"}

	for i, token := range want {
		if result[i] != token {
			t.Errorf("Tokenize(\"[1,1,3,1,1]\") = %s; want []string{'[','1','1','3','1','1',']'}", result)
		}
	}
}

func TestTokenizeStringComplex(t *testing.T) {
	result := day13.Tokenize("[[1],[2,3,4]]")

	want := []string{"[", "[", "1", "]", "[", "2", "3", "4", "]", "]"}

	for i, token := range want {
		if result[i] != token {
			t.Errorf("Tokenize(\"[[1],[2,3,4]]\") = %s; want []string{'[','[','1',']','[','2','3','4',']',']'}", result)
		}
	}
}

func TestTokenizeStringDoubleDigit(t *testing.T) {
	result := day13.Tokenize("[10]")

	want := []string{"[", "10", "]"}

	for i, token := range want {
		if result[i] != token {
			t.Errorf("Tokenize(\"[10]\") = %s; want []string{'[','10',']'}", result)
		}
	}

}

func TestCompareListEmpty(t *testing.T) {
	empty := []string{"[", "]"}
	result, reason := day13.CompareLists(empty, empty)

	if !result {
		t.Errorf("CompareLists([], ][]) = %t; want true", result)
	}

	if reason != "Lists are the same length" {
		t.Errorf("CompareLists([], ][]) reason = %s; want 'Lists are the same length'", reason)
	}
}
