package domain

import "github.com/graphql-go/graphql"

type Demo struct {
	Header Header
	Ticks  []Tick
	Events []GameEvent
}

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
				Args: nil,
				Resolve: func(resolvParams graphql.ResolveParams) (interface{}, error) {
					return repository.CurrentDemo.Events, nil
				},
			},
		},
	})
}

type DemoRepository struct {
	CurrentDemo Demo
}
