package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type ToDo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	//url := "https://jsonplaceholder.typicode.com/todos/200"
	todoItem := &ToDo{1, 1, "vikrant this side", true}

	//converting back to json( marshal)
	todo, err := json.MarshalIndent(todoItem, "", " ")
	if err != nil {
		log.Fatal("Encode error: ", err)
	}
	fmt.Println(string(todo))
}
