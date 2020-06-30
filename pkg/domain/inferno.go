package domain

import (
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
)

// Inferno represents a burning fire.
type Inferno struct {
	Hull []Position `json:"hull"`
}

func ToInferno(inferno common.Inferno) Inferno {
	return Inferno{
		Hull: FromPoints(inferno.Fires().ConvexHull2D()),
	}
}
