package domain

import (
	"github.com/graphql-go/graphql"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
)

type RoundStartedEvent struct {
	Name      string `json:"name"`
	Tick      int    `json:"tick"`
	FragLimit int    `json:"fragLimit"`
	Objective string `json:"objective"`
	TimeLimit int    `json:"timeLimit"`
}

type RoundEndEvent struct {
	Name string `json:"name"`
	Tick int    `json:"tick"`
}

func RoundStarted(tick int, e events.RoundStart) GameEvent {
	return GameEvent{
		Name: "ROUND_STARTED",
		RealEvent: RoundStartedEvent{
			Name:      "ROUND_STARTED",
			Tick:      tick,
			FragLimit: e.FragLimit,
			Objective: e.Objective,
			TimeLimit: e.TimeLimit,
		},
	}
}

func RoundEnded(tick int, e events.RoundEnd) GameEvent {
	return GameEvent{
		Name: "ROUND_ENDED",
		RealEvent: RoundEndEvent{
			Name: "ROUND_ENDED",
			Tick: tick,
		},
	}
}

var RoundStartedType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RoundStarted",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Name: "name",
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "ROUND_STARTED", nil
			},
		},
		"tick": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				x := p.Source.(GameEvent)
				return x.RealEvent.(RoundStartedEvent).Tick, nil
			},
		},
		"timeLimit": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				x := p.Source.(GameEvent)
				return x.RealEvent.(RoundStartedEvent).TimeLimit, nil
			},
		},
		"objective": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				x := p.Source.(GameEvent)
				return x.RealEvent.(RoundStartedEvent).Objective, nil
			},
		},
		"fragLimit": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				x := p.Source.(GameEvent)
				return x.RealEvent.(RoundStartedEvent).FragLimit, nil
			},
		},
	},
	IsTypeOf: func(p graphql.IsTypeOfParams) bool {
		eventName := p.Value.(GameEvent).Name

		return eventName == "ROUND_STARTED"
	},
})

var RoundEndedType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RoundEnded",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Name: "name",
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "ROUND_ENDED", nil
			},
		},
		"tick": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				x := p.Source.(GameEvent)
				return x.RealEvent.(RoundEndEvent).Tick, nil
			},
		},
	},
	IsTypeOf: func(p graphql.IsTypeOfParams) bool {
		eventName := p.Value.(GameEvent).Name

		return eventName == "ROUND_ENDED"
	},
})
