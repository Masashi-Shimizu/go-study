package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(math.Pi)
	fmt.Println(add(42,31))
	fmt.Println()
}

func add(a int,b int) int {
	return a + b
}

func swap(x, y string) (string, string) {
	return y, x
}