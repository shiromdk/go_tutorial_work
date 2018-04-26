package main

import (
	"fmt"
	"time"
)

var chain map[int64]int64

func main() {
	start := time.Now()
	chain = make(map[int64]int64)
	var max int64 = 0
	var maxIndex int = 0
	for i := 1; i < 1000000; i++ {
		var temp int64 = collatz(int64(i))
		if temp > max {
			max = temp
			maxIndex = i
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("It took %s to find longest chain %d with length of %d", elapsed, maxIndex, max)
	fmt.Println()
}

func collatz(num int64) int64 {
	var nextNum int64
	if chain[num] != 0 {
		return chain[num]
	}
	if num == 1 {
		return 1
	}
	if num%2 == 0 {
		nextNum = num / 2
	} else {
		nextNum = (3 * num) + 1
	}
	chain[num] = 1 + collatz(nextNum)
	return chain[num]

}
