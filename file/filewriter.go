package main

import (
	"fmt"
	"os"
)

func main(){
	var text string
	fmt.Println("Please write text: ")
	fmt.Scan(&text)

	var file_name string
	fmt.Println("Please write file name: ")
	fmt.Scan(&file_name)

	if file_is_exist(file_name){
		file, err := os.OpenFile(file_name,os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil{

			fmt.Println(err)
		}
		defer file.Close()
		file.WriteString(text)
	}else{
		create_file(file_name)
		new_file, err := os.OpenFile(file_name,os.O_APPEND|os.O_WRONLY, 0600)

		if err != nil{
			fmt.Println(err)
		}
		defer new_file.Close()
		new_file.WriteString(text)
	}
}

func file_is_exist(file_name string) bool{

	_, err := os.Stat(file_name)

	if err != nil{
		return false
	}
    return true
}

func create_file(file_name string) {
	_, err := os.Create(file_name)
	if err != nil{
		fmt.Println(err)
	}
}