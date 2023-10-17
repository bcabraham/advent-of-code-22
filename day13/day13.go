package day13

func Compare(left, right int) (bool, string) {
	if left < right {
		return true, "Left side is smaller, so inputs are in the right order"
	}

	if left == right {
		return true, "Integers are the same"
	}

	return false, "Left side is larger, so inputs are not in the right order"
}
