package main

import "fmt"

func main() {

	var (
		a, b, c float64
	)

	fmt.Print("a=")
	fmt.Scan(&a)

	fmt.Print("b=")
	fmt.Scan(&b)

	fmt.Print("с=")
	fmt.Scan(&c)

	x := -(b / 2 * a)
	y := a*(x*x) + b*x + c
	fmt.Print("\nThe vertex of the parabola at the point\nx=", x, "\ny=", y)

}
