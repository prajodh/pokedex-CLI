# Pokedex CLI

A command-line Pokedex application built in Go. This application allows you to explore the world of Pokémon using the PokeAPI, catch Pokémon, and inspect the Pokémon you've caught. The application is designed to be interactive and easy to use, operating in a REPL (Read-Eval-Print Loop) environment.

## Features

- **REPL Environment**: Interact with the Pokedex using simple commands.
- **Explore the Pokémon World**: Use the `map` and `mapb` commands to navigate through Pokémon locations.
- **Caching**: Fast access to previously visited locations through a caching mechanism.
- **Catch Pokémon**: Attempt to catch Pokémon in specific locations and add them to your Pokedex.
- **Inspect Pokémon**: View detailed information about Pokémon you've caught.
- **Pokedex Overview**: See a list of all Pokémon you've caught.

## Commands

### `help`
Displays a help message describing how to use the REPL.

### `exit`
Exits the program.

### `map`
Displays the names of 20 location areas in the Pokémon world. Each subsequent call displays the next 20 locations.

### `mapb`
Displays the previous 20 location areas. If you're on the first "page," this command will print an error message.

### `explore <location-area>`
Explores the specified location area, listing the Pokémon that can be found there.

### `catch <pokemon-name>`
Attempts to catch the specified Pokémon by name. The success rate depends on the Pokémon's base experience.

### `inspect <pokemon-name>`
Displays detailed information about the specified Pokémon if it has been caught. If not, it will inform you that the Pokémon hasn't been caught yet.

### `pokedex`
Displays a list of all the Pokémon you've caught.

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/prajodh/pokedex-CLI.git
   cd pokedex-cli
   ```
   
2. **Build the application:**

   
  ```bash
  go build -o pokedexcli
  ```
3. **Run application**


  ```bash
  go build -o pokedexcli
  ```




