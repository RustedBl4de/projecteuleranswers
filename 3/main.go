/*
Answers to projecteuler.net

#3 : https://projecteuler.net/problem=3

Largest prime factor
Problem 3

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143?
*/

package main

import (
	"fmt"
	"time"
)

const myNumber int = 600851475143

var answer int

func isPrime(number int) bool {
	var result bool

	if number <= 1 {
		result = false
	} else if (number % 2) == 0 {
		result = false
	} else if len(getFactors(number, true)) > 0 {
		result = false
	} else {
		result = true
	}

	return result
}

func getFactors(number int, quick bool) []int {
	var results []int

	for i := 2; i < number; i++ {
		if (number % i) == 0 {
			if number == myNumber {
				if isPrime((number / i)) {
					answer = (number / i)
					break
				}
			}
			results = append(results, i)
		}

		if quick && len(results) > 1 {
			break
		}
	}

	return results
}

func main() {
	start := time.Now()

	getFactors(myNumber, false)
	fmt.Println("The answer is:", answer)

	duration := time.Since(start)

	fmt.Println("Execution Time:", duration)
}
