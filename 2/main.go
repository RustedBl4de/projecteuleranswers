/*
Answers to projecteuler.net

#2 : https://projecteuler.net/problem=2

Even Fibonacci numbers
Problem 2

Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
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
	start := time.Now()

	var upperBound int = 4000000
	var fibSeq []int
	var fibStart, fibPrev, fibCurr, fibNext int

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
