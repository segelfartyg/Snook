package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) ConnectionHandler(ws *websocket.Conn) {
	s.conns[ws] = true
	s.ReadLoop(ws)
}


func (s *Server) ReadLoop(ws *websocket.Conn){
buf := make([]byte, 1024)
	for {

		n, err := ws.Read(buf)

		if err != nil {

			fmt.Println(err)
		}
		
	msg := buf[:n]
	fmt.Println(string(msg))
	fmt.Println("new message")		

	s.MessageToClients(msg)

	}
}

func (s *Server) MessageToClients(msg []byte){
	fmt.Println("Sending message to all clients")
	for connection := range s.conns {
		connection.Write(msg)
	}

}

func main() {
	server := NewServer()
	http.Handle("/ws", websocket.Handler(server.ConnectionHandler))
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		panic("Listening failed" + err.Error())
	}

}
