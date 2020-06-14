package usecase

import (
	"encoding/json"
	"log"

	"github.com/graphql-go/graphql"
)

func CreateJson(schema graphql.Schema, query string) string {
	params := graphql.Params{Schema: schema, RequestString: query}
	log.Print("Request using ", query)
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	log.Print("Start mashalling ...");
	rJSON, _ := json.Marshal(r)

	return string(rJSON)
}
