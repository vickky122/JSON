package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ToDo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"-"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/200"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		todoItem := ToDo{}

		decoder := json.NewDecoder(response.Body)
		decoder.DisallowUnknownFields()

		if err := decoder.Decode(&todoItem); err != nil {
			log.Fatal("Decode error: ", err)
		}

		fmt.Printf("Data from Api: %+v\n", todoItem)
	}
}
