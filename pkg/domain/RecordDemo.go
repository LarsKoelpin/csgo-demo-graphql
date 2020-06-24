package domain

import (
	"io"
	"log"
	"math"

	dem "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
)

// RecordDemo models the process of replaying the demo while recording all events.
func RecordDemo(file io.Reader, freq float64) Demo {
	p := dem.NewParser(file)
	header, _ := p.ParseHeader()

	allEvents := make([]GameEvent, 0)
	firing := make(map[int]bool)
	p.RegisterEventHandler(func(e events.BombPlanted) {
		allEvents = append(allEvents, GameEvent{
			Name: "BOMB_PLANTED",
			RealEvent: BombPlanted{
				Name:     "BOMB_PLANTED",
				Player:   CreateParticipant(e.Player, firing[e.Player.EntityID]),
				Bombsite: int32(e.Site),
			},
		})
	})

	p.RegisterEventHandler(func(e events.WeaponFire) {
		firing[e.Shooter.EntityID] = true
		allEvents = append(allEvents, GameEvent{
			Name: "WEAPON_FIRED",
			RealEvent: WeaponFired{
				Shooter: CreateParticipant(e.Shooter, firing[e.Shooter.EntityID]),
				Weapon:  FromEquipment(e.Weapon),
			},
		})
	})

	smokes := make([]Smoke, 0)

	p.RegisterEventHandler(func(e events.SmokeStart) {
		allEvents = append(allEvents, SmokeStarted(p.GameState().IngameTick(), e))
		smokes = append(smokes, Smoke{
			Id: e.GrenadeEntityID,
			Position: Position{
				X: e.Position.X,
				Y: e.Position.Y,
			}})
	})

	p.RegisterEventHandler(func(e events.SmokeExpired) {
		allEvents = append(allEvents, SmokeExpired(p.GameState().IngameTick(), e))
		smokes = Remove(smokes, e.GrenadeEntityID)
	})

	p.RegisterEventHandler(func(e events.FireGrenadeStart) {
		allEvents = append(allEvents, FireStarted(p.GameState().IngameTick(), e))
	})

	p.RegisterEventHandler(func(e events.FireGrenadeExpired) {
		allEvents = append(allEvents, FireExpired(p.GameState().IngameTick(), e))
	})

	p.RegisterEventHandler(func(e events.RoundStart) {
		allEvents = append(allEvents, RoundStarted(p.GameState().IngameTick(), e))
    smokes = make([]Smoke, 0)
	})

	p.RegisterEventHandler(func(e events.RoundEnd) {
		allEvents = append(allEvents, RoundEnded(p.GameState().IngameTick(), e))
		smokes = make([]Smoke, 0)
	})

	p.RegisterEventHandler(func(e events.MatchStart) {
		allEvents = append(allEvents, NewMatchStartedEvent(p.GameState().IngameTick(), e))
    smokes = make([]Smoke, 0)
	})

	p.RegisterEventHandler(func(e events.FlashExplode) {
		allEvents = append(allEvents, NewFlashExplosion(p.GameState().IngameTick(), e))
	})

	p.RegisterEventHandler(func(e events.HeExplode) {
		allEvents = append(allEvents, NewHeExplosion(p.GameState().IngameTick(), e))
	})

	snapshotRate := int(math.Round(header.FrameRate() / freq))
	renderedTicks := make([]Tick, 0)
	p.RegisterEventHandler(
		func(e events.FrameDone) {
			tick := p.GameState().IngameTick()
			players := make([]Player, 0)
			grenades := make([]Grenade, 0)
			infernos := make([]Inferno, 0)

			if tick%snapshotRate == 0 {
				for _, pl := range p.GameState().Participants().Playing() {
					e := CreateParticipant(pl, firing[pl.EntityID])

					players = append(players, e)
					firing[pl.EntityID] = false
				}

				for _, grenade := range p.GameState().GrenadeProjectiles() {
					e := NewProjectile(*grenade, firing)
					grenades = append(grenades, e)
				}

				for _, pl := range p.GameState().Infernos() {
					infernos = append(infernos, ToInferno(*pl))
				}

				renderedTicks = append(renderedTicks, Tick{
					Tick:              tick,
					Players:           players,
					Grenades:          grenades,
					Infernos:          infernos,
					Smokes:            smokes,
					TotalRoundsPlayed: p.GameState().TotalRoundsPlayed(),
				})
			}

		})

	p.ParseToEnd()

	log.Print("Recording finished")

	return Demo{
		Header: Header{
			MapName:  header.MapName,
			TickRate: header.FrameRate(),
			Fps:      int(freq),
		},
		Ticks:  renderedTicks,
		Events: allEvents,
	}
}
