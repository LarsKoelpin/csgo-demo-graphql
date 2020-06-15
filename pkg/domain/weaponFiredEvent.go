package domain

import (
	"github.com/graphql-go/graphql"
)

type WeaponFired struct {
	Name    string    `json:"name"`
	Shooter Player    `json:"shooter"`
	Weapon  Equipment `json:"weapon"`
}

var WeaponFiredType = graphql.NewObject(graphql.ObjectConfig{
	Name: "WeaponFired",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Name: "name",
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "WEAPON_FIRED", nil
			},
		},
		"shooter": &graphql.Field{
			Name: "shooter",
			Type: ParticipantType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				x := p.Source.(GameEvent)
				event := x.RealEvent.(WeaponFired)
				return event.Shooter, nil
			},
		},
		"weapon": &graphql.Field{
			Name: "weapon",
			Type: EquipmentType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				x := p.Source.(GameEvent)
				event := x.RealEvent.(WeaponFired)
				return event.Weapon, nil
			},
		},
	},
	IsTypeOf: func(p graphql.IsTypeOfParams) bool {
		eventName := p.Value.(GameEvent).Name

		return eventName == "WEAPON_FIRED"
	},
})
