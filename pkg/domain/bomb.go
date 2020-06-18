package domain

import "github.com/graphql-go/graphql"

// Bomb is the structure representing the json of the bombstate in a tick.
type Bomb struct {
	LastOnGroundPosition Position `json:"lastOnGroundPosition"`
	Carrier              Player   `json:"carrier"`
}

// BombType is the graphql implementation for the bomb.
var BombType = graphql.NewObject(graphql.ObjectConfig{
	Name: "bomb",
	Fields: graphql.Fields{
		"lastOnGroundPosition": &graphql.Field{
			Type: PositionType,
		},
		"carrier": &graphql.Field{
			Type: PlayerType,
		},
	},
})
