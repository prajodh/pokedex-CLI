package pokerepl

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)




type Commands struct{
	name string
	description string
	function func(...string) error
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
	commands["explore"] = Commands{
		name:"mapb",
		description: "explore pokemon in a given location",
		function: functionEXplore,
	}
	commands["catch"] = Commands{
		name:"catch",
		description: "try and catch the pokemon",
		function: functionCatch,
	}
	commands["help"] = Commands{
		name : "help",
		description :  "Welcome to the Pokedex!\n\nUsage:\nhelp: Displays a help message \n\n",
		function: functionHelp,
	}
	return commands
}


func StartRepl(){
	
	commands:=loadCommands()
	for {
		fmt.Printf("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		if text == ""{
			continue
		}
		var arguments []string
		args := strings.Split(text, " ")
		if len(args) > 1{
			for i := range(len(args)){
				if i == 0{
					continue
				}
				arguments = append(arguments, args[1])
			}
		}
		command, ok:= commands[args[0]]
		if !ok{
			fmt.Println("invalid command")
			continue
		}
		err := command.function(arguments...)
		if err!= nil{
			log.Fatal(err)
		}
	}
}