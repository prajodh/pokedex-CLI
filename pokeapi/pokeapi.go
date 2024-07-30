package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"github.com/prajodh/pokedex-Cli-golang/pokecache"
)


var cache = pokecache.CreateCache()

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