package domain

import (
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
)

func NewProjectile(e common.GrenadeProjectile, firing map[int]bool) Grenade {
	return Grenade{
		Thrower:        CreateParticipant(e.Thrower, firing[e.Thrower.EntityID]),
		Owner:          CreateParticipant(e.Owner, firing[e.Owner.EntityID]),
		Trajectory:     FromVectors(e.Trajectory),
		WeaponInstance: FromEquipment(e.WeaponInstance),
		Position:       FromVector(e.Position()),
		UniqueId:       e.UniqueID(),
	}
}

type Grenade struct {
	Thrower        Player     `json:"thrower"`
	Owner          Player     `json:"owner"`
	Trajectory     []Position `json:"trajectory"`
	WeaponInstance Equipment  `json:"weapon"`
	Position       Position   `json:"position"`
	UniqueId       int64      `json:"uniqueId"`
}
