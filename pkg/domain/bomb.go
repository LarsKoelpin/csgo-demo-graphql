package domain

import "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"

// Bomb is the structure representing the json of the bombstate in a tick.
type Bomb struct {
	LastOnGroundPosition Position
	Carrier              int
	Position             Position
	IsPlanted            bool
}

func NewBomb(b *common.Bomb, planted bool) Bomb {
	carrier := -1
	if b.Carrier != nil {
		carrier = b.Carrier.EntityID
	}
	return Bomb{
		LastOnGroundPosition: FromVector(b.LastOnGroundPosition),
		Position:             FromVector(b.Position()),
		Carrier:              carrier,
		IsPlanted:            planted,
	}
}

func RenderBomb(template map[string]interface{}, b Bomb) map[string]interface{} {
	renderedBomb := map[string]interface{}{}
	_, hasCarrier := template["carrier"]
	lastOngroundTpl, hasLastOnGround := template["lastOnGroundPosition"]
	posTemplate, hasPosition := template["position"]
	_, hasPlanted := template["planted"]

	if hasLastOnGround {
		tpl, ok := lastOngroundTpl.(map[string]interface{})
		if ok {
			renderedBomb["lastOnGroundPosition"] = RenderPosition(tpl, b.LastOnGroundPosition)
		}
	}

	if hasCarrier {
		renderedBomb["carrier"] = b.Carrier
	}

	if hasPosition {
		tpl, ok := posTemplate.(map[string]interface{})
		if ok {
			renderedBomb["position"] = RenderPosition(tpl, b.Position)
		}
	}

	if hasPlanted {
		renderedBomb["planted"] = b.IsPlanted
	}

	return renderedBomb
}
