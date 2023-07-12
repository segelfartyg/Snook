package main

import SnookWebSocket "Snook.Client/Snook.Client.Modules/Snook.Client.Modules.SnookWebSocket"

func main() {

	ws := SnookWebSocket.Init()
	SnookWebSocket.ReadConn(ws)
	
}