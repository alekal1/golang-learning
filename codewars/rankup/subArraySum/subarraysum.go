package subarraysum

func MaximumSubarraySum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	// return iteration(numbers)
	return kadane(numbers)
}

//////////////////////////////////////////////////////////////
/*					Kadane's algorithm						*/
//////////////////////////////////////////////////////////////

func kadane(numbers []int) int {
	maxSum := 0
	currentSum := 0

	for _, v := range numbers {
		currentSum = kadaneMax(0, currentSum+v)
		maxSum = kadaneMax(maxSum, currentSum)
	}

	return maxSum
}

func kadaneMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//////////////////////////////////////////////////////////////
/*					Iteration method						*/
//////////////////////////////////////////////////////////////

func iteration(numbers []int) int {
	currMax := 0

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			subArraySum := calculateSubArraySum(numbers[i:j])
			if subArraySum > currMax {
				currMax = subArraySum
			}
		}
	}

	return currMax
}

func calculateSubArraySum(subArray []int) int {
	var res int
	for _, v := range subArray {
		res += v
	}
	return res
}
