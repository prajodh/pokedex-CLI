package pokerepl

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"github.com/prajodh/pokedex-Cli-golang/pokecache"
)


type Commands struct{
	name string
	description string
	function func() error
}

func loadCommands() map[string]Commands{
	commands := make(map[string]Commands)
	
	commands["exit"] = Commands{
		name : "exit",
		description : "Exit the Pokedex",
		function: functionExit,
	}
	commands["map"] = Commands{
		name:"map",
		description: "gets next new 20 locations",
		function: functionmap,

	}
	commands["mapb"] = Commands{
		name:"mapb",
		description: "move back to prev 20 locations",
		function: functionmapb,
	}
	commands["help"] = Commands{
		name : "help",
		description :  "Welcome to the Pokedex!\n\nUsage:\nhelp: Displays a help message \n\n",
		function: functionHelp,
	}
	return commands
}


func StartRepl(){
	pokecache.Pokecache()
	
	commands:=loadCommands()
	for {
		fmt.Printf("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		if text == ""{
			continue
		}
		command, ok:= commands[text]
		if !ok{
			fmt.Println("invalid command")
			continue
		}
		err := command.function()
		if err!= nil{
			log.Fatal(err)
		}
	}
}