package SnookWebSocket

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"golang.org/x/net/websocket"
)

var ws *websocket.Conn

func Init()(*websocket.Conn) {
	fmt.Println("client starting")

	origin := "http://localhost/"
	url := "ws://localhost:3000/ws"
	ws, err := websocket.Dial(url, "", origin) 
	if err != nil {
		log.Fatal(err.Error())	
	}
	return ws
}


func ReadConn(ws *websocket.Conn){
	for {
		
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("What do you want to send?...")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)

		if _, err := ws.Write([]byte(text)); err != nil {
			fmt.Println(err.Error())
			continue
		}
		
		var msg = make([]byte, 512)
	 
		if _, err := ws.Read(msg); err != nil{
			fmt.Println(err.Error())
			continue
		}
	}
}