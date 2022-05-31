package algo

import (
	"fmt"
	"testing"
)

func fibonacciDP(n int, cache map[int]int) int {
	if v, ok := cache[n]; ok {
		return v
	}
	if n < 2 {
		return n
	}
	result := fibonacciDP(n-1, cache) + fibonacciDP(n-2, cache)
	cache[n] = result
	return cache[n]
}

func TestFibonacciDP(t *testing.T) {
	cache := make(map[int]int)
	fmt.Println(fibonacciDP(100, cache))
}
