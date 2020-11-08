package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"

	"./process"
)

var list []*process.Process

func server() {
	server, err := net.Listen("tcp", ":3000")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		client, err := server.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleClient(client)
	}
}
func handleClient(client net.Conn) {
	var process *process.Process
	err := gob.NewDecoder(client).Decode(&process)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("hey new client conected and send ", process)
	if process.NewProcess == true {
		process, list = list[len(list)-1], list[:len(list)-1]
		process.Kill()
		fmt.Println(process)
		err = gob.NewEncoder(client).Encode(process)
	} else {
		process.Iterate = true
		process.Start()
		list = append(list, process)
	}
	//process.Start()
}
func main() {
	for i := 0; i < 5; i++ {
		process := &process.Process{Id: i, Counter: 0, Iterate: true}
		go process.Start()
		list = append(list, process)
		time.Sleep(time.Second * 1)

	}

	go server()
	fmt.Println("Server listen on port 3000")
	var input string
	fmt.Scanln(&input)
}
