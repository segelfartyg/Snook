package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

func ConnectionHandler(ws *websocket.Conn) {

	buf := make([]byte, 1024)
	for {

		n, err := ws.Read(buf)

		if err != nil {

			fmt.Println(err)
		}
		
	msg := buf[:n]

	ws.Write(msg)

	fmt.Println(string(msg))
	fmt.Println("new message")		
	}

}

func main() {

	http.Handle("/ws", websocket.Handler(ConnectionHandler))
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		panic("Listening failed" + err.Error())
	}

}
