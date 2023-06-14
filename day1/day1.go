package day1

import (
	"advent-of-code-22/lib"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Run() {
	calories := loadCalories("./data/calories.txt")
	totals := getCalorieTotals(calories)

	fmt.Printf("Max Calories: %d\n", getMaxCalories(totals))

	topThree := getTopCalories(totals, 3)

	fmt.Printf("Top Three: %+v\nTotal: %d\n", topThree, sum(topThree))
}

func loadCalories(filename string) []int {
	file, err := os.Open(filename)
	lib.HandleError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	calories := []int{}
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) > 0 {
			i, err := strconv.Atoi(s)
			lib.HandleError(err)
			calories = append(calories, i)
		} else {
			calories = append(calories, -1)
		}
	}

	lib.HandleError(scanner.Err())

	return calories
}

func getCalorieTotals(calories []int) []int {
	totals := []int{}
	elf := 0

	for _, cal := range calories {
		if cal > 0 {
			elf += cal
		} else {
			totals = append(totals, elf)
			elf = 0
		}
	}

	return totals
}

func getMaxCalories(calories []int) int {
	maxCalories := 0

	for _, cal := range calories {
		maxCalories = max(cal, maxCalories)
		fmt.Printf("cal: %d max: %d\n", cal, maxCalories)
	}

	return maxCalories
}

func getMaxIndex(calories []int) int {
	maxCalories, index := 0, 0

	for i, cal := range calories {
		if cal > maxCalories {
			maxCalories = cal
			index = i
		}

	}

	return index
}

func getTopCalories(calories []int, num int) []int {
	top := []int{}

	for i := 0; i < num; i++ {
		index := getMaxIndex(calories)
		top = append(top, calories[index])
		calories = remove(calories, index)
	}

	return top
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func sum(s []int) int {
	total := 0
	for _, v := range s {
		total += v
	}

	return total
}
