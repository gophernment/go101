package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	f, err := os.Open("oscar_male.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)

	data, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	_ = data
}
