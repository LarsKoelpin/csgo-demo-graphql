package domain

import "github.com/graphql-go/graphql"

type Tick struct {
	Tick              int       `json:"tick"`
	Players           []Player  `json:"players"`
	Grenades          []Grenade `json:"grenades"`
	Bomb              Bomb      `json:"bomb"`
	TotalRoundsPlayed int       `json:"totalRoundsPlayed"`
}

var TickType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Tick",
	Fields: graphql.Fields{
		"tick": &graphql.Field{
			Type: graphql.Int,
		},
		"players": &graphql.Field{
			Type: graphql.NewList(PlayerType),
		},
		"bomb": &graphql.Field{
			Type: BombType,
		},
		"grenades": &graphql.Field{
			Type: graphql.NewList(GrenadeProjectileType),
		},
		"totalRoundsPlayed": &graphql.Field{
			Type: graphql.Int,
		},
	},
})
