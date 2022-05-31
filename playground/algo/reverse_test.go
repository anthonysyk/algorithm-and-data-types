package algo

import (
	"fmt"
	"testing"
)

func reverseRecursive(input []int, result []int) []int {
	if len(input) == 0 {
		return result
	}
	result = append(result, input[0:1]...)
	return reverseRecursive(input[1:], result)
}

func reverseIterative(input []int) []int {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}

	return input
}

func TestReverse(t *testing.T) {
	test := []int{1, 2, 3, 4, 5, 7, 8, 9, 10}
	fmt.Println(reverseIterative(test))
	fmt.Println(reverseRecursive(test, nil))
}
