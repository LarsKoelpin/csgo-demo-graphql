package usecase

import (
	"io"
	"log"
	"math"

	"github.com/larskoelpin/csgo-demo-graphql/pkg/domain"
	dem "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
)

func RecordDemo(file io.Reader, freq float64) domain.Demo {
	p := dem.NewParser(file)
	header, _ := p.ParseHeader()

	allEvents := make([]domain.GameEvent, 0)

	p.RegisterEventHandler(func(e events.BombPlanted) {
		allEvents = append(allEvents, domain.GameEvent{
			Name: "BOMB_PLANTED",
			RealEvent: domain.BombPlanted{
				Name:     "BOMB_PLANTED",
				Player:   domain.CreateParticipant(e.Player),
				Bombsite: int32(e.Site),
			},
		})
	})

	p.RegisterEventHandler(func(e events.WeaponFire) {
		allEvents = append(allEvents, domain.GameEvent{
			Name: "WEAPON_FIRED",
			RealEvent: domain.WeaponFired{
				Shooter: domain.CreateParticipant(e.Shooter),
				Weapon:  domain.FromEquipment(e.Weapon),
			},
		})
	})

	p.RegisterEventHandler(func(e events.SmokeStart) {
		allEvents = append(allEvents, domain.SmokeStarted(p.GameState().IngameTick(), e))
	})

	p.RegisterEventHandler(func(e events.SmokeExpired) {
		allEvents = append(allEvents, domain.SmokeExpired(p.GameState().IngameTick(), e))
	})

	p.RegisterEventHandler(func(e events.FireGrenadeStart) {
		allEvents = append(allEvents, domain.FireStarted(p.GameState().IngameTick(), e))
	})

	p.RegisterEventHandler(func(e events.FireGrenadeExpired) {
		allEvents = append(allEvents, domain.FireExpired(p.GameState().IngameTick(), e))
	})

	p.RegisterEventHandler(func(e events.RoundStart) {
		allEvents = append(allEvents, domain.RoundStarted(p.GameState().IngameTick(), e))
	})

	p.RegisterEventHandler(func(e events.RoundEnd) {
		allEvents = append(allEvents, domain.RoundEnded(p.GameState().IngameTick(), e))
	})

	p.RegisterEventHandler(func(e events.MatchStart) {
		allEvents = append(allEvents, domain.NewMatchStartedEvent(p.GameState().IngameTick(), e))
	})

	p.RegisterEventHandler(func(e events.FlashExplode) {
		allEvents = append(allEvents, domain.NewFlashExplosion(p.GameState().IngameTick(), e))
	})

	snapshotRate := int(math.Round(header.FrameRate() / freq))
	renderedTicks := make([]domain.Tick, 0)
	p.RegisterEventHandler(
		func(e events.FrameDone) {
			tick := p.CurrentFrame()
			players := make([]domain.Player, 0)
			grenades := make([]domain.Grenade, 0)

			if tick%snapshotRate == 0 {
				for _, pl := range p.GameState().Participants().Playing() {
					e := domain.CreateParticipant(pl)

					players = append(players, e)
				}

				for _, grenade := range p.GameState().GrenadeProjectiles() {
					e := domain.NewProjectile(*grenade)
					grenades = append(grenades, e)
				}

				renderedTicks = append(renderedTicks, domain.Tick{
					Tick:              tick,
					Players:           players,
					Grenades:          grenades,
					TotalRoundsPlayed: p.GameState().TotalRoundsPlayed(),
				})
			}
		})

	p.ParseToEnd()

	log.Print("Recording finished")

	return domain.Demo{
		Header: domain.Header{
			MapName:      header.MapName,
			TickRate:     header.FrameRate(),
			SnapshotRate: 1,
		},
		Ticks:  renderedTicks,
		Events: allEvents,
	}
}
