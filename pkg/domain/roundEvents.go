package domain

import (
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
)

type RoundStartedEvent struct {
	Name      string `json:"name"`
	Tick      int    `json:"tick"`
	FragLimit int    `json:"fragLimit"`
	Objective string `json:"objective"`
	TimeLimit int    `json:"timeLimit"`
}

type RoundEndEvent struct {
	Name string `json:"name"`
	Tick int    `json:"tick"`
}

func RenderRoundStarted(template map[string]interface{}, event RoundStartedEvent) map[string]interface{} {
	renderedRoundStarted := map[string]interface{}{}
	_, hasName := template["name"]
	_, hasTick := template["tick"]
	_, hasTimeLimit := template["timelimit"]

	if hasName {
		renderedRoundStarted["name"] = event.Name
	}
	if hasTick {
		renderedRoundStarted["tick"] = event.Tick
	}
	if hasTimeLimit {
		renderedRoundStarted["timelimit"] = event.TimeLimit
	}

	return renderedRoundStarted
}

func RoundStarted(tick int, e events.RoundStart) RoundStartedEvent {
	return RoundStartedEvent{
		Name:      "ROUND_STARTED",
		Tick:      tick,
		FragLimit: e.FragLimit,
		Objective: e.Objective,
		TimeLimit: e.TimeLimit,
	}
}

func RenderRoundEnded(template map[string]interface{}, event RoundEndEvent) map[string]interface{} {
	result := map[string]interface{}{}
	_, hasName := template["name"]
	_, hasTick := template["tick"]

	if hasName {
		result["name"] = event.Name
	}
	if hasTick {
		result["tick"] = event.Tick
	}

	return result
}

func RoundEnded(tick int, e events.RoundEnd) RoundEndEvent {
	return RoundEndEvent{
		Name: "ROUND_ENDED",
		Tick: tick,
	}
}
