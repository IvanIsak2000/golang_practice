package main

import (
	"fmt"
	"math"
)

func getHalfAxle() float64 {
	var halfAxle float64
	fmt.Print("Please enter half-axle of a planet in astronomical unit(for example: 1.5): ")
	fmt.Scan(&halfAxle)
	return halfAxle
}

func main() {
	var res float64
	halfAxle := getHalfAxle()
	res = halfAxle * math.Sqrt(halfAxle)
	fmt.Print(res, "year/s")
}
