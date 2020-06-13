package usecase

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/larskoelpin/csgo-demo-graphql/pkg/domain"
)

func SchemaFromDemo(demo domain.Demo) graphql.Schema {
	demoType := domain.CreateDemoType(demo)
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: graphql.Fields{
		"demo": {
			Name: "Demo",
			Type: demoType,
			Args: nil,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return demo, nil
			},
			DeprecationReason: "",
			Description:       "",
		},
	}}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	return schema
}
