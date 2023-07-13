package main

import (
	"fmt"

	Chat "Snook.Client/Snook.Client.Modules/Snook.Client.Modules.Chat"
	Config "Snook.Client/Snook.Client.Modules/Snook.Client.Modules.Config"
	Menu "Snook.Client/Snook.Client.Modules/Snook.Client.Modules.Menu"
)

func main() {

	// CONFIGURE THE CLIENT
	configuration := Config.NewConfiguration()
	configuration.ConfigureClient()
	fmt.Println("Welcome, ", configuration.User.UserName)
	fmt.Println("You will connect through the server: ", configuration.Servers[0].Name)

	// ASSEMBLING MENU
	menu := Menu.NewMenu()
	menu.AssembleMenu()
	fmt.Println(menu.MenuItems[0])

	// INITIALIZING THE CHAT
	ws := Chat.Init(configuration.Servers[0].ChatUrl, configuration.Servers[0].Origin)
	Chat.ReadConn(ws, configuration.User.UserName)
}