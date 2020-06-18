package domain

import (
	"math"

	"github.com/graphql-go/graphql"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
)

var PlayerType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Player",
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
		"npc": &graphql.Field{
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
		"planting": &graphql.Field{
			Type: graphql.Boolean,
		},
		"defusing": &graphql.Field{
			Type: graphql.Boolean,
		},
		"inBuyzone": &graphql.Field{
			Type: graphql.Boolean,
		},
		"firing": &graphql.Field{
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

func CreateParticipant(pl *common.Player, fireing bool) Player {
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
