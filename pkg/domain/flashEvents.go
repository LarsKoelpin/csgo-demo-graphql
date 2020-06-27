package domain

import (
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
)

type FlashExplosionEvent struct {
	Position Position `json:"position"`
	Name     string   `json:"name"`
	Tick     int      `json:"tick"`
}

func RenderFlashExplosionEvent(template map[string]interface{}, e FlashExplosionEvent) map[string]interface{} {
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

func NewFlashExplosion(tick int, e events.FlashExplode) FlashExplosionEvent {
	return FlashExplosionEvent{
		Position: Position{
			X: e.Position.X,
			Y: e.Position.Y,
			Z: e.Position.Z,
		},
		Name: "FLASH_EXPLOSION",
		Tick: tick,
	}
}
