package pokerepl

import (
	"bufio"
	"fmt"
	"os"
	"github.com/prajodh/pokedex-Cli-golang/pokeapi"
	"github.com/prajodh/pokedex-Cli-golang/pokecache"
	"errors"
)


type Commands struct{
	name string
	description string
}

func help(c map[string]Commands) string{
	res := "Welcome to the Pokedex!\n\nUsage:\nhelp: Displays a help message \n\n"
	for _,val:= range c{
		res += val.name+": "+val.description+"\n\n"
	}
	return res
}

func loadCommands() map[string]Commands{
	commands := make(map[string]Commands)
	
	commands["exit"] = Commands{
		name : "exit",
		description : "Exit the Pokedex",
	}
	commands["map"] = Commands{
		name:"map",
		description: "gets next new 20 locations",
	}
	commands["mapb"] = Commands{
		name:"mapb",
		description: "move back to prev 20 locations",
	}
	helpDescription := help(commands)
	commands["help"] = Commands{
		name : "help",
		description : helpDescription,
	}
	return commands
}

func location(pagination *int){
	for i:=1; i < 21; i++{
	res, err := pokeapi.GetLocations(*pagination+i)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(res)
	}
	*pagination = (*pagination + 20) % 1036
}

func locationReverse(pagination *int){
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


func StartRepl(){
	pokecache.Pokecache()
	var pagination int = 0
	var paginationPtr *int = &pagination
	commands:=loadCommands()
	for i:=0 ;; i++{
		fmt.Printf("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		test := scanner.Text()
		if val, ok := commands[test]; ok{
			if val.name == "help"{
			fmt.Printf(val.description)
			} else if val.name == "exit"{
				os.Exit(0)
			} else if val.name == "map"{
				location(paginationPtr)
			} else if val.name == "mapb"{
				locationReverse(paginationPtr)
			} 
		}else{
				fmt.Println("Not a Valid Command")
			}
	}
}