package domain

type Tick struct {
	Tick              int       `json:"tick"`
	Players           []Player  `json:"players"`
	Grenades          []Grenade `json:"grenades"`
	Infernos          []Inferno `json:"infernos"`
	Smokes            []Smoke   `json:"smokes"`
	Bomb              Bomb      `json:"bomb"`
	TotalRoundsPlayed int       `json:"totalRoundsPlayed"`
}

type TickTemplate map[string]interface{}
type RenderedTick map[string]interface{}

func RenderTick(template TickTemplate, t Tick) RenderedTick {
	renderedTick := map[string]interface{}{}

	smokesTpl, hasSmokes := template["smokes"]
	_, players := template["players"]

	if template["tick"] == true {
		renderedTick["tick"] = t.Tick
	}
	if players {
		renderedTick["players"] = renderPlayers(DefaultPlayerTemplate, t.Players)
	}

	if template["grenades"] == true {
		renderedTick["grenades"] = t.Grenades
	}

	if hasSmokes {
		smokesTemplate := smokesTpl.(map[string]interface{})
		renderedTick["smokes"] = RenderSmokes(smokesTemplate, t.Smokes)
	}

	if template["bomb"] == true {
		renderedTick["bomb"] = t.Bomb
	}

	if template["totalRoundsPlayed"] == true {
		renderedTick["totalRoundsPlayed"] = t.TotalRoundsPlayed
	}
	return renderedTick
}

func renderPlayers(template PlayerTemplate, p []Player) []RenderedPlayer {
	result := make([]RenderedPlayer, 0)
	for _, x := range p {
		result = append(result, RenderPlayer(template, x))
	}
	return result
}
