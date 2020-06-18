package domain

import (
	"github.com/graphql-go/graphql"
)

// BombPlanted represents the event, when the bomb was planted.
type BombPlanted struct {
	Name     string`json:"name"`
	Player   Player `json:"player"`
	Bombsite int32 `json:"bombsite"`
}

// BombPlantedType represents the BombPlantedEvent as GraphQL Type.
var BombPlantedType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BombPlanted",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Name: "name",
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "BOMB_PLANTED", nil
			},
		},
		"player": &graphql.Field{
			Name: "player",
			Type: PlayerType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				x := p.Source.(GameEvent)
				return x.RealEvent.(BombPlanted).Player, nil
			},
		},
		"bombsite": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				x := p.Source.(GameEvent)
				return x.RealEvent.(BombPlanted).Bombsite, nil
			},
		},
	},
	IsTypeOf: func(p graphql.IsTypeOfParams) bool {
		eventName := p.Value.(GameEvent).Name

		return eventName == "BOMB_PLANTED"
	},
})
