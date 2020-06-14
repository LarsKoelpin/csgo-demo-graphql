package domain

import "github.com/graphql-go/graphql"

type Header struct {
	MapName      string  `json:"map"`
	TickRate     float64 `json:"tickRate"`     // How many ticks per second
	SnapshotRate int     `json:"snapshotRate"` // How many ticks per snapshot
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
		"snapshotRate": &graphql.Field{
			Type: graphql.Float,
		},
	},
})
