package main

import (
	"fmt"
	"godrive/message"
	"log"
	"net"
	"os"
)

func main() {
	// commandline processing
	// path manipulation

	userInput := os.Args
	fmt.Printf("%s\n", userInput)
	fmt.Printf("%s\n", userInput[1])
	fmt.Printf("%d\n", len(userInput))

	conn, err := net.Dial("tcp", userInput[1])
	// conn, err := net.Dial("tcp", ":9999") // connect to localhost port 9999
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	defer conn.Close()
	var msg *message.Message

	fileStat, err := os.Stat(userInput[3])
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	fileSize := fileStat.Size()
	log.Printf("File Size: %d\n", fileSize)

	// figure out put or get userInput[2]
	if userInput[2] == "put" {
		msg = message.New(0, fileSize, userInput[3]) //use os.stat
		// msg.Send(conn)
	} else if userInput[2] == "get" {
		msg = message.New(1, fileSize, userInput[3])
		// msg.Get(conn)
	} else if userInput[2] == "search" {
		msg = message.New(2, fileSize, userInput[3])
	} else if userInput[2] == "delete" {
		msg = message.New(3, fileSize, userInput[3])
	} else {
		log.Fatalln(err.Error())
	}

	// something := message.SearchRequest
	// m := message.Message{Name: "GoDrive"}
	// fmt.Println(m, something)

	// msg := message.New(message.StorageRequest, 300, userInput[3])
	// fmt.Printf("%T\n", msg)
	msg.Print()
	msg.Send(conn) // pass in our connection

	// move to the constructor open the file
	// file, err := os.OpenFile("test.txt", os.O_RDONLY, 0666)
	// if err != nil {
	// 	log.Fatalln(err.Error())
	// 	return
	// }
	// io.Copy(conn, file)
}
