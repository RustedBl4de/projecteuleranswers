package main

/*
Answers to projecteuler.net

#1 : https://projecteuler.net/problem=1

Multiples of 3 or 5
Problem 1
*/

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

	var results3, results5 int
	var results []int

	results3 = getSum(getMultipliers(3, 1000))
	results5 = getSum(getMultipliers(5, 1000))

	results = append(results, results3)
	results = append(results, results5)

	fmt.Println("The sum of all multipliers of 3 less than 1000:", results3)
	fmt.Println("The sum of all multipliers of 5 less than 1000:", results5)
	fmt.Println("The sum of all multipliers of 3 or 5 less than 1000:", getSum(results))

	duration := time.Since(start)

	fmt.Println("Execution Time:", duration)
}
