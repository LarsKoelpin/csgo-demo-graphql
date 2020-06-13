package main

import "github.com/graphql-go/graphql"

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

var PositionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "position",
	Fields: graphql.Fields{
		"x": &graphql.Field{
			Type: graphql.Float,
		},
		"y": &graphql.Field{
			Type: graphql.Float,
		},
		"z": &graphql.Field{
			Type: graphql.Float,
		},
	},
})
