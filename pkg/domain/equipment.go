package domain

import (
	"github.com/graphql-go/graphql"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
)

type Equipment struct {
	Type           int `json:"type"`
	AmmoInMagazine int `json:"ammoInMagazine"`
	AmmoReserve    int `json:"ammoReserve"`
}

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

func ToEntityEquipment(eq []*common.Equipment) []Equipment {
	var equipmentForPlayer = make([]Equipment, 0, len(eq))

	for _, equipment := range eq {
		equipmentForPlayer = append(equipmentForPlayer, Equipment{
			Type:           int(equipment.Type),
			AmmoInMagazine: equipment.AmmoInMagazine(),
			AmmoReserve:    equipment.AmmoReserve(),
		})
	}

	return equipmentForPlayer
}
