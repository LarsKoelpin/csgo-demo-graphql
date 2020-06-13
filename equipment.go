package main

import (
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
)

type EntityEquipment struct {
	Type           int `json:"type"`
	AmmoInMagazine int `json:"ammoInMagazine"`
	AmmoReserve    int `json:"ammoReserve"`
}

func toEntityEquipment(eq []*common.Equipment) []EntityEquipment {
	var equipmentForPlayer = make([]EntityEquipment, 0, len(eq))

	for _, equipment := range eq {
		equipmentForPlayer = append(equipmentForPlayer, EntityEquipment{
			Type:           int(equipment.Type),
			AmmoInMagazine: equipment.AmmoInMagazine(),
			AmmoReserve:    equipment.AmmoReserve(),
		})
	}

	return equipmentForPlayer
}
