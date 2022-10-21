/*
Answers to projecteuler.net

#2 : https://projecteuler.net/problem=2

Even Fibonacci numbers
Problem 2

*/

package main

import (
	"fmt"
	"time"
)

func isEven(number int) bool {
	var result bool

	if number%2 == 0 {
		result = true
	} else {
		result = false
	}

	return result
}

func main() {
	var upperBound int = 4000000
	var fibSeq []int
	var fibStart, fibPrev, fibCurr, fibNext int

	start := time.Now()

	fibStart = 0
	fibSeq = append(fibSeq, fibStart)

	fibPrev = 0
	fibCurr = 1
	fibSeq = append(fibSeq, fibCurr)

	fibNext = 1

	for fibNext < upperBound {
		if isEven(fibNext) {
			fibSeq = append(fibSeq, fibNext)
		}

		fibPrev = fibCurr

		fibCurr = fibNext

		fibNext = fibPrev + fibCurr
	}

	fmt.Println(fibSeq)

	duration := time.Since(start)

	fmt.Println("Execution Time:", duration)
}
