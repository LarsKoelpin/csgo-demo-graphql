package domain

import (
	"github.com/graphql-go/graphql"
)

type BombPlanted struct {
	Name     string
	Player   Participant
	Bombsite int32
}

var BombPlantedType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BombPlanted",
	Fields: graphql.Fields{
		"player": &graphql.Field{
			Name: "player",
			Type: ParticipantType,
		},
		"bombSite": &graphql.Field{
			Type: graphql.Int,
		},
	},
	IsTypeOf: func(p graphql.IsTypeOfParams) bool {
		return true
	},
})
