package domain

import "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"

type WeaponFired struct {
	Name    string    `json:"name"`
	Shooter Player    `json:"shooter"`
	Weapon  Equipment `json:"weapon"`
	Tick    int       `json:"tick"`
}

// NewWeaponFired creates an Event
func NewWeaponFired(tick int, e events.WeaponFire) WeaponFired {
	return WeaponFired{
		Name:   "WEAPON_FIRED",
		Tick:   tick,
		Weapon: FromEquipment(e.Weapon),
	}
}

func RenderWeaponFired(template map[string]interface{}, e WeaponFired) map[string]interface{} {
	result := map[string]interface{}{}
	_, hasName := template["name"]
	_, hasTick := template["tick"]
	weaponTpl, hasWeapon := template["weapon"]

	if hasName {
		result["name"] = e.Name
	}

	if hasTick {
		result["tick"] = e.Tick
	}

	if hasWeapon {
		weaponTemplate := weaponTpl.(map[string]interface{})
		result["weapon"] = RenderEquipment(weaponTemplate, e.Weapon)
	}

	return result
}
