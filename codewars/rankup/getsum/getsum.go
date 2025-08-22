package getsum

/**
Get sum between two numbers
*/
func GetSum(a, b int) int {
	if a == b {
		return a
	}

	var res int

	start, end := getStartAndEnd(a, b)
	for {
		res += start

		if start < end {
			start++
			continue
		}

		if start == end {
			break
		}
	}

	return res
}

func getStartAndEnd(a, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}
