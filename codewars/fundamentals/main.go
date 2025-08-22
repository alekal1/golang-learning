package main

import (
	"aleksale/fundamentals/tribonacci"
)

// Main func to test output from submodules under fundamentals package
func main() {
	tribonacci.Tribonacci([3]float64{1, 1, 1}, 10)
}
