package main

import (
	"fmt"
	"time"
	"net"
	"os"

	"github.com/fatih/color"
)

func main(){
	host_port := get_host_port()
	fmt.Printf("Listening:\nHost: %s\nPort: %s\n______________________\n",host_port[0],host_port[1] )
	connect(host_port)

}
	
func get_host_port() []string{	

	host := os.Args[1]
	port := os.Args[2]

	if host ==   "_" && port == "_"{
		host_port := []string{"localhost", "8080"}
		return host_port

	}else{
		host_port := []string{host, port}
		return host_port
	} 
	
}


func connect(host_port []string) string{

	listener, _ := net.Listen("tcp",host_port[0] + ":"+ host_port[1])
	
	for { 
		color.Yellow("Wait connection...")
		conn, err := listener.Accept() 

		if err != nil{
			fmt.Println(err)
		}

		current_time := time.Now().Format("02-01-2006 15:04:05")
		conn.Close()

		fmt.Println("Connection from: ", conn)
		data := fmt.Sprintf("From:%s\nTime:%s\n_______________________", conn, current_time)
		
		write_in_file(data)
	}
			
}

func write_in_file(data string){
	file_name := "server_log.txt" 

	file, err := os.Open(file_name)

	if err != nil{
		fmt.Println("File was created!")
		new_file, err := os.Create(file_name)
		if err != nil{
			fmt.Println(err)
		}

		new_file.WriteString(data)
		new_file.Close()
	}

	file.WriteString(data)
	file.Close()
}