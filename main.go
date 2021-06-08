package main

import (
	"fmt"
	_ "net/http/pprof"
	"testing/math"
	"time"
)

func main() {
	fmt.Println("Addition : ", add(3, 4))
	fmt.Println("Subtraction : ", subtract(3, 4))
	fmt.Println("Divistion : ", math.Div(3, 4))
	time.Sleep(5 * time.Minute)
}

func add(a, b int) int {
	var c int
	c = a + b
	if a > 10 {
		c = c + 5
	}
	return c
}

func subtract(a, b int) int {
	var c int
	c = b - a
	return c
}
