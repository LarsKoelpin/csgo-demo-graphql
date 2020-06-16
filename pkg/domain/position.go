package domain

import (
	"github.com/golang/geo/r3"
	"github.com/graphql-go/graphql"
)

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

var PositionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Position",
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

func FromVector(e r3.Vector) Position {
	return Position{
		X: e.X,
		Y: e.Y,
		Z: e.Z,
	}
}

func FromVectors(e []r3.Vector) []Position {
	result := make([]Position, len(e))
	for _, val := range e {
		result = append(result, Position{
			X: val.X,
			Y: val.Y,
			Z: val.Z,
		})
	}
	return result
}
