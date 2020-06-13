package domain

import "github.com/graphql-go/graphql"

var GrenadeProjectileType = graphql.NewObject(graphql.ObjectConfig{
	Name: "grenadeProjectTile",
	Fields: graphql.Fields{
		"thrower": &graphql.Field{
			Type: ParticipantType,
		},
		"owner": &graphql.Field{
			Type: ParticipantType,
		},
		"trajectory": &graphql.Field{
			Type: graphql.NewList(PositionType),
		},
		"weapon": &graphql.Field{
			Type: EquipmentType,
		},
	},
})

type GrenadeProjectile struct {
	Thrower        Participant `json:"thrower"`
	Owner          Participant `json:"owner"`
	Trajectory     []Position  `json:"trajectory"`
	WeaponInstance Equipment   `json:"weapon"`
	Position       Position    `json:"position"`
	UniqueId       int64       `json:"uniqueId"`
}
