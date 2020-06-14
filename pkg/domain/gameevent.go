package domain

import (
	"github.com/graphql-go/graphql"
)

type GameEvent struct {
	Name  string      `json:"name"`
	Event interface{} `json:"event"`
}

var GameEventType = graphql.NewObject(graphql.ObjectConfig{
	Name: "gameEvent",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})
