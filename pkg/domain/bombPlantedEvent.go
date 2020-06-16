package domain

import (
	"github.com/graphql-go/graphql"
)

type BombPlanted struct {
	Name     string
	Player   Player
	Bombsite int32 `json:"bombsite"`
}

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
