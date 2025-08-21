package smallestInt

import "math"

func SmallestIntegerFinder(numbers []int) int {
	currMin := math.MaxInt
	for _, v := range numbers {
		if v < currMin {
			currMin = v
		}
	}
	return currMin
}
