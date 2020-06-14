package usecase

import (
	"log"
	"math"
	"os"

	"github.com/larskoelpin/csgo-demo-graphql/pkg/domain"
	dem "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
)

func RecordDemo(filePath string, freq float64) domain.Demo {
	f, err := os.Open(filePath)
	log.Print("Recording demo ", filePath)

	if err != nil {
		log.Print("File not found!", filePath)
		os.Exit(0)
	}
	p := dem.NewParser(f)
	header, _ := p.ParseHeader()

	allEvents := make([]interface{}, 0)

	p.RegisterEventHandler(func(e events.BombPlanted) {
		allEvents = append(allEvents, domain.BombPlanted{
			Name:     "BOMB_PLANTED",
			Player:   domain.CreateParticipant(e.Player),
			Bombsite: int32(e.Site),
		})
	})

	snapshotRate := int(math.Round(header.FrameRate() / freq))
	renderedTicks := make([]domain.Tick, 0)
	p.RegisterEventHandler(
		func(e events.FrameDone) {
			tick := p.CurrentFrame()
			players := make([]domain.Participant, 0)

			if tick%snapshotRate == 0 {
				for _, pl := range p.GameState().Participants().Playing() {
					e := domain.CreateParticipant(pl)

					players = append(players, e)
				}
				renderedTicks = append(renderedTicks, domain.Tick{
					Tick:              tick,
					Participants:      players,
					TotalRoundsPlayed: p.GameState().TotalRoundsPlayed(),
				})
			}
		})

	err = p.ParseToEnd()

	if err != nil {
		log.Print("Error when paring the demo", err)
		os.Exit(0)
	}

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
