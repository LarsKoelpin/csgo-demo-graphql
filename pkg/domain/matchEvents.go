package domain

import (
	"github.com/graphql-go/graphql"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
)

type MatchStartedEvent struct {
	Name string `json:"name"`
	Tick int    `json:"tick"`
}

func NewMatchStartedEvent(t int, e events.MatchStart) GameEvent {
	return GameEvent{
		Name: "MATCH_STARTED",
		RealEvent: MatchStartedEvent{
			Name: "MATCH_STARTED",
			Tick: t,
		},
	}
}

var MatchStartedEventTpe = graphql.NewObject(graphql.ObjectConfig{
	Name: "MatchStarted",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Name: "name",
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "MATCH_STARTED", nil
			},
		},
		"tick": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				x := p.Source.(GameEvent)
				return x.RealEvent.(MatchStartedEvent).Tick, nil
			},
		},
	},
	IsTypeOf: func(p graphql.IsTypeOfParams) bool {
		eventName := p.Value.(GameEvent).Name

		return eventName == "MATCH_STARTED"
	},
})
