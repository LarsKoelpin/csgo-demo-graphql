package domain

import (
	"io"
	"log"
	"math"

	dem "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
)

// RecordDemo models the process of replaying the demo while recording all events.
func RecordDemo(file io.Reader, freq float64, demoTemplate DemoTemplate) RenderedDemo {
	p := dem.NewParser(file)
	renderedDemo := map[string]interface{}{}
	p.ParseHeader()
	allEvents := make([]map[string]interface{}, 0)
	firing := make(map[int]bool)

	smokes := make([]Smoke, 0)

	eventsTemplate, hasEvents := demoTemplate["events"]

	if hasEvents {
		eventTypes := eventsTemplate.(map[string]interface{})
		bombPlantedTemplate, hasBombPlanted := eventTypes["BOMB_PLANTED"]
		smokeStartedTemplate, hasSmokeStarted := eventTypes["SMOKE_STARTED"]
		smokeExpiredTemplate, hasSmokeExpired := eventTypes["SMOKE_EXPIRED"]
		fireStartedTemplate, hasFireStarted := eventTypes["FIRE_STARTED"]
		fireExpiredTemplate, hasFireExpired := eventTypes["FIRE_EXPIRED"]
		roundStartedTemplate, hasRoundStarted := eventTypes["ROUND_STARTED"]
		roundEndedTemplate, hasRoundEnded := eventTypes["ROUND_ENDED"]
		matchStartedTemplate, hasMatchStarted := eventTypes["MATCH_STARTED"]
		flashExplosionTemplate, hasFlashExplosion := eventTypes["FLASH_EXPLOSION"]
		heExplosionTemplate, hasHeExplosion := eventTypes["HE_EXPLOSION"]
		weaponFiredTemplate, hasWeaponFired := eventTypes["WEAPON_FIRED"]

		if hasBombPlanted {
			p.RegisterEventHandler(func(e events.BombPlanted) {
				x := bombPlantedTemplate.(map[string]interface{})
				event := NewBombPlanted(e)
				renderedBombPlantedEvent := RenderBombPlanted(x, event)
				allEvents = append(allEvents, renderedBombPlantedEvent)
			})
		}

		p.RegisterEventHandler(func(e events.WeaponFire) {
			firing[e.Shooter.EntityID] = true
			if hasWeaponFired {
				x := weaponFiredTemplate.(map[string]interface{})
				wf := NewWeaponFired(p.GameState().IngameTick(), e)
				renderedEvent := RenderWeaponFired(x, wf)
				allEvents = append(allEvents, renderedEvent)
			}
		})

		if hasSmokeStarted {
			x := smokeStartedTemplate.(map[string]interface{})
			p.RegisterEventHandler(func(e events.SmokeStart) {
				smokeEvent := SmokeStarted(p.GameState().IngameTick(), e)
				renderedSmokeEvent := RenderSmokeStarted(x, smokeEvent)
				allEvents = append(allEvents, renderedSmokeEvent)
				smokes = append(smokes, Smoke{
					Id: e.GrenadeEntityID,
					Position: Position{
						X: e.Position.X,
						Y: e.Position.Y,
					}})
			})
		}

		if hasSmokeExpired {
			x := smokeExpiredTemplate.(map[string]interface{})
			p.RegisterEventHandler(func(e events.SmokeExpired) {
				smokeExpiredEvent := SmokeExpired(p.GameState().IngameTick(), e)
				renderedExpiredEvent := RenderSmokeExpired(x, smokeExpiredEvent)
				allEvents = append(allEvents, renderedExpiredEvent)
				smokes = Remove(smokes, e.GrenadeEntityID)
			})
		}

		if hasFireStarted {
			x := fireStartedTemplate.(map[string]interface{})
			p.RegisterEventHandler(func(e events.FireGrenadeStart) {
				fireStartedEvent := FireStarted(p.GameState().IngameTick(), e)
				renderedStartedEvent := RenderFireStarted(x, fireStartedEvent)
				allEvents = append(allEvents, renderedStartedEvent)
			})
		}

		if hasFireExpired {
			x := fireExpiredTemplate.(map[string]interface{})
			p.RegisterEventHandler(func(e events.FireGrenadeExpired) {
				fireExpiredEvent := FireExpired(p.GameState().IngameTick(), e)
				renderedExpiredEvent := RenderFireExpired(x, fireExpiredEvent)
				allEvents = append(allEvents, renderedExpiredEvent)
				smokes = Remove(smokes, e.GrenadeEntityID)
			})
		}

		if hasRoundStarted {
			p.RegisterEventHandler(func(e events.RoundStart) {
				tpl := roundStartedTemplate.(map[string]interface{})
				rs := RoundStarted(p.GameState().IngameTick(), e)
				renderedRoundStarted := RenderRoundStarted(tpl, rs)
				allEvents = append(allEvents, renderedRoundStarted)
				smokes = make([]Smoke, 0)
			})
		}

		if hasRoundEnded {
			p.RegisterEventHandler(func(e events.RoundEnd) {
				tpl := roundEndedTemplate.(map[string]interface{})
				rs := RoundEnded(p.GameState().IngameTick(), e)
				rendered := RenderRoundEnded(tpl, rs)
				allEvents = append(allEvents, rendered)
				smokes = make([]Smoke, 0)
			})
		}

		if hasMatchStarted {
			p.RegisterEventHandler(func(e events.MatchStart) {
				tpl := matchStartedTemplate.(map[string]interface{})
				matchStarted := NewMatchStartedEvent(p.GameState().IngameTick(), e)
				renderedMatchStarted := RenderMatchStartedEvent(tpl, matchStarted)
				allEvents = append(allEvents, renderedMatchStarted)
				smokes = make([]Smoke, 0)
			})
		}

		if hasMatchStarted {
			p.RegisterEventHandler(func(e events.MatchStart) {
				tpl := matchStartedTemplate.(map[string]interface{})
				matchStarted := NewMatchStartedEvent(p.GameState().IngameTick(), e)
				renderedMatchStarted := RenderMatchStartedEvent(tpl, matchStarted)
				allEvents = append(allEvents, renderedMatchStarted)
				smokes = make([]Smoke, 0)
			})
		}

		if hasFlashExplosion {
			p.RegisterEventHandler(func(e events.FlashExplode) {
				tpl := flashExplosionTemplate.(map[string]interface{})
				explosion := NewFlashExplosion(p.GameState().IngameTick(), e)
				renderedMatchStarted := RenderFlashExplosionEvent(tpl, explosion)
				allEvents = append(allEvents, renderedMatchStarted)
			})
		}

		if hasHeExplosion {
			p.RegisterEventHandler(func(e events.HeExplode) {
				tpl := heExplosionTemplate.(map[string]interface{})
				explosion := NewHeExplosion(p.GameState().IngameTick(), e)
				renderedHeExplosion := RenderHEExplosionEvent(tpl, explosion)
				allEvents = append(allEvents, renderedHeExplosion)
			})
		}

	}

	snapshotRate := int(math.Round(p.TickRate() / freq))
	renderedTicks := make([]RenderedTick, 0)
	p.RegisterEventHandler(
		func(e events.FrameDone) {
			tick := p.GameState().IngameTick()
			players := make([]Player, 0)
			grenades := make([]Grenade, 0)
			infernos := make([]Inferno, 0)

			ticksTemplate, hasTicks := demoTemplate["ticks"]
			if hasTicks {
				if tick%snapshotRate == 0 {
					templateOfDemo, ok := ticksTemplate.(map[string]interface{})
					if !ok {
						return
					}
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
					renderedTick := RenderTick(templateOfDemo, Tick{
						Tick:              tick,
						Players:           players,
						Grenades:          grenades,
						Infernos:          infernos,
						Smokes:            smokes,
						TotalRoundsPlayed: p.GameState().TotalRoundsPlayed(),
					})
					renderedTicks = append(renderedTicks, renderedTick)
				}
				renderedDemo["ticks"] = renderedTicks
			}
		})

	p.ParseToEnd()

	renderedDemo["events"] = allEvents

	log.Print("Recording finished")

	return renderedDemo
}
