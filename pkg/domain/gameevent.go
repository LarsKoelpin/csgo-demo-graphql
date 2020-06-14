package domain

import (
	"github.com/graphql-go/graphql"
)

type GameEvent struct {
	Name      string `json:"name"`
	RealEvent interface{}
}

var GameEventType = graphql.NewUnion(graphql.UnionConfig{
	Name: "AnyGameEvent",
	Types: []*graphql.Object{
		BombPlantedType,
		WeaponFiredType,
	},
})
