package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/net/websocket"
)

func main() {
	fmt.Println("client starting")

	origin := "http://localhost/"
	url := "ws://localhost:3000/Tick"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		fmt.Println("111")
		log.Fatal(err.Error())
		
	}

	for {

		var input string

		fmt.Print("What you want to send?")
		fmt.Scan(&input)
		fmt.Print(input)
		var n1 int
		if n1, err = ws.Write([]byte(input)); err != nil {
			fmt.Println("1")
			fmt.Println(err.Error())
			continue
		}


		var msg = make([]byte, 512)
		var n int 
		if n, err = ws.Read(msg); err != nil{
			fmt.Println("2")
			fmt.Println(err.Error())
			continue
		}
		fmt.Printf("Recieved:", string(msg[:n]))
		fmt.Printf("Recieved:", string(msg[:n1]))

		time.Sleep(time.Second * 2)
	}
	

}