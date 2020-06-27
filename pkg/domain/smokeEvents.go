package domain

import (
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
)

type RenderedSmokeStarted map[string]interface{}

func RenderSmokeStarted(template map[string]interface{}, event SmokeStartedEvent) RenderedSmokeStarted {
	renderedSmokeStartedEvent := map[string]interface{}{}

	positionTemplate, hasPosition := template["position"]
	_, hasTick := template["tick"]
	_, hasName := template["name"]
	_, hasId := template["id"]

	if hasId {
		renderedSmokeStartedEvent["id"] = event.Id
	}

	if hasPosition {
		posTpl := positionTemplate.(map[string]interface{})
		renderedSmokeStartedEvent["position"] = RenderPosition(posTpl, event.Position)
	}

	if hasTick {
		renderedSmokeStartedEvent["tick"] = event.Tick
	}

	if hasName {
		renderedSmokeStartedEvent["name"] = event.Name
	}

	return renderedSmokeStartedEvent
}

func SmokeStarted(tick int, e events.SmokeStart) SmokeStartedEvent {
	return SmokeStartedEvent{
		Position: Position{
			X: e.Position.X,
			Y: e.Position.Y,
			Z: e.Position.Z,
		},
		Name: "SMOKE_STARTED",
		Tick: tick,
		Id:   e.GrenadeEntityID,
	}
}

type SmokeExpiredTemplate map[string]map[string]interface{}
type RenderedSmokeExpired map[string]interface{}

func RenderSmokeExpired(template map[string]interface{}, event SmokeExpiredEvent) RenderedSmokeExpired {
	renderedSmokeExpiredEvent := map[string]interface{}{}

	positionTemplate, hasPosition := template["position"]
	_, hasTick := template["tick"]
	_, hasName := template["name"]
	_, hasId := template["id"]

	if hasPosition {
		posTpl := positionTemplate.(map[string]interface{})
		renderedSmokeExpiredEvent["position"] = RenderPosition(posTpl, event.Position)
	}

	if hasTick {
		renderedSmokeExpiredEvent["tick"] = event.Tick
	}

	if hasName {
		renderedSmokeExpiredEvent["name"] = event.Name
	}

	if hasId {
		renderedSmokeExpiredEvent["id"] = event.Id
	}

	return renderedSmokeExpiredEvent
}

func SmokeExpired(tick int, e events.SmokeExpired) SmokeExpiredEvent {
	return SmokeExpiredEvent{
		Id: e.GrenadeEntityID,
		Position: Position{
			X: e.Position.X,
			Y: e.Position.Y,
			Z: e.Position.Z,
		},
		Name: "SMOKE_EXPIRED",
		Tick: tick,
	}
}

type SmokeStartedEvent struct {
	Id       int      `json:"id"`
	Position Position `json:"position"`
	Name     string   `json:"name"`
	Tick     int      `json:"tick"`
}

type SmokeExpiredEvent struct {
	Id       int      `json:"id"`
	Position Position `json:"position"`
	Name     string   `json:"name"`
	Tick     int      `json:"tick"`
}
