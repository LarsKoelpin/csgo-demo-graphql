package main

import "github.com/graphql-go/graphql"

func CreateDemoType(d Demo) *graphql.Object {
  return graphql.NewObject(graphql.ObjectConfig{
    Name: "demo",
    Fields: graphql.Fields{
      "header": &graphql.Field{
        Name:              "header",
        Type:              HeaderType,
        Args:              nil,
        Resolve: func(resolvParams graphql.ResolveParams) (interface{}, error) {
          return d.Header, nil
        },
      },
    },
  })
}


type Demo struct {
  Header Header
}


type Tick struct {
  Tick  int    `json:"tick"`
  Player Player `json:"player"`
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

type Position struct {
  x float32
  y float32
  z float32
}

type Event struct {

}

type EntityEquipment struct {
  Type           int `json:"type" msgpack:"type"`
  AmmoInMagazine int `json:"ammoInMagazine" msgpack:"ammoInMagazine"`
  AmmoReserve    int `json:"ammoReserve" msgpack:"ammoReserve"`
}
