package domain

type Header struct {
	MapName  string  `json:"mapName"`
	TickRate float64 `json:"tickRate"` // How many ticks per second
	Fps      int     `json:"fps"`
}

// RenderHeader renders the header to the user selection
func RenderHeader(template map[string]interface{}, h Header) map[string]interface{} {
	result := map[string]interface{}{}

	_, hasMapName := template["mapName"]
	_, hasTickrate := template["tickrate"]
	_, hasFps := template["fps"]

	if hasMapName {
		result["mapName"] = h.MapName
	}

	if hasTickrate {
		result["tickrate"] = h.TickRate
	}

	if hasFps {
		result["fps"] = h.Fps
	}

	return result
}
