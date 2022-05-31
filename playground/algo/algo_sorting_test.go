package algo

import (
	"fmt"
	"testing"
)

func bubble(input []int) []int { // O(n^2)
	for i := 0; i < len(input); i++ { // O(n)
		for j := 0; j < len(input); j++ { // O(n)
			if input[i] > input[j] {
				input[i], input[j] = input[j], input[i]
			}
		}
	}
	return input
}

func TestBubble(t *testing.T) {
	data := []int{5, 1, 8, 9, 4, 6, 3, 7, 2}
	fmt.Println(bubble(data))
}
