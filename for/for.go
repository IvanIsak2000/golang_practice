package main

import (
	"fmt"
	"time"
)


func EvenNumbers(count int){
	for i := 0; i <= count; i++{
		if i%2 == 0{
			fmt.Println(i)
		}
	}
}

func UnEvenNumbers(count int){
	for i := 0; i <= count; i++{
		if i%2 != 0{
			fmt.Println(i)
		}
	}
}


func main() {
	count := 10
	go EvenNumbers(count)
	go UnEvenNumbers(count)
	time.Sleep(5 * time.Second)
}
