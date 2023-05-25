//скрипт для вычисления площадей поверхности цилиндра

package main

import (
	"fmt"
	// "math"
)

const (
	pi float32 = 3.14
)

func get_vars() (float32, float32) {
	var ares float32
	fmt.Print("Enter a ares: ")
	fmt.Scan(&ares)

	var height float32
	fmt.Print("Enter a height: ")
	fmt.Scan(&height)

	return ares, height

}

func vars_is_valid(ares, height float32) bool {

	if ares != 0 && height != 0 {
		return true
	}

	return false

}

func main() {

	r, h := get_vars()

	if vars_is_valid(r, h) {
		Sq_base := pi * r * r
		Sq_side_surface := 2 * pi * r * h
		Sq_full_surface := 2 * pi * r * (r + h)

		fmt.Println("Sq base: ", Sq_base)
		fmt.Println("Sq side surface: ", Sq_side_surface)
		fmt.Println("Sq full surface ", Sq_full_surface)
	} else {
		fmt.Print("Is not valid vars!!!")
	}

}
