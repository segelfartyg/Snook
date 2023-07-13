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

		var msg = make([]byte, 512)
		n, err := ws.Read(msg); if err != nil{
			fmt.Println(err.Error())
			continue
		}

		readMessage := msg[:n]
		fmt.Println(string(readMessage))

		go sendMessage(ws)
	}
}

func sendMessage(ws *websocket.Conn){

	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')

	if _, err := ws.Write([]byte(text)); err != nil {
		fmt.Println(err.Error())
	}


}