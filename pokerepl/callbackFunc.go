package pokerepl

import (
	// "errors"
	"fmt"
	"os"
	"github.com/prajodh/pokedex-Cli-golang/pokeapi"
	"math/rand"
)

var pokedexCaught = make(map[string]pokeapi.Pokemon)
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


func functionCatch(args ...string) error{
	default_url := "https://pokeapi.co/api/v2/pokemon/"
	default_url += args[0]
	pokemon, err := pokeapi.Catch(default_url)
	if err != nil {
		return err
	}
	if catchPokemon(args[0], pokemon.BaseExperience){
		_,ok := pokedexCaught[args[0]]
		if !ok{
			pokedexCaught[args[0]] = pokemon
		}
	}
	return nil
}


func catchPokemon(name string, basehealth int) bool {
	fmt.Println("throwing a ball at "+ name+"...")
	for {
		minhp := 55
		maxhp := 420
		chance := rand.Intn(maxhp - minhp + 1) + minhp
		if chance > basehealth{
			fmt.Println(name+" was caught")
			return true
		} else {
			fmt.Println(name+" fled the scene")
			return false
		}
	}
}





















