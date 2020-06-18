package domain

import (
	"github.com/graphql-go/graphql"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
)

// Inferno represents a burning fire.
type Inferno struct {
	Hull []Position `json:"hull"`
}

// InfernoType represents the GraphQL Type of an inferno.
var InfernoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Inferno",
	Fields: graphql.Fields{
		"hull": &graphql.Field{
			Type: graphql.NewList(PositionType),
		},
	},
})

func ToInferno(inferno common.Inferno) Inferno {
	return Inferno{
		Hull: FromPoints(inferno.Fires().ConvexHull2D()),
	}
}
