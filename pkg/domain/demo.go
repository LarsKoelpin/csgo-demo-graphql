package domain

import (
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

// Demo represents the json of a whole demo.
type Demo struct {
	Header Header      `json:"header"`
	Ticks  []Tick      `json:"ticks"`
	Events []GameEvent `json:"events"`
}

// CreateDemoType creats a GraphQL Schema representing the whole query.
func CreateDemoType(repository *DemoRepository) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "demo",
		Fields: graphql.Fields{
			"header": &graphql.Field{
				Name: "header",
				Type: HeaderType,
				Args: nil,
				Resolve: func(resolvParams graphql.ResolveParams) (interface{}, error) {
					return repository.CurrentDemo.Header, nil
				},
			},
			"ticks": &graphql.Field{
				Name: "ticks",
				Type: graphql.NewList(TickType),
				Args: nil,
				Resolve: func(resolvParams graphql.ResolveParams) (interface{}, error) {
					return repository.CurrentDemo.Ticks, nil
				},
			},
			"events": &graphql.Field{
				Name: "events",
				Type: graphql.NewList(GameEventType),
				Args: graphql.FieldConfigArgument{
					"type": &graphql.ArgumentConfig{
						Type: graphql.NewList(graphql.String),
					},
				},
				Resolve: func(resolvParams graphql.ResolveParams) (interface{}, error) {
					demoFile := resolvParams.Args["type"]

					if demoFile == nil {
						log.Print("No Event filters. Returning no events.")
						return []GameEvent{}, nil
					}

					x := demoFile.([]interface{})
					setOfStrings := make(map[string]bool)
					for _, v := range x {
						setOfStrings[fmt.Sprint(v)] = true
					}

					filteredEvents := make([]interface{}, 0)

					for _, val := range repository.CurrentDemo.Events {
						_, exists := setOfStrings[val.Name]
						if exists {
							filteredEvents = append(filteredEvents, val)
						}
					}

					return filteredEvents, nil
				},
			},
		},
	})
}

// DemoRepository holds the state of the whole Operation. In this case a simple in memory database
type DemoRepository struct {
	CurrentDemo Demo
}
