package tribonacci

func Tribonacci(signature [3]float64, n int) []float64 {
	if n <= 0 {
		return []float64{}
	}

	res := make([]float64, 0, n)

	sequence := tribonacci(signature)

	for i := 0; i < n; i++ {
		res = append(res, sequence())
	}

	return res
}

func tribonacci(signature [3]float64) func() float64 {
	first, second, third := signature[0], signature[1], signature[2]
	return func() float64 {
		ret := first
		first, second, third = second, third, first+second+third
		return float64(ret)
	}
}
