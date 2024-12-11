package main

import (
	"fmt"
	"math"
)

func Area(a int, b int, c int) int32 {

	p := (a + b + c) / 2
	act1 := p * (p - a) * (p - b) * (p - c)
	act2 := math.Pow(float64(act1), (0.5))

	return int32(act2)

}

func main() {

	var a int
	fmt.Println("Enter a side:")
	fmt.Scan(&a)

	var b int
	fmt.Println("Enter b side:")
	fmt.Scan(&b)

	var c int
	fmt.Println("Enter c side:")
	fmt.Scan(&c)

	area := Area(a, b, c)
	fmt.Println("Area= ", area)
}
