package domain

import (
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
)

type MatchStartedEvent struct {
	Name string `json:"name"`
	Tick int    `json:"tick"`
}

func NewMatchStartedEvent(t int, e events.MatchStart) MatchStartedEvent {
	return MatchStartedEvent{
		Name: "MATCH_STARTED",
		Tick: t,
	}
}

func RenderMatchStartedEvent(template map[string]interface{}, e MatchStartedEvent) map[string]interface{} {
	result := map[string]interface{}{}
	_, hasName := template["name"]
	_, hasTick := template["tick"]

	if hasName {
		result["name"] = e.Name
	}
	if hasTick {
		result["tick"] = e.Tick
	}
	return result
}
