package usecase

import (
	"log"
	"os"
)

func CreateJsonFile(filename string, json []byte) {
	file, _ := os.Create(filename)
	log.Print("Writing JSON file ...")
	file.Write([]byte(json))
}
