package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"time"
)

const url string = "http://127.0.0.1:11434/api/generate"

type Ollama struct {
	Model string
	Stream bool
}

type Query struct{
	Ollama *Ollama
	Prompt string
}

func (o *Ollama) StartServer() {
	cmd := exec.Command("ollama", "serve")
	cmd.Run()
	fmt.Println("Ollama server started.")
}

func (q *Query) Ask() int{
	json_data, err := json.Marshal(q)
	if err != nil{
		fmt.Println(err)
	}
	
	request, err := http.NewRequest("POST", url, bytes.NewReader(json_data))
	if err != nil{
		fmt.Println(err)
	}
	client := http.Client{Timeout: 10 * time.Second}
	response, err :=  client.Do(request)
	if err != nil{
		fmt.Printf("Ollama server cannot get request: %s", err)
	}
	fmt.Println(response)
	return response.StatusCode
}