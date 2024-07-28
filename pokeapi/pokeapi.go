package pokeapi

import (
	"encoding/json"
	"errors"

	// "fmt"
	"io"
	"net/http"
	"strconv"
)

type City struct {
    ID          int      
    Name        string   
    Region      Region   
    Areas       []Area   
    GameIndices []struct {
        GameIndex int 
        Generation struct {
            Name string 
            URL  string 
        } 
    }
}
type Region struct {
    Name string 
    URL  string 
}
type Area struct {
    Name string 
    URL  string 
}

func GetLocations(id int) (string, error) {
	res, err := http.Get("https://pokeapi.co/api/v2/location/"+ strconv.Itoa(id))
	if err != nil{
		return "", errors.New("error wil fetching the location")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil{
		return "", errors.New("error will reading bytes")
	}
	loc := &City{}
	err_ := json.Unmarshal(body, loc)
	if err_ != nil{
		return "", errors.New("error while unmarshalling")
	}
	return loc.Name, nil
}