package usecase

import (
	"log"
	"os"
)

func CreateJsonFile(json string) {
	file, _ := os.Create("examples/test.json")
	log.Print("Writing JSON file ...")
	file.Write([]byte(json))
}
