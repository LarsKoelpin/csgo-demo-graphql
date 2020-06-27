package domain

type RenderedGameEvent = map[string]interface{}

type GameEvent struct {
	Name      string `json:"name"`
	RealEvent interface{}
}
