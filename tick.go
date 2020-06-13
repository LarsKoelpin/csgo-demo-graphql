package main

import "github.com/graphql-go/graphql"

var TickType = graphql.NewObject(graphql.ObjectConfig{
  Name: "ticks",
  Fields: graphql.Fields{
    "tick": &graphql.Field{
      Type: graphql.Int,
    },
    "players": &graphql.Field{
      Type: graphql.NewList(PlayerType),
    },
  },
})


var PlayerType = graphql.NewObject(graphql.ObjectConfig{
  Name: "players",
  Fields: graphql.Fields{
    "entityId": &graphql.Field{
      Type: graphql.Int,
    },
    "team": &graphql.Field{
      Type: graphql.Int,
    },
    "position": &graphql.Field{
      Type: PositioType,
    },
  },
})


type Tick struct {
  Tick   int      `json:"tick"`
  Players []Player `json:"players"`
}
type Player struct {
  EntityID      int               `json:"entityId"`
  Team          int               `json:"team,omitempty"`
  Positions     Position          `json:"position,omitempty"` // This allows us smoother replay with less overhead compared to higher snapshot rate
  AngleX        int               `json:"angleX,omitempty"`
  AngleY        int               `json:"angleY,omitempty"`
  Hp            int               `json:"hp,omitempty"`
  Armor         int               `json:"armor,omitempty"`
  FlashDuration float32           `json:"flashDuration,omitempty"`
  IsNpc         bool              `json:"isNpc,omitempty"`
  HasHelmet     bool              `json:"hasHelmet,omitempty"`
  HasDefuseKit  bool              `json:"hasDefuseKit,omitempty"`
  Equipment     []EntityEquipment `json:"equipment,omitempty"`
  IsPlanting    bool              `json:"isPlanting,omitempty"`
  IsDefusing    bool              `json:"isDefusing,omitempty"`
  IsInBuyzone   bool              `json:"isInBuyzone,omitempty"`
  Money         int               `json:"money,omitempty"`
  Kills         int               `json:"kills,omitempty"`
  Deaths        int               `json:"deaths,omitempty"`
}
