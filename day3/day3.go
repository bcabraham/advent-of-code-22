package day3

import (
	"advent-of-code-22/lib"
	"fmt"
	"strings"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var charMap = map[string]int{}

func Run() {
	for i, c := range chars {
		charMap[string(c)] = i + 1
	}

	rucksacks, err := lib.ReadLines("day3", "rucksacks.txt")
	lib.HandleError(err)
	compareRucksacks(rucksacks)
	findBadges(rucksacks)
}

// Find the item type that corresponds to the badges of each three-Elf group.
// What is the sum of the priorities of those item types?
func findBadges(rucksacks []string) {
	sacks := []string{}
	count := 0
	priorityTotal := 0
	for i, r := range rucksacks {
		sacks = append(sacks, r)
		count += 1

		if count == 3 {
			badge := findBadge(sacks)
			priority := getPriority(badge)
			priorityTotal += priority
			fmt.Printf("%d %v: %s (%d) = %d\n", i, sacks, badge, priority, priorityTotal)

			sacks = nil
			count = 0
		}
	}

	fmt.Printf("Total Priority: %d\n", priorityTotal)
}

func findBadge(rucksacks []string) string {
	if len(rucksacks) < 3 {
		return ""
	}

	matches := findMatches(rucksacks[0], rucksacks[1])
	match := findMatches(matches, rucksacks[2])

	return string(match[0])
}

func findMatches(str1, str2 string) string {
	matches := ""
	for _, c := range str1 {
		if strings.Contains(str2, string(c)) {
			matches += string(c)
		}
	}

	return matches
}

func compareRucksacks(rucksacks []string) {
	total := 0
	for _, str := range rucksacks {
		left, right := splitLine(str)
		match := findMatch(left, right)
		priority := getPriority(match)

		total += priority
		fmt.Printf("%s: %s %s. Both contain: %s (%d)\n", str, left, right, match, priority)
	}

	fmt.Printf("Total: %d\n", total)
}

func splitLine(str string) (string, string) {
	i := len(str) / 2
	return str[:i], str[i:]
}

func findMatch(str1, str2 string) string {
	for _, c := range str1 {
		if strings.Contains(str2, string(c)) {
			return string(c)
		}
	}

	return ""
}

func getPriority(c string) int {
	return charMap[c]
}
