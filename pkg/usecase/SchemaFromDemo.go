package usecase

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/graphql-go/graphql"
	"github.com/larskoelpin/csgo-demo-graphql/pkg/domain"
)

func SchemaFromDemo(demoFile io.Reader, repository domain.DemoRepository) graphql.Schema {
	demoType := domain.CreateDemoType(&repository)
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: graphql.Fields{
		"demo": {
			Name: "Demo",
			Type: demoType,
			Args: graphql.FieldConfigArgument{
				"freq": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Float),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				freq := p.Args["freq"]
				theFreq, okFreq := freq.(float64)
				if !okFreq {
					log.Print("Invalid Frequency Value")
					os.Exit(0)
				}
				newDemo := domain.RecordDemo(demoFile, theFreq)
				log.Print("The Demo was queried using freq" + fmt.Sprintf("%f", freq))
				repository.CurrentDemo = newDemo
				return repository.CurrentDemo, nil
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
	log.Print("Created schema")

	return schema
}
