package main

import "fmt"

const one_joules = 0.00000000000000000016

func select_mode() int {
	var mode int
	fmt.Scan(&mode)
	return mode
}
func get_number() float32 {
	var number float32
	fmt.Scan(&number)
	return number

}

func main() {

	fmt.Print("Welcome to translation!\nPlease select mode by number:\n" +
		"[1] Joules -> electron volts\n" +
		"[2] Electron volt -> joules\nMode: ")

	switch select_mode() {

	case 1:
		fmt.Print("Please write a joules number [for example: 4.16*10^19]: ")
		number := get_number()
		result := number / one_joules
		fmt.Printf("___________________________________\n"+
			"In exponential representation: %E electron volts", result)
		fmt.Printf("\nEqual: %.f electron volts\n", result)

	case 2:
		fmt.Print("Please write a electron volts number [for example: 2.59]:  ")
		number := get_number()
		result := number * one_joules
		fmt.Printf("\n__________________________________\n%f electron volts = %E joules\n", number, result)

	default:
		fmt.Println("Error!\nPlease write 1 or 2!!!")

	}
}
