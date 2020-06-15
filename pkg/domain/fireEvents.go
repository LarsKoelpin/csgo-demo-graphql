package domain

import (
	"github.com/graphql-go/graphql"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
)

type FireStartedEvent struct {
  Position Position `json:"position"`
  Name     string   `json:"name"`
  Tick     int   `json:"tick"`
}

type FireExpiredEvent struct {
  Position Position `json:"position"`
  Name     string   `json:"name"`
  Tick     int   `json:"tick"`
}

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

		return eventName == "SMOKE_STARTED"
	},
})

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
