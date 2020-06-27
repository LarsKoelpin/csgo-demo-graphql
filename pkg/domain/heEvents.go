package domain

import (
	"github.com/graphql-go/graphql"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
)

// HEExplosionEvent represents an event of a explodion HE-Grenade.
type HEExplosionEvent struct {
	Position Position `json:"position"`
	Name     string   `json:"name"`
	Tick     int      `json:"tick"`
}

func RenderHEExplosionEvent(template map[string]interface{}, e HEExplosionEvent) map[string]interface{} {
	result := map[string]interface{}{}
	_, hasName := template["name"]
	_, hasTick := template["tick"]
	posTemplate, hasPosition := template["position"]

	if hasName {
		result["name"] = e.Name
	}

	if hasTick {
		result["tick"] = e.Tick
	}

	if hasPosition {
		positionTemplate := posTemplate.(map[string]interface{})
		result["position"] = RenderPosition(positionTemplate, e.Position)
	}

	return result
}

// NewHeExplision creates a new HEExplosion Event from the core event.
func NewHeExplosion(tick int, e events.HeExplode) HEExplosionEvent {
	return HEExplosionEvent{
		Position: Position{
			X: e.Position.X,
			Y: e.Position.Y,
			Z: e.Position.Z,
		},
		Name: "HE_EXPLOSION",
		Tick: tick,
	}
}

// HEExplosionType represents the GraphQL Type of a HEGrenade.
var HEExplosionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "HEExploded",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Name: "name",
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "HE_EXPLOSION", nil
			},
		},
		"position": &graphql.Field{
			Type: PositionType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				x := p.Source.(GameEvent)
				return x.RealEvent.(HEExplosionEvent).Position, nil
			},
		},
		"tick": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				x := p.Source.(GameEvent)
				return x.RealEvent.(HEExplosionEvent).Tick, nil
			},
		},
	},
	IsTypeOf: func(p graphql.IsTypeOfParams) bool {
		eventName := p.Value.(GameEvent).Name

		return eventName == "HE_EXPLOSION"
	},
})
