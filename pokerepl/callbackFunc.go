package pokerepl

import (
	// "errors"
	"fmt"
	"os"
	"github.com/prajodh/pokedex-Cli-golang/pokeapi"
)

type NextPrev struct{
	nextUrl string
	prevUrl string
}

var state NextPrev

func functionExit() error{
	os.Exit(0)
	return nil
}

func functionHelp() error {
	c := loadCommands()
	res := c["help"].description
	for _,val:= range c{
		if val.name == "help"{
			continue
		}
		res += val.name+": "+val.description+"\n\n"
	}
	fmt.Println(res)
	return nil
}


func functionmap() error { 
	default_url := "https://pokeapi.co/api/v2/location/" 
	if state.nextUrl != ""{
		default_url = state.nextUrl
	}
	nextUrl, prevUrl, err := pokeapi.Map(default_url)
	if err != nil{
		return err
	}
	state.nextUrl = nextUrl
	state.prevUrl = prevUrl
	return nil
}


func functionmapb() error { 
	default_url := "https://pokeapi.co/api/v2/location/" 
	if state.prevUrl != ""{
		default_url = state.prevUrl
	}else{
		fmt.Println("you are already at the start")
	}
	nextUrl, prevUrl, err := pokeapi.Map(default_url)
	if err != nil{
		return err
	}
	state.nextUrl = nextUrl
	state.prevUrl = prevUrl
	return nil
}























