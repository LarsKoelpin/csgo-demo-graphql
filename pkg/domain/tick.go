package domain

import "github.com/graphql-go/graphql"

type Tick struct {
	Tick              int           `json:"tick"`
	Participants      []Participant `json:"participants"`
	Bomb              Bomb          `json:"bomb"`
	TotalRoundsPlayed int           `json:"totalRoundsPlayed"`
}

var TickType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ticks",
	Fields: graphql.Fields{
		"tick": &graphql.Field{
			Type: graphql.Int,
		},
		"participants": &graphql.Field{
			Type: graphql.NewList(ParticipantType),
		},
		"bomb": &graphql.Field{
			Type: BombType,
		},
		"totalRoundsPlayed": &graphql.Field{
			Type: graphql.Int,
		},
	},
})
