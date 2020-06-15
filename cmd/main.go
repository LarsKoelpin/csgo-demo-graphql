package main

import (
	"log"

	"github.com/larskoelpin/csgo-demo-graphql/pkg/domain"
	usecase "github.com/larskoelpin/csgo-demo-graphql/pkg/usecase"
)

func main() {
	demoRepository := domain.DemoRepository{}
	log.Print("Reading User query ...")
	userQuery := usecase.ReadQuery("examples/events-query.query")
	log.Print("Creating a schema ...")
	schema := usecase.SchemaFromDemo(demoRepository)
	json := usecase.CreateJson(schema, userQuery)
	usecase.CreateJsonFile("examples/example.json", json)
}
