package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Output   string    `json:"output"`
	Students []Student `json:"students"`
}

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Marks int    `json:"marks"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: app <input.json>")
		os.Exit(1)
	}

	inputFile := os.Args[1]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		log.Fatal(err)
	}

	if cfg.Output == "" {
		log.Fatal("output file not specified")
	}

	file, err := os.Create(cfg.Output)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"name", "age", "marks"})

	for _, s := range cfg.Students {
		writer.Write([]string{
			s.Name,
			strconv.Itoa(s.Age),
			strconv.Itoa(s.Marks),
		})
	}

	fmt.Println("CSV written to:", cfg.Output)
}
