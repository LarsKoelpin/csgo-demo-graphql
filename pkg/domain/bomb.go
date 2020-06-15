package domain

import "github.com/graphql-go/graphql"

type Bomb struct {
	LastOnGroundPosition Position `json:"lastOnGroundPosition"`
	Carrier              Player   `json:"carrier"`
}

var BombType = graphql.NewObject(graphql.ObjectConfig{
	Name: "bomb",
	Fields: graphql.Fields{
		"lastOnGroundPosition": &graphql.Field{
			Type: PositionType,
		},
		"carrier": &graphql.Field{
			Type: ParticipantType,
		},
	},
})
