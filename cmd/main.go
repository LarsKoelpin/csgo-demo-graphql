package main

import (
  "encoding/json"
  "flag"
  "fmt"
  "log"
  "os"
  "time"

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
	queryPtr := fl.String("query", "", "The query file")
	outPath := fl.String("outPath", "", "The resulting JSON")

	err := fl.Parse(os.Args[1:])
	if err != nil {
		// Some parsing problem, the flag.Parse() already prints the error to stderr
		return
	}

	if *queryPtr == "" {
		log.Fatal("No query file specified. try using --query=/path/to/ile  :(", queryPtr)
	}

	if *outPath == "" {
		log.Print("No output specified. Using default current directory")
		*outPath = "./"
	}

	if *inPath == "" {
		log.Fatal("No Input specified. try using --demo=/path/to/file :(")
	}

	file, err := os.Open(*inPath)
	if err != nil {
		log.Fatal("DemoFile does not exist")
	}

	if err != nil {
		log.Fatal("DemoFile does not exist")
	}

	prodQuery := domain.DefaultDemoTemplate
	//schema := usecase.SchemaFromDemo(file, demoRepository)
	//json := usecase.CreateJson(schema, userQuery)
	start := time.Now()
	demo := domain.RecordDemo(file, 16, prodQuery)
	log.Print("Create JSON")
	asJson, _ := json.Marshal(demo)
	usecase.CreateJsonFile("out.json", asJson)
	elapsed := time.Since(start)
	log.Printf("Time rendering took %s", elapsed)
}
