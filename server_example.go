package main

import (
    "fmt"
    "net/http"
)

func connection(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Conencted")
    fmt.Fprintf(w, "Successful connection!")
}


func main() {

    http.HandleFunc("/connection", connection)
	fmt.Println("Server is working")
    http.ListenAndServe(":8000", nil)
}