package golang

import (
	"fmt"
	"golang.org/x/tour/tree"
	"testing"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		ch <- t.Value
		Walk(t.Left, ch)
	}
	if t.Right != nil {
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {

	return false
}

func Walker(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	//close the channel to avoid panic
	close(ch)
}

func TestBT(t *testing.T) {
	t5 := tree.New(5)
	ch1 := make(chan int)
	go Walker(t5, ch1)
	for v := range ch1 {
		fmt.Println(v)
	}
}
