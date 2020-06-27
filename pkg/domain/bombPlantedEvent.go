package domain

import (
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
)

// BombPlanted represents the event, when the bomb was planted.
type BombPlanted struct {
	Name     string `json:"name"`
	Bombsite int32  `json:"bombsite"`
}

func NewBombPlanted(e events.BombPlanted) BombPlanted {
	return BombPlanted{
		Name:     "BOMB_PLANTED",
		Bombsite: int32(e.Site),
	}
}

// RenderBombPlanted renders a bombplanted event rastered to the given template.
func RenderBombPlanted(template map[string]interface{}, p BombPlanted) map[string]interface{} {
	result := map[string]interface{}{}
	_, hasName := template["name"]
	_, hasBombSite := template["bombsite"]

	if hasName {
		result["name"] = p.Name
	}

	if hasBombSite {
		result["bombsite"] = p.Bombsite
	}

	return result
}
