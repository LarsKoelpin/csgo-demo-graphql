package domain

import (
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
)

// FireStartedEvent represents the JSON of the fire started event.
type FireStartedEvent struct {
	Position Position `json:"position"`
	Name     string   `json:"name"`
	Tick     int      `json:"tick"`
	Id       int      `json:"id"`
}

// FireExpiredEvent represents the JSON of the fire expired event.
type FireExpiredEvent struct {
	Position Position `json:"position"`
	Name     string   `json:"name"`
	Tick     int      `json:"tick"`
	Id       int      `json:"id"`
}

type RenderedFireStarted map[string]interface{}
type RenderedFireExpired map[string]interface{}

func RenderFireStarted(template map[string]interface{}, event FireStartedEvent) RenderedFireStarted {
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

// FireStarted creates a GameEvent for FIRE_STARTED of the real Event.
func FireStarted(tick int, e events.FireGrenadeStart) FireStartedEvent {
	return FireStartedEvent{
		Position: Position{
			X: e.Position.X,
			Y: e.Position.Y,
			Z: e.Position.Z,
		},
		Name: "FIRE_STARTED",
		Tick: tick,
		Id:   e.GrenadeEntityID,
	}
}

func RenderFireExpired(template map[string]interface{}, event FireExpiredEvent) RenderedFireStarted {
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

// FireExpired creates a GameEvent for FIRE_EXPIRED of the real Event.
func FireExpired(tick int, e events.FireGrenadeExpired) FireExpiredEvent {
	return FireExpiredEvent{
		Position: Position{
			X: e.Position.X,
			Y: e.Position.Y,
			Z: e.Position.Z,
		},
		Name: "FIRE_EXPIRED",
		Tick: tick,
		Id:   e.GrenadeEntityID,
	}
}
