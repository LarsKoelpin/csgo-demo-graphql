package usecase

import (
	"log"
	"os"
)

func CreateJsonFile(filename string, json string) {
	file, _ := os.Create(filename)
	log.Print("Writing JSON file ...")
	file.Write([]byte(json))
}
