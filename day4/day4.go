package day4

import (
	"advent-of-code-22/lib"
	"fmt"
	"strings"
)

type CleaningOrder struct {
	Start int
	End   int
}

type Comparable func(a, b CleaningOrder) bool

type CleaningOrderList [][]CleaningOrder

func Run() {
	cleaningOrders, err := lib.ReadLines("camp-cleanup.txt")
	lib.HandleError(err)

	cleaningOrdersList := parseCleaningOrderList(cleaningOrders)

	var overlaps CleaningOrderList
	var noOverlaps CleaningOrderList

	overlaps, noOverlaps = findOverlaps(cleaningOrdersList, hasCompleteOverlap)
	display("hasCompleteOverlap", overlaps)

	overlaps, noOverlaps = findOverlaps(noOverlaps, hasPartialOverlap)
	display("hasPartialOverlap", overlaps)
	display("noOverlaps", noOverlaps)
}

func display(message string, cleaningOrders CleaningOrderList) {
	fmt.Println(message)
	for _, o := range cleaningOrders {
		fmt.Printf("%+v, %+v\n", o[0], o[1])
	}

	fmt.Printf("Total: %d\n", len(cleaningOrders))
}

func findOverlaps(cleaningOrders CleaningOrderList, fn Comparable) (CleaningOrderList, CleaningOrderList) {
	overlaps := CleaningOrderList{}
	noOverlaps := CleaningOrderList{}

	for _, o := range cleaningOrders {
		a, b := o[0], o[1]

		if fn(a, b) {
			overlaps = append(overlaps, []CleaningOrder{a, b})
		} else {
			noOverlaps = append(noOverlaps, []CleaningOrder{a, b})
		}
	}

	return overlaps, noOverlaps
}

func parseCleaningOrderList(cleaningOrders []string) CleaningOrderList {
	cleaningOrdersList := CleaningOrderList{}
	for _, o := range cleaningOrders {
		a, b := parseCleaningOrders(o)
		cleaningOrdersList = append(cleaningOrdersList, []CleaningOrder{a, b})
	}

	return cleaningOrdersList
}

func parseCleaningOrders(str string) (CleaningOrder, CleaningOrder) {
	orders := strings.SplitN(str, ",", 2)
	if len(orders) != 2 {
		panic("error parsing CleaningOrder")
	}

	order1 := parseCleaningOrder(orders[0])
	order2 := parseCleaningOrder(orders[1])

	return order1, order2
}

func parseCleaningOrder(str string) CleaningOrder {
	order := strings.SplitN(str, "-", 2)

	if len(order) != 2 {
		panic("error parsing CleaningOrder")
	}

	return CleaningOrder{
		lib.StrToInt(order[0]),
		lib.StrToInt(order[1]),
	}
}

// In how many assignment pairs does one range fully contain the other?
func hasCompleteOverlap(a, b CleaningOrder) bool {
	return a.Start >= b.Start && a.End <= b.End || b.Start >= a.Start && b.End <= a.End
}

// In how many assignment pairs does one range overlap at all with the other?
func hasPartialOverlap(a, b CleaningOrder) bool {
	return a.End >= b.Start && a.End <= b.End || b.End >= a.Start && b.End <= a.End
}
