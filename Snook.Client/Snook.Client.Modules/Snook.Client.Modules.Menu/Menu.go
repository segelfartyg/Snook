package Menu

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Menu struct {
	MenuItems map[int]string
}

type SnookModules struct{
	Names []string `json:"snookModules"`
}

func NewMenu() *Menu {
	return &Menu{}
}

// retrieves config from MenuConfig file, to fill the menu that the method is called from.
func (m *Menu)AssembleMenu(){

	menuConfig, err := os.Open("./Snook.Client.Modules/Snook.Client.Modules.Menu/MenuConfig.json")

	if err != nil {
		fmt.Println(err)
	}
	defer menuConfig.Close()
	
	byteValue, _ := ioutil.ReadAll(menuConfig)
	
	var modules SnookModules
	json.Unmarshal(byteValue, &modules)
	
	// initialize the menu
	items := make(map[int]string)
	for i := 0; i < len(modules.Names); i++ {
		items[i] = modules.Names[i]
	}
	m.MenuItems = items
}