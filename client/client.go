package main

import (
	"encoding/gob"
	"fmt"
	"net"

	"./process"
)

var p *process.Process

func closeConection() {
	fmt.Println("Hey now launch defer")
	client, err := net.Dial("tcp", ":3000")
	if err != nil {
		fmt.Println(err)
		return
	}
	p.NewProcess = false
	p.Kill()
	err = gob.NewEncoder(client).Encode(p)
	if err != nil {
		fmt.Println(err)
	}

}

func client() {
	client, err := net.Dial("tcp", ":3000")
	if err != nil {
		fmt.Println(err)
		return
	}
	p = &process.Process{0, 0, false, true}
	err = gob.NewEncoder(client).Encode(p)
	err = gob.NewDecoder(client).Decode(&p)
	p.Iterate = true
	p.Start()
	if err != nil {
		fmt.Println(err)
	}
	client.Close()
}

func main() {
	go client()
	defer closeConection()
	var input string
	fmt.Scanln(&input)
}
