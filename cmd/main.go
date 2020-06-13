package main

import (
  usecase "github.com/larskoelpin/csgo-demo-graphql/pkg/usecase"
)

func main() {
  userQuery := usecase.ReadQuery("examples/data.query")
	demo := usecase.DemoFromFile("examples/test.dem")
	schema := usecase.SchemaFromDemo(demo)
	json := usecase.CreateJson(schema, userQuery)
	usecase.CreateJsonFile(json);
}
