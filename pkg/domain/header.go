package domain

import "github.com/graphql-go/graphql"

type Header struct {
	MapName  string  `json:"map"`
	TickRate float64 `json:"tickRate"` // How many ticks per second
	Fps      int     `json:"fps"`
}

var HeaderType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Header",
	Fields: graphql.Fields{
		"mapName": &graphql.Field{
			Type: graphql.String,
		},
		"tickRate": &graphql.Field{
			Type: graphql.Int,
		},
		"fps": &graphql.Field{
			Type: graphql.Int,
		},
	},
})
