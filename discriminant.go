package main

import (
	"fmt"
	"math"
)

func main() {

	var a float64
	var b float64
	var c float64

	fmt.Print("Write coefficient at a: ")
	fmt.Scan(&a)

	fmt.Print("Write coefficient at b: ")
	fmt.Scan(&b)

	fmt.Print("Write coefficient  c: ")
	fmt.Scan(&c)

	D := (b * b) - 4*a*c

	if D > 0 {

		x1 := (-b + (math.Sqrt(D))) / 2 * a
		x2 := (-b - (math.Sqrt(D))) / 2 * a

		fmt.Println("D = ", D)

		fmt.Println("x1=", x1)
		fmt.Println("x2=", x2)
	}

	if D == 0 {
		x1 := (-b) / 2 * a
		fmt.Println("D = ", D)
		fmt.Println("x1=", x1)
	}

	if D < 0 {
		fmt.Println("D = ", D)
		fmt.Println("No roots!")
	}

}
