package main

import (
	"math/rand"
	"time"
)

func makePassword(digits, letters, symbols, lenght int) string {
	toMake := ""
	res := ""

	if digits == 1 {
		toMake += "0123456789"
	}

	if letters == 1 {
		toMake += "abcdefghijklmnopqrstuvwxyz"
	}

	if symbols == 1 {
		toMake += "!$%&()*+,-.:;<=>?@[]^_{|}~"
	}

	for i := 0; i < lenght; i++ {

		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(len(toMake))
		res += toMake[index : index+1]
	}
	return res

}
