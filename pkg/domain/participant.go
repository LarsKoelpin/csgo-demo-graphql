package domain

import (
	"github.com/graphql-go/graphql"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
)

var ParticipantType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Participant",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"entityId": &graphql.Field{
			Type: graphql.Int,
		},
		"team": &graphql.Field{
			Type: graphql.Int,
		},
		"position": &graphql.Field{
			Type: PositionType,
		},
		"angleX": &graphql.Field{
			Type: graphql.Float,
		},
		"angleY": &graphql.Field{
			Type: graphql.Float,
		},
		"hp": &graphql.Field{
			Type: graphql.Int,
		},
		"armor": &graphql.Field{
			Type: graphql.Int,
		},
		"flashDuration": &graphql.Field{
			Type: graphql.Float,
		},
		"isNpc": &graphql.Field{
			Type: graphql.Boolean,
		},
		"hasHelmet": &graphql.Field{
			Type: graphql.Boolean,
		},
		"hasDefuseKit": &graphql.Field{
			Type: graphql.Boolean,
		},
		"equipment": &graphql.Field{
			Type: graphql.NewList(EquipmentType),
		},
		"isPlanting": &graphql.Field{
			Type: graphql.Boolean,
		},
		"isDefusing": &graphql.Field{
			Type: graphql.Boolean,
		},
		"isInBuyzone": &graphql.Field{
			Type: graphql.Boolean,
		},
		"money": &graphql.Field{
			Type: graphql.Int,
		},
		"kills": &graphql.Field{
			Type: graphql.Int,
		},
		"deaths": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

type Participant struct {
	Name          string      `json:"name"`
	EntityID      int         `json:"entityId"`
	Team          int         `json:"team,omitempty"`
	Position      Position    `json:"position,omitempty"` // This allows us smoother replay with less overhead compared to higher snapshot rate
	AngleX        int         `json:"angleX,omitempty"`
	AngleY        int         `json:"angleY,omitempty"`
	Hp            int         `json:"hp,omitempty"`
	Armor         int         `json:"armor,omitempty"`
	FlashDuration float32     `json:"flashDuration,omitempty"`
	IsNpc         bool        `json:"isNpc,omitempty"`
	HasHelmet     bool        `json:"hasHelmet,omitempty"`
	HasDefuseKit  bool        `json:"hasDefuseKit,omitempty"`
	Equipment     []Equipment `json:"equipment,omitempty"`
	IsPlanting    bool        `json:"isPlanting,omitempty"`
	IsDefusing    bool        `json:"isDefusing,omitempty"`
	IsInBuyzone   bool        `json:"isInBuyzone,omitempty"`
	Money         int         `json:"money,omitempty"`
	Kills         int         `json:"kills,omitempty"`
	Deaths        int         `json:"deaths,omitempty"`
}

func CreateParticipant(pl *common.Player) Participant {
	return Participant{
		Name:          pl.Name,
		EntityID:      pl.EntityID,
		Hp:            pl.Health(),
		Armor:         pl.Armor(),
		FlashDuration: 0.1, // Round to nearest 0.1 sec - saves space in JSON
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
		IsDefusing:   pl.IsDefusing,
		IsPlanting:   pl.IsPlanting,
		Money:        pl.Money(),
		Kills:        pl.Kills(),
		Deaths:       pl.Deaths(),
		IsInBuyzone:  pl.IsInBuyZone(),
	}
}
