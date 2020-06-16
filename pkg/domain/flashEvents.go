package domain

import (
	"github.com/graphql-go/graphql"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
)

type FlashExplosionEvent struct {
	Position Position `json:"position"`
	Name     string   `json:"name"`
	Tick     int      `json:"tick"`
}

func NewFlashExplosion(tick int, e events.FlashExplode) GameEvent {
	return GameEvent{
		Name: "FLASH_EXPLOSION",
		RealEvent: SmokeStartedEvent{
			Position: Position{
				X: e.Position.X,
				Y: e.Position.Y,
				Z: e.Position.Z,
			},
			Name: "FLASH_EXPLOSION",
			Tick: tick,
		},
	}
}

var FlashExplosionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FlashExplosion",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Name: "name",
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "FLASH_EXPLOSION", nil
			},
		},
		"position": &graphql.Field{
			Type: PositionType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				x := p.Source.(GameEvent)
				return x.RealEvent.(SmokeStartedEvent).Position, nil
			},
		},
		"tick": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				x := p.Source.(GameEvent)
				return x.RealEvent.(SmokeStartedEvent).Tick, nil
			},
		},
	},
	IsTypeOf: func(p graphql.IsTypeOfParams) bool {
		eventName := p.Value.(GameEvent).Name

		return eventName == "FLASH_EXPLOSION"
	},
})
