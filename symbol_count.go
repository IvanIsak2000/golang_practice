package main

import "fmt"

func get_string() string {
	var s string
	fmt.Println("Please enter a letters string: ")
	fmt.Scan(&s)
	return s
}

func main() {
	get_str := get_string()
	count := 0

	for i := 0; i < len(get_str); i++ {
		count++
	}
	fmt.Println("All symbols without spaces: ", count)
}
