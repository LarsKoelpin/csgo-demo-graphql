package domain

// Demo represents the json of a whole demo.
type Demo struct {
	Header Header      `json:"header"`
	Ticks  []Tick      `json:"ticks"`
	Events []GameEvent `json:"events"`
}

// DemoTemplate is the main data structure. It get costructed in RecordDemo.go.
type DemoTemplate map[string]interface{}
type RenderedDemo map[string]interface{}

// DemoRepository holds the state of the whole Operation. In this case a simple in memory database
type DemoRepository struct {
	CurrentDemo Demo
}
