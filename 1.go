package main

import (
	"fmt"
	"time"
)

func getMultipliers(multiple int, upperBound int) []int {
	var multipliers []int
	var result int

	for i := 1; i <= upperBound; i++ {
		result = i * multiple

		if result <= upperBound {
			multipliers = append(multipliers, result)
		}
	}

	return multipliers
}

func getSum(addends []int) int {
	var sum int

	for _, n := range addends {
		sum += n
	}

	return sum
}

func main() {
	start := time.Now()

	fmt.Println("The sum of all multipliers of 3 less than 1000:", getSum(getMultipliers(3, 1000)))
	fmt.Println("The sum of all multipliers of 5 less than 1000:", getSum(getMultipliers(5, 1000)))

	duration := time.Since(start)

	fmt.Println("Execution Time:", duration)
}
