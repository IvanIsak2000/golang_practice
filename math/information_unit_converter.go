package main

import (
	"fmt"
	"math"
)

func getNumber() float64 {
	var number float64
	fmt.Print("Please enter a number: ")
	fmt.Scan(&number)
	return number
}

func getUnitOfInformation() int {
	var numberUnit int
	fmt.Println("\n(1)Bit\n(2)Byte\n(3)Kbyte\n(4)Mbyte\n(5)Gbyte\n(6)Tbyte")
	fmt.Scan(&numberUnit)
	return numberUnit
}

func calculation(number float64, primaryUnit int, finishUnit int) float64 {
	p := float64(primaryUnit)
	f := float64(finishUnit)
	if f != p {
		if p == 1 || p == 2 && f == 1 || f == 2 {
			if p > f {
				return number / 8
			}
			return number * 8
		}

		if p > f {
			diff := p - f
			return number * math.Pow(1024, diff)
			//switch diff {
			//case 1:
			//	return number * 1024
			//case 2:
			//	return number * 1024 * 1024
			//
			//}

		}
		if f > p {
			return number / (1024 * (p - f))
		}
	}
	return -1
}


func main() {
	units := [...]string{1: "Bit", 2: "Byte", 3: "Kbyte", 4: "Mbyte", 5: "Gbyte", 6: "Tbyte"}
	number := getNumber()

	fmt.Print("Please write a primary the  unit of information:")
	primaryUnit := getUnitOfInformation()

	fmt.Print("Please write the finish unit of information:")
	finishUnit := getUnitOfInformation()

	answer := calculation(number, primaryUnit, finishUnit)
	if answer != -1 {
		fmt.Printf("Answer: %f %s = %f %s", number, units[primaryUnit], answer, units[finishUnit])
	}
}
