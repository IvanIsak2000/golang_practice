package main

import (
	"math/rand"
	"time"
)

func make_password(digits, letters, symbols, lenght int )string{
	to_make := ""
	res := ""

	if digits == 1{
		to_make += "0123456789"
	}

	if letters == 1 {
		to_make += "abcdefghijklmnopqrstuvwxyz"
	}

	if symbols == 1{
		to_make +=  "!$%&()*+,-.:;<=>?@[]^_{|}~"
	}

	for i:=0; i< lenght; i++{

		rand.Seed(time.Now().UnixNano())
		index:=(rand.Intn(len(to_make)))
		res +=to_make[index:index+1]
	}
	return res


}