package main

import (
	"fmt"
	"log"
	"encoding/json"
)

type Movie struct {
    Title  string
    Year   int  `json:"released"`
    Color  bool `json:"color,omitempty"`
    Actors []string
}

var movies = []Movie{
    {Title: "Casablanca", Year: 1942, Color: false,
        Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
    {Title: "Cool Hand Luke", Year: 1967, Color: true,
        Actors: []string{"Paul Newman"}},
    {Title: "Bullitt", Year: 1968, Color: true,
        Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
    // ...
}

func main() {
	data, err := json.MarshalIndent(movies, "", "	")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("%s\n", data)
	var deco []Movie
	err = json.Unmarshal(data, &deco)
	fmt.Println(deco[0].Title)
}