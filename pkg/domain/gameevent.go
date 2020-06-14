package domain

import (
	"github.com/graphql-go/graphql"
)

var GameEventType = graphql.NewUnion(graphql.UnionConfig{
	Name: "AnyGameEvent",
	Types: []*graphql.Object{
		BombPlantedType,
	},
})
