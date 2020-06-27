package domain

import (
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
)

// Equipment represents the Equipment of a player in the json format.
type Equipment struct {
	Type           int `json:"type"`
	AmmoInMagazine int `json:"ammoInMagazine"`
	AmmoReserve    int `json:"ammoReserve"`
	AmmoType       int `json:"ammoType"`
}

func RenderEquipment(template map[string]interface{}, eq Equipment) map[string]interface{} {
	result := map[string]interface{}{}
	_, hasType := template["type"]
	_, hasAmmoInMagazine := template["ammoInMagazine"]
	_, hasAmmoReserve := template["ammoReserve"]
	_, hasAmmoType := template["ammoType"]

	if hasType {
		result["type"] = eq.Type
	}

	if hasAmmoInMagazine {
		result["ammoInMagazine"] = eq.AmmoInMagazine
	}

	if hasAmmoReserve {
		result["ammoReserve"] = eq.AmmoReserve
	}

	if hasAmmoType {
		result["ammoType"] = eq.AmmoType
	}

	return result
}

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
