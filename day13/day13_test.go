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
