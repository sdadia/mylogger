package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {

	body, err := os.ReadFile("data.txt")
	if err != nil {
		return
	}
	log.Printf("Data is \n%s", body)

	f, _ := os.Open("./data.txt")
	reader := csv.NewReader(f)
	line, err := reader.ReadAll()
	if err != nil {
		return
	}
	log.Printf("Data is %s", line)

}
