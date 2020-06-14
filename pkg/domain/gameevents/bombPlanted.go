package gameevents

import "github.com/larskoelpin/csgo-demo-graphql/pkg/domain"

type BombPlanted struct {
	Player   domain.Participant
	Bombsite int32
}
