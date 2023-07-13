package Config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Configuration struct {

	Servers [] Server `json:"servers"`
	User User `json:"user"`

}

type Server struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Url string `json:"url"`
	ChatUrl string `json:"chaturl"`
	Origin string `json:"origin"`
}

type User struct {
	UserName string `json:"name"`
	Motto string `json:"motto"`
}


func NewConfiguration() *Configuration{
	return &Configuration{}
}

func (c *Configuration)ConfigureClient(){
	fmt.Println("Configuration started")

	menuConfig, err := os.Open("./Snook.Client.Modules/Snook.Client.Modules.Config/Config.json")

	if err != nil {
		fmt.Println(err)
	}
	defer menuConfig.Close()
	
	byteValue, _ := ioutil.ReadAll(menuConfig)
	
	var configs Configuration
	json.Unmarshal(byteValue, &configs)
	

	fmt.Println(configs)

	c.Servers = configs.Servers
	c.User = configs.User
	

}