package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"

	// "fmt"
	"io"
	"net/http"
	// "strconv"
)
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
	res ,err := http.Get(url)
	if err != nil{
		return "", "", errors.New("error wil fetching the location")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil{
		return "", "", errors.New("error will reading bytes")
	}
	loc := &LocationAreas{}
	err_ := json.Unmarshal(body, loc)
	if err_ != nil{
		return "", "", errors.New("error while unmarshalling")
	}
	for i:= 0; i < len(loc.Results); i++{
		fmt.Println("-"+loc.Results[i].Name)
	}
	return loc.Next, loc.Previous, nil
}