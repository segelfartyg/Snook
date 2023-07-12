package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)


func TickTock(ws *websocket.Conn){
	
	buf := make([]byte, 1024)

	n, err := ws.Read(buf)
	
	
	if err != nil {
	
			fmt.Println(err)
	}

	msg := buf[:n]

	ws.Write(msg)

	println(n)
	println(string(msg))
}

func main() {
	

	
	http.Handle("/Tick", websocket.Handler(TickTock))
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		panic("Listening failed" + err.Error())
	}



}