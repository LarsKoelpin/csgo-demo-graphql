package domain

import (
	"log"
	"math"

	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
)

type PlayerTemplate map[string]interface{}
type RenderedPlayer map[string]interface{}

func RenderPlayer(template PlayerTemplate, player Player) RenderedPlayer {
	renderedPlayer := map[string]interface{}{}

	_, hasName := template["name"]
	_, hasEntityId := template["entityId"]
	_, hasTeam := template["team"]
	posTemplate, hasPosition := template["position"]
	_, hasAngleX := template["angleX"]
	_, hasAngleY := template["angleY"]
	_, hasHp := template["hp"]
	_, hasArmor := template["armor"]
	_, hasFlashDuration := template["flashDuration"]
	_, hasNpc := template["npc"]
	_, hasHelmet := template["hasHelmet"]
	_, hasDefuse := template["hasDefuseKit"]
	equipmentTpl, equipment := template["equipment"]
	_, hasPlanting := template["planting"]
	_, hasDefusing := template["defusing"]
	_, hasBuyzone := template["inBuyzone"]
	_, hasMoney := template["money"]
	_, hasKills := template["kills"]
	_, hasDeaths := template["deaths"]
	_, hasFiring := template["firing"]

	if hasName {
		renderedPlayer["name"] = player.Name
	}

	if hasEntityId {
		renderedPlayer["entityId"] = player.EntityID
	}
	if hasTeam {
		renderedPlayer["team"] = player.Team
	}

	if hasPosition {
		positionTemplate, ok := posTemplate.(map[string]interface{})
		if ok {
			renderedPlayer["position"] = RenderPosition(positionTemplate, player.Position)
		} else {
			log.Panic("WTF")
		}
	}

	if hasAngleX {
		renderedPlayer["angleX"] = player.AngleX
	}

	if hasAngleY {
		renderedPlayer["angleY"] = player.AngleY
	}

	if hasHp {
		renderedPlayer["hp"] = player.Hp
	}

	if hasArmor {
		renderedPlayer["armor"] = player.HasHelmet
	}

	if hasFlashDuration {
		renderedPlayer["flashDuration"] = player.FlashDuration
	}

	if hasNpc {
		renderedPlayer["npc"] = player.Npc
	}

	if hasHelmet {
		renderedPlayer["hasHelmet"] = player.HasHelmet
	}

	if hasDefuse {
		renderedPlayer["hasDefuseKit"] = player.HasDefuseKit
	}

	if hasPlanting {
		renderedPlayer["planting"] = player.Planting
	}

	if hasDefusing {
		renderedPlayer["defusing"] = player.Defusing
	}

	if hasBuyzone {
		renderedPlayer["inBuyzone"] = player.InBuyzone
	}

	if hasMoney {
		renderedPlayer["money"] = player.Money
	}

	if hasKills {
		renderedPlayer["kills"] = player.Kills
	}

	if hasDeaths {
		renderedPlayer["deaths"] = player.Deaths
	}

	if hasFiring {
		renderedPlayer["firing"] = player.Firing
	}

	if equipment {
		castedEquipmentTpl := equipmentTpl.(map[string]interface{})
		equipmentList := make([]map[string]interface{}, len(player.Equipment))
		for _, singleEq := range player.Equipment {
			equipmentList = append(equipmentList, RenderEquipment(castedEquipmentTpl, singleEq))
		}
	}

	return renderedPlayer
}

type Player struct {
	Name          string      `json:"name"`
	EntityID      int         `json:"entityId"`
	Team          int         `json:"team,omitempty"`
	Position      Position    `json:"position,omitempty"` // This allows us smoother replay with less overhead compared to higher snapshot rate
	AngleX        int         `json:"angleX,omitempty"`
	AngleY        int         `json:"angleY,omitempty"`
	Hp            int         `json:"hp,omitempty"`
	Armor         int         `json:"armor,omitempty"`
	FlashDuration float32     `json:"flashDuration,omitempty"`
	Npc           bool        `json:"npc,omitempty"`
	HasHelmet     bool        `json:"hasHelmet,omitempty"`
	HasDefuseKit  bool        `json:"hasDefuseKit,omitempty"`
	Equipment     []Equipment `json:"equipment,omitempty"`
	Planting      bool        `json:"planting,omitempty"`
	Defusing      bool        `json:"defusing,omitempty"`
	InBuyzone     bool        `json:"inBuyzone,omitempty"`
	Money         int         `json:"money,omitempty"`
	Kills         int         `json:"kills,omitempty"`
	Deaths        int         `json:"deaths,omitempty"`
	Firing        bool        `json:"isFiring, omitempty"`
}

func NewPlayer(pl *common.Player, fireing bool) Player {
	return Player{
		Name:          pl.Name,
		EntityID:      pl.EntityID,
		Hp:            pl.Health(),
		Armor:         pl.Armor(),
		FlashDuration: float32(math.Round(float64(pl.FlashDuration*100)) / 100),
		Position: Position{
			X: pl.Position().X,
			Y: pl.Position().Y,
			Z: pl.Position().Z,
		},
		AngleX:       int(pl.ViewDirectionX()),
		AngleY:       int(pl.ViewDirectionY()),
		HasHelmet:    pl.HasHelmet(),
		HasDefuseKit: pl.HasDefuseKit(),
		Equipment:    ToEntityEquipment(pl.Weapons()),
		Team:         int(pl.Team),
		Defusing:     pl.IsDefusing,
		Planting:     pl.IsPlanting,
		Money:        pl.Money(),
		Kills:        pl.Kills(),
		Deaths:       pl.Deaths(),
		InBuyzone:    pl.IsInBuyZone(),
		Firing:       fireing,
	}
}
