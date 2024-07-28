module github.com/prajodh/pokedex-Cli-golang/pokerepl

go 1.22.5


replace github.com/prajodh/pokedex-Cli-golang/pokeapi v0.0.0 => ../pokeapi
replace github.com/prajodh/pokedex-Cli-golang/pokecache v0.0.0 => ../pokecache

require( 
        github.com/prajodh/pokedex-Cli-golang/pokeapi v0.0.0
        github.com/prajodh/pokedex-Cli-golang/pokecache v0.0.0

)