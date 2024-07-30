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

func functionExit(args ...string) error{
	os.Exit(0)
	return nil
}

func functionHelp(args ...string) error {
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


func functionmap(args ...string) error { 
	default_url := "https://pokeapi.co/api/v2/location-area/" 
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


func functionmapb(args ...string) error { 
	default_url := "https://pokeapi.co/api/v2/location-area/" 
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


func functionEXplore(args ...string) error {	
	default_url := "https://pokeapi.co/api/v2/location-area/"
	default_url+=args[0]
	err := pokeapi.Explore(default_url)
	if err != nil{
		return err
	}
	return nil

}























