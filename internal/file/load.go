package file

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Define struct to be used for sample dataset
type Films struct {
	Films []Film `json:"films"`
}

type Film struct {
	Name     string `json:"name"`
	Year     int    `json:"year"`
	Genre    string `json:"genre"`
	Language string `json:"language"`
	Social   Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func JsonParser() {
	// Open the JSON file
	jsonFile, err := os.Open("dataset.json")

	// Error handling
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened dataset.json")

	// defer the closing of the json file
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var films Films

	json.Unmarshal(byteValue, &films)

	for i := 0; i < len(films.Films); i++ {
		fmt.Println("Film Name: " + films.Films[i].Name)
		fmt.Println("Film Year: " + strconv.Itoa(films.Films[i].Year))
		fmt.Println("Film Gengre: " + films.Films[i].Genre)
		fmt.Println("Film Language: " + films.Films[i].Language)
		fmt.Println("Facebook URL: " + films.Films[i].Social.Facebook)
	}
}
