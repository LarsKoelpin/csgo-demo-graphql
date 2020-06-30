package domain

import (
  dem "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs"
  "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
  "io"
  "log"
  "math"
)

// RecordDemo models the process of replaying the demo while recording all events.
func RecordDemo(file io.Reader, freq float64, demoTemplate DemoTemplate) RenderedDemo {
	p := dem.NewParser(file)
	renderedDemo := map[string]interface{}{}

	eventsTemplate, hasEvents := demoTemplate["events"]
	headerTemplate, hasHeaderTemplate := demoTemplate["header"]

	h, _ := p.ParseHeader()
	header := Header{
		MapName:  h.MapName,
		TickRate: p.TickRate(),
		Fps:      int(freq),
		FrameRate: int(math.Round(h.FrameRate())),
	}
	allEvents := make([]map[string]interface{}, 0)
	smokes := make([]Smoke, 0)
	firing := make(map[int]bool)

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
				wf := NewWeaponFired(p.CurrentFrame(), e)
				renderedEvent := RenderWeaponFired(x, wf)
				allEvents = append(allEvents, renderedEvent)
			}
		})

		if hasSmokeStarted {
			x := smokeStartedTemplate.(map[string]interface{})
			p.RegisterEventHandler(func(e events.SmokeStart) {
				smokeEvent := SmokeStarted(p.CurrentFrame(), e)
				renderedSmokeEvent := RenderSmokeStarted(x, smokeEvent)
				allEvents = append(allEvents, renderedSmokeEvent)

				newSmoke := Smoke{
					Id: e.GrenadeEntityID,
					Position: Position{
						X: e.Position.X,
						Y: e.Position.Y,
					}}
				smokes = append(smokes, newSmoke)
			})
		}

		if hasSmokeExpired {
			x := smokeExpiredTemplate.(map[string]interface{})
			p.RegisterEventHandler(func(e events.SmokeExpired) {
				smokeExpiredEvent := SmokeExpired(p.CurrentFrame(), e)
				renderedExpiredEvent := RenderSmokeExpired(x, smokeExpiredEvent)
				allEvents = append(allEvents, renderedExpiredEvent)
				smokes = Remove(smokes, e.GrenadeEntityID)
			})
		}

		if hasFireStarted {
			x := fireStartedTemplate.(map[string]interface{})
			p.RegisterEventHandler(func(e events.FireGrenadeStart) {
				fireStartedEvent := FireStarted(p.CurrentFrame(), e)
				renderedStartedEvent := RenderFireStarted(x, fireStartedEvent)
				allEvents = append(allEvents, renderedStartedEvent)
			})
		}

		if hasFireExpired {
			x := fireExpiredTemplate.(map[string]interface{})
			p.RegisterEventHandler(func(e events.FireGrenadeExpired) {
				fireExpiredEvent := FireExpired(p.CurrentFrame(), e)
				renderedExpiredEvent := RenderFireExpired(x, fireExpiredEvent)
				allEvents = append(allEvents, renderedExpiredEvent)
				smokes = Remove(smokes, e.GrenadeEntityID)
			})
		}

		if hasRoundStarted {
			p.RegisterEventHandler(func(e events.RoundStart) {
				tpl := roundStartedTemplate.(map[string]interface{})
				rs := RoundStarted(p.CurrentFrame(), e)
				renderedRoundStarted := RenderRoundStarted(tpl, rs)
				allEvents = append(allEvents, renderedRoundStarted)
				smokes = make([]Smoke, 0)
			})
		}

		if hasRoundEnded {
			p.RegisterEventHandler(func(e events.RoundEnd) {
				tpl := roundEndedTemplate.(map[string]interface{})
				rs := RoundEnded(p.CurrentFrame(), e)
				rendered := RenderRoundEnded(tpl, rs)
				allEvents = append(allEvents, rendered)
				smokes = make([]Smoke, 0)
			})
		}

		if hasMatchStarted {
			p.RegisterEventHandler(func(e events.MatchStart) {
				tpl := matchStartedTemplate.(map[string]interface{})
				matchStarted := NewMatchStartedEvent(p.CurrentFrame(), e)
				renderedMatchStarted := RenderMatchStartedEvent(tpl, matchStarted)
				allEvents = append(allEvents, renderedMatchStarted)
				smokes = make([]Smoke, 0)
			})
		}

		if hasMatchStarted {
			p.RegisterEventHandler(func(e events.MatchStart) {
				tpl := matchStartedTemplate.(map[string]interface{})
				matchStarted := NewMatchStartedEvent(p.CurrentFrame(), e)
				renderedMatchStarted := RenderMatchStartedEvent(tpl, matchStarted)
				allEvents = append(allEvents, renderedMatchStarted)
				smokes = make([]Smoke, 0)
			})
		}

		if hasFlashExplosion {
			p.RegisterEventHandler(func(e events.FlashExplode) {
				tpl := flashExplosionTemplate.(map[string]interface{})
				explosion := NewFlashExplosion(p.CurrentFrame(), e)
				renderedMatchStarted := RenderFlashExplosionEvent(tpl, explosion)
				allEvents = append(allEvents, renderedMatchStarted)
			})
		}

		if hasHeExplosion {
			p.RegisterEventHandler(func(e events.HeExplode) {
				tpl := heExplosionTemplate.(map[string]interface{})
				explosion := NewHeExplosion(p.CurrentFrame(), e)
				renderedHeExplosion := RenderHEExplosionEvent(tpl, explosion)
				allEvents = append(allEvents, renderedHeExplosion)
			})
		}

	}

	frameRate := math.Round(h.FrameRate())
	tickRate := math.Round(p.TickRate())

	var snapshotRate = -1;
  // This is perfecly fine :shurg;
	if(frameRate == 32 && tickRate == 128) {
	  log.Print("Use Framerate 32/128")
    snapshotRate = int(math.Round(tickRate/ freq))
  }
  // This is perfecly fine :shurg;
  if(frameRate == 64 && tickRate == 128) {
	  log.Print("Use Framerate 64/128")
    snapshotRate = int(math.Round(tickRate/ freq)) / 2
  }

  // This is perfecly fine :shurg;
  if(frameRate == 32 && tickRate == 64) {
    log.Print("Use Framerate 32/64")
    snapshotRate = int(math.Round(tickRate/ freq)) / 2
  }

  if (snapshotRate == -1 || snapshotRate == 0) {
    log.Fatal("Could not determine snapshotrate for framerate ", frameRate)
  }

  log.Print("Use tickrate", tickRate)
  log.Print("Use snapshotrate", snapshotRate)


  renderedTicks := make([]RenderedTick, 0)
	p.RegisterEventHandler(
		func(e events.FrameDone) {
			tick := p.CurrentFrame()
			players := make([]Player, 0)
			grenades := make([]Grenade, 0)
			infernos := make([]Inferno, 0)

			ticksTemplate, hasTicks := demoTemplate["ticks"]
			if hasTicks {
				if tick%snapshotRate == 0 {
					templateOfDemo, ok := ticksTemplate.(map[string]interface{})
					if !ok {
						log.Panic("NOPE")
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
						Tick:              int(tick),
						Players:           players,
						Grenades:          grenades,
						Infernos:          infernos,
						Smokes:            smokes,
						TotalRoundsPlayed: p.GameState().TotalRoundsPlayed(),
					})
					renderedTicks = append(renderedTicks, renderedTick)
				}
			}
		})

	p.ParseToEnd()

	renderedDemo["events"] = allEvents
	renderedDemo["ticks"] = renderedTicks
	if hasHeaderTemplate {
		casted, _ := headerTemplate.(map[string]interface{})
		renderedDemo["header"] = RenderHeader(casted, header)
	}

	log.Print("Recording finished")

	return renderedDemo
}
