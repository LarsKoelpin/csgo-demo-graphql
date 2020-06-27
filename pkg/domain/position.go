package domain

import (
	"github.com/golang/geo/r2"
	"github.com/golang/geo/r3"
	"github.com/graphql-go/graphql"
)

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type PositionTemplate map[string]interface{}
type RenderedPosition map[string]interface{}

func RenderPosition(template PositionTemplate, position Position) RenderedPosition {
	result := map[string]interface{}{}
	_, hasX := template["x"]
	if hasX {
		result["x"] = position.X
	}
	_, hasY := template["y"]
	if hasY {
		result["y"] = position.Y
	}
	_, hasZ := template["z"]
	if hasZ {
		result["z"] = position.Z
	}

	return result
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

// FromPoints creats a position array from a point array.
func FromPoints(e []r2.Point) []Position {
	result := make([]Position, 0)
	for _, x := range e {
		result = append(result, FromPoint(x))
	}

	return result
}

// FromPoints creats a position array from a point array.
func FromPoint(e r2.Point) Position {
	return Position{
		X: e.X,
		Y: e.Y,
		Z: 0,
	}
}

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
