package Chat

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"golang.org/x/net/websocket"
)

var ws *websocket.Conn

func Init(chaturl string, origin string)(*websocket.Conn) {
	fmt.Println("client starting")
	ws, err := websocket.Dial(chaturl, "", origin) 
	if err != nil {
		log.Fatal(err.Error())	
	}
	return ws
}


func ReadConn(ws *websocket.Conn, username string){

	for {
		go sendMessage(ws, username)
		var msg = make([]byte, 512)
		n, err := ws.Read(msg); if err != nil{
			fmt.Println(err.Error())
			continue
		}

		readMessage := msg[:n]
		fmt.Println(string(readMessage))

		
	}
}

func sendMessage(ws *websocket.Conn, username string){

	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')

	usernameMessage := username + ": " + text

	if _, err := ws.Write([]byte(usernameMessage)); err != nil {
		fmt.Println(err.Error())
	}


}