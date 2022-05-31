package golang

import (
	"fmt"
	"testing"
)

func sum(c chan int, input ...int) {
	acc := 0
	for _, i := range input {
		acc += i
	}
	c <- acc
}

func TestChannel(t *testing.T) {
	c := make(chan int)
	go sum(c, 5, 5)
	go sum(c, 1, 1)
	sum1, sum2 := <-c, <-c
	fmt.Println(sum1, sum2)
}

func fibonacciChannel(n int, c chan int) {
	x, y := 0, 1
	for i := 2; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func TestChannel_Fibonacci(t *testing.T) {
	c := make(chan int, 10)
	go fibonacciChannel(cap(c), c)
	for x := range c {
		fmt.Println(x)
	}
}

func TestChannel_Fibonacci_Select(t *testing.T) {
	c := make(chan interface{})
	quit := make(chan interface{})

	x, y := 0, 1
	go func() {
		for i := 0; i < 10; i++ {
			c <- nil
		}
		quit <- nil
	}()

	for {
		select {
		case <-c:
			x, y = y, x+y
			fmt.Println(x)
		case <-quit:
			return
		}
	}
}
