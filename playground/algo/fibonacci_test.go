package algo

import (
	"fmt"
	"testing"
)

func fibIterative(n int) int {
	x, y := 0, 1
	for i := 2; i < n+1; i++ {
		x, y = y, x+y
	}

	return y
}
func fibRecursive(n int) int {
	if n < 2 {
		return n
	}
	return fibRecursive(n-1) + fibRecursive(n-2)
}

func TestFibonacci(t *testing.T) {
	fmt.Println(fibIterative(10))
	fmt.Println(fibRecursive(10))
}
