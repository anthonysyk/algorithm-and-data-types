package algo

import (
	"fmt"
	"testing"
)

func factorialRecursive(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorialIterative(n-1)
}

func factorialIterative(n int) int {
	if n < 2 {
		return 1
	}
	acc := n
	for i := n - 1; i > 1; i-- {
		acc = acc * i
	}
	return acc
}

func TestFactorial(t *testing.T) {
	n := 5
	fmt.Println(factorialIterative(n))
	fmt.Println(factorialRecursive(n))
}
