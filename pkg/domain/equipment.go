package domain

import (
	"github.com/graphql-go/graphql"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
)

// Equipment represents the Equipment of a player in the json format.
type Equipment struct {
	Type           int `json:"type"`
	AmmoInMagazine int `json:"ammoInMagazine"`
	AmmoReserve    int `json:"ammoReserve"`
	AmmoType       int `json:"ammoType"`
}

// EquipmentType represents the Equipment GraphQL Type.
var EquipmentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "equipment",
	Fields: graphql.Fields{
		"type": &graphql.Field{
			Type: graphql.Int,
		},
		"ammoInMagazine": &graphql.Field{
			Type: graphql.Int,
		},
		"ammoReserve": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

// ToEntityEquipment maps an array of "core" Equipment to the JSON Model.
func ToEntityEquipment(eq []*common.Equipment) []Equipment {
	var equipmentForPlayer = make([]Equipment, 0, len(eq))

	for _, equipment := range eq {
		equipmentForPlayer = append(equipmentForPlayer, FromEquipment(equipment))
	}

	return equipmentForPlayer
}

// FromEquipment maps a single "core" Equipment to the JSON Model.
func FromEquipment(eq *common.Equipment) Equipment {
	return Equipment{
		Type:           int(eq.Type),
		AmmoInMagazine: eq.AmmoInMagazine(),
		AmmoReserve:    eq.AmmoReserve(),
		AmmoType:       eq.AmmoType(),
	}
}
