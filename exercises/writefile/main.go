package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Create("textfile.txt")
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write([]byte{65, 114, 105, 115, 101})
	if err != nil {
		log.Fatal(err)
	}

}
