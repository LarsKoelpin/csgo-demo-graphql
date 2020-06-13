package main

import "github.com/graphql-go/graphql"

type Demo struct {
	Header Header
	Ticks  []Tick
}

func CreateDemoType(d Demo) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "demo",
		Fields: graphql.Fields{
			"header": &graphql.Field{
				Name: "header",
				Type: HeaderType,
				Args: nil,
				Resolve: func(resolvParams graphql.ResolveParams) (interface{}, error) {
					return d.Header, nil
				},
			},
			"ticks": &graphql.Field{
				Name: "header",
				Type: graphql.NewList(TickType),
				Args: nil,
				Resolve: func(resolvParams graphql.ResolveParams) (interface{}, error) {
					return d.Ticks, nil
				},
			},
		},
	})
}
