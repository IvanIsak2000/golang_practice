package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){

	arg1 := os.Args[1]
	arg2 := os.Args[2]
	arg3 := os.Args[3]
	arg4 := os.Args[4]


	err_message := "Must be: go run . <digits [1-yes, no - not 1]>  <lettes [1-yes, no - not 1]> <special symbol 1-yes, no - not 1]> <length [integer number]>\nExample: go run . 1 1 0 16\nResult:h1rd4ybkjij3vc52"


	digits, err := strconv.Atoi(arg1)
    if err != nil {
            fmt.Println(err_message)
	}
	letters, err := strconv.Atoi(arg2)
    if err != nil {
            fmt.Println(err_message)
	}

	symbols, err := strconv.Atoi(arg3)
    if err != nil {
            fmt.Println(err_message)
	}


	lenght, err := strconv.Atoi(arg4)
    if err != nil {
            fmt.Println(err_message)
	}


	res := make_password(digits, letters, symbols, lenght)
	fmt.Println(res)
}