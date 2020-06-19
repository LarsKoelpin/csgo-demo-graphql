package domain

import "github.com/graphql-go/graphql"

// Smoke represents an active smoke in the world.
type Smoke struct {
	Id       int      `json:"id"`
	Position Position `json:"position"`
}

// SmokeType represents the GraphQL Type of a Smoke
var SmokeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "smoke",
	Fields: graphql.Fields{
		"position": &graphql.Field{
			Type: PositionType,
		},
		"id": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

// Remove removes a Smoke Id from smokes
func Remove(smokes []Smoke, toRemove int) []Smoke {
	result := make([]Smoke, 0)
	for _, single := range smokes {
		if single.Id != toRemove {
			result = append(result, single)
		}
	}

	return result
}
