package domain

// Demo represents the json of a whole demo.
type Demo struct {
	Header Header      `json:"header"`
	Ticks  []Tick      `json:"ticks"`
	Events []GameEvent `json:"events"`
}

type DemoTemplate map[string]interface{}
type RenderedDemo map[string]interface{}

func RenderDemo(demoTemplate DemoTemplate, d Demo) RenderedDemo {
	renderedDemo := map[string]interface{}{}
	ticksTemplate, hasTicks := demoTemplate["ticks"]
	if hasTicks {
		ticksTemplate, _ := ticksTemplate.(map[string]interface{})
		renderedDemo["ticks"] = renderTicks(ticksTemplate, d.Ticks)
	}
	return renderedDemo
}

func renderTicks(template TickTemplate, p []Tick) []RenderedTick {
	result := make([]RenderedTick, 0)
	for _, x := range p {
		result = append(result, RenderTick(template, x))
	}
	return result
}

// DemoRepository holds the state of the whole Operation. In this case a simple in memory database
type DemoRepository struct {
	CurrentDemo Demo
}
