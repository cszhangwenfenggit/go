package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Movie struct {
		Title string
		//Year   int `json:"released"`
		Year   int
		Actors []string
	}

	var movies = []Movie{
		{
			Title:  "钢铁侠",
			Year:   2020,
			Actors: []string{"Robert", "Alice"},
		},
		{
			Title:  "变形金刚",
			Year:   2021,
			Actors: []string{"大黄蜂", "擎天柱"},
		},
	}

	data, err := json.Marshal(movies)
	if err != nil {
		fmt.Printf("JSON marshaling fail:%s", err)
	}

	fmt.Printf("%s\n", data)

	//var titles []struct{ Title string }
	var titles []struct{ Year int }

	if err := json.Unmarshal(data, &titles); err != nil {
		fmt.Printf("JSON unmarshal fail:%s", err)
	}
	fmt.Println(titles)
}
