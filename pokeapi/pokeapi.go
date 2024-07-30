package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/prajodh/pokedex-Cli-golang/pokecache"
)


var cache = pokecache.CreateCache(time.Minute* 10) // interval after which we clear the cache

type LocationAreas struct {
	Count    int    
	Next     string 
	Previous string    
	Results  []struct {
		Name string 
		URL  string 
	} 
}


func Map(url string)  (string, string, error){
	var data []byte
	data, err := cache.Get(url)
	if err != nil {
		res ,err := http.Get(url)
		if err != nil{
			return "", "", errors.New("error wil fetching the location")
		}
		data, err = io.ReadAll(res.Body)
		if err != nil{
			return "", "", errors.New("error will reading bytes")
		}
		cache.Add(url, data)
	}
	loc := &LocationAreas{}
	err_ := json.Unmarshal(data, loc)
	if err_ != nil{
		return "", "", errors.New("error while unmarshalling")
	}
	for i:= 0; i < len(loc.Results); i++{
		fmt.Println("-"+loc.Results[i].Name)
	}
	return loc.Next, loc.Previous, nil
}



type LocationAreasPokemon struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func Explore(url string) error{
	data, err := cache.Get(url)
	if err != nil{
		res, err := http.Get(url)
		if err != nil{
			return err
		}
		data, err = io.ReadAll(res.Body)
		if err != nil{
			return err
		}
		cache.Add(url, data)
	}
	unmarshallStruct := &LocationAreasPokemon{}
	err_ := json.Unmarshal(data, unmarshallStruct)
	if err_ != nil{
		return err_
	}
	data_json := *unmarshallStruct
	for _, pokemon := range data_json.PokemonEncounters{
		fmt.Println("-"+pokemon.Pokemon.Name)
	}
	return nil

}