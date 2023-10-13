package jsons

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Dimension struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Data struct {
	Species     string    `json:"species"`
	Description string    `json:"description"`
	Dimension   Dimension `json:"dimension"`
}

func Jsons() {
	f, err := os.Open("input.json")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var data Data

	// jsonを構造体に変換する
	err = json.NewDecoder(f).Decode(&data)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", data)
}
