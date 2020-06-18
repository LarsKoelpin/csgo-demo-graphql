package domain

import (
	"github.com/graphql-go/graphql"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
)

// FireStartedEvent represents the JSON of the fire started event.
type FireStartedEvent struct {
	Position Position `json:"position"`
	Name     string   `json:"name"`
	Tick     int      `json:"tick"`
}

// FireExpiredEvent represents the JSON of the fire expired event.
type FireExpiredEvent struct {
	Position Position `json:"position"`
	Name     string   `json:"name"`
	Tick     int      `json:"tick"`
}

// FireStarted creates a GameEvent for FIRE_STARTED of the real Event.
func FireStarted(tick int, e events.FireGrenadeStart) GameEvent {
	return GameEvent{
		Name: "FIRE_STARTED",
		RealEvent: FireStartedEvent{
			Position: Position{
				X: e.Position.X,
				Y: e.Position.Y,
				Z: e.Position.Z,
			},
			Name: "FIRE_STARTED",
			Tick: tick,
		},
	}
}

// FireExpired creates a GameEvent for FIRE_EXPIRED of the real Event.
func FireExpired(tick int, e events.FireGrenadeExpired) GameEvent {
	return GameEvent{
		Name: "FIRE_EXPIRED",
		RealEvent: FireExpiredEvent{
			Position: Position{
				X: e.Position.X,
				Y: e.Position.Y,
				Z: e.Position.Z,
			},
			Name: "FIRE_EXPIRED",
			Tick: tick,
		},
	}
}

// FireStartedType is the GraphQL type of the FireStartedEvent
var FireStartedType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FireStarted",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Name: "name",
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "FIRE_STARTED", nil
			},
		},
		"position": &graphql.Field{
			Type: PositionType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				x := p.Source.(GameEvent)
				return x.RealEvent.(FireStartedEvent).Position, nil
			},
		},
		"tick": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				x := p.Source.(GameEvent)
				return x.RealEvent.(FireStartedEvent).Tick, nil
			},
		},
	},
	IsTypeOf: func(p graphql.IsTypeOfParams) bool {
		eventName := p.Value.(GameEvent).Name

		return eventName == "FIRE_STARTED"
	},
})

// FireExpiredType is the GraphQL type of the FireStartedEvent
var FireExpiredType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FireExpired",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Name: "name",
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "FIRE_EXPIRED", nil
			},
		},
		"position": &graphql.Field{
			Type: PositionType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				x := p.Source.(GameEvent)
				return x.RealEvent.(FireExpiredEvent).Position, nil
			},
		},
		"tick": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				x := p.Source.(GameEvent)
				return x.RealEvent.(FireExpiredEvent).Tick, nil
			},
		},
	},
	IsTypeOf: func(p graphql.IsTypeOfParams) bool {
		eventName := p.Value.(GameEvent).Name

		return eventName == "FIRE_EXPIRED"
	},
})
