package usecase

import (
	"encoding/json"
	"log"

	"github.com/graphql-go/graphql"
)

func CreateJson(schema graphql.Schema, query string) string {
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)

	return string(rJSON)
}
