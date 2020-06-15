package domain

import (
	"github.com/graphql-go/graphql"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
)

func SmokeStarted(tick int, e events.SmokeStart) GameEvent {
	return GameEvent{
		Name: "SMOKE_STARTED",
		RealEvent: SmokeStartedEvent{
			Position: Position{
				X: e.Position.X,
				Y: e.Position.Y,
				Z: e.Position.Z,
			},
			Name: "SMOKE_STARTED",
			Tick: tick,
		},
	}
}

func SmokeExpired(tick int, e events.SmokeExpired) GameEvent {
	return GameEvent{
		Name: "SMOKE_EXPIRED",
		RealEvent: SmokeExpiredEvent{
			Position: Position{
				X: e.Position.X,
				Y: e.Position.Y,
				Z: e.Position.Z,
			},
			Name: "SMOKE_EXPIRED",
			Tick: tick,
		},
	}
}

type SmokeStartedEvent struct {
	Position Position `json:"position"`
	Name     string   `json:"name"`
	Tick     int      `json:"tick"`
}

type SmokeExpiredEvent struct {
	Position Position `json:"position"`
	Name     string   `json:"name"`
	Tick     int      `json:"tick"`
}

var SmokeStartedType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SmokeStarted",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Name: "name",
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "SMOKE_STARTED", nil
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

		return eventName == "SMOKE_STARTED"
	},
})

var SmokeExpiredType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SmokeExpired",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Name: "name",
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "SMOKE_EXPIRED", nil
			},
		},
		"position": &graphql.Field{
			Type: PositionType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				x := p.Source.(GameEvent)
				return x.RealEvent.(SmokeExpiredEvent).Position, nil
			},
		},
		"tick": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				x := p.Source.(GameEvent)
				return x.RealEvent.(SmokeExpiredEvent).Tick, nil
			},
		},
	},
	IsTypeOf: func(p graphql.IsTypeOfParams) bool {
		eventName := p.Value.(GameEvent).Name

		return eventName == "SMOKE_EXPIRED"
	},
})
