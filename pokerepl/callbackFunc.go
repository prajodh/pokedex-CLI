package pokerepl

import (
	"errors"
	"fmt"
	"os"
	"github.com/prajodh/pokedex-Cli-golang/pokeapi"
)

var page int = 0
var pagination *int = &page

func functionExit(){
	os.Exit(0)
}

func functionHelp() {
	c := loadCommands()
	res := c["help"].description
	for _,val:= range c{
		res += val.name+": "+val.description+"\n\n"
	}
	fmt.Println(res)
}


func functionmap(){
	for i:=1; i < 21; i++{
	res, err := pokeapi.GetLocations(*pagination+i)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(res)
	}
	*pagination = (*pagination + 20) % 1036
}

func functionmapb(){
	*pagination -= 40
	if *pagination < 0{
		*pagination = 0
		pagErr := errors.New("your all the way back to the start")
		fmt.Println(pagErr)
	}
	for i:=1; i < 21; i++{
	res, err := pokeapi.GetLocations(*pagination+i)
	if err != nil{
		fmt.Println("errors")
	}
	fmt.Println(res)
	}
	*pagination = (*pagination + 20) % 1036
}























