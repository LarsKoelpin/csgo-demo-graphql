package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/larskoelpin/csgo-demo-graphql/pkg/domain"
	usecase "github.com/larskoelpin/csgo-demo-graphql/pkg/usecase"
)

func main() {
	fl := new(flag.FlagSet)
	fl.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage of convertcsgo:")
		fl.PrintDefaults()
		fmt.Fprintln(os.Stderr)
	}

	inPath := fl.String("demo", "", "Path to the .dem file")
	freqPtr := fl.String("query", "", "The query file")
	outPath := fl.String("outPath", "", "The resulting JSON")

	err := fl.Parse(os.Args[1:])
	if err != nil {
		// Some parsing problem, the flag.Parse() already prints the error to stderr
		return
	}

	if *freqPtr == "" {
		log.Fatal("No query file specified. try using --query=/path/to/ile  :(", freqPtr)
	}

	if *outPath == "" {
		log.Print("No output specified. Using default current directory")
		*outPath = "./"
	}

	if *inPath == "" {
		log.Fatal("No Input specified. try using --demo=/path/to/file :(")
	}

	demoRepository := domain.DemoRepository{}
	log.Print("Reading User query ...")
	userQuery := usecase.ReadQuery(*freqPtr)
	log.Print("Creating a schema ...")
	file, err := os.Open(*inPath)
	if err != nil {
		log.Fatal("DemoFile does not exist")
	}
	schema := usecase.SchemaFromDemo(file, demoRepository)
	json := usecase.CreateJson(schema, userQuery)
	usecase.CreateJsonFile("out.json", json)
}
