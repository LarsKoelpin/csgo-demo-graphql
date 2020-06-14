package usecase

import (
	"log"
	"math"
	"os"

	"github.com/larskoelpin/csgo-demo-graphql/pkg/domain"
	dem "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
)

func DemoFromFile(filePath string, freq float64) domain.Demo {
	f, err := os.Open(filePath)

	if err != nil {
		log.Print("File not found!", filePath)
		os.Exit(0)
	}
	p := dem.NewParser(f)
	header, _ := p.ParseHeader()

	snapshotRate := int(math.Round(header.FrameRate() / freq))
	renderedTicks := make([]domain.Tick, 0)
	p.RegisterEventHandler(
		func(e events.FrameDone) {
			tick := p.CurrentFrame()
			players := make([]domain.Participant, 0)

			if tick%snapshotRate == 0 {
				for _, pl := range p.GameState().Participants().Playing() {
					e := domain.Participant{
						Name:          pl.Name,
						EntityID:      pl.EntityID,
						Hp:            pl.Health(),
						Armor:         pl.Armor(),
						FlashDuration: 0.1, // Round to nearest 0.1 sec - saves space in JSON
						Position: domain.Position{
							X: pl.Position().X,
							Y: pl.Position().Y,
							Z: pl.Position().Z,
						},
						AngleX:       int(pl.ViewDirectionX()),
						AngleY:       int(pl.ViewDirectionY()),
						HasHelmet:    pl.HasHelmet(),
						HasDefuseKit: pl.HasDefuseKit(),
						Equipment:    domain.ToEntityEquipment(pl.Weapons()),
						Team:         int(pl.Team),
						IsDefusing:   pl.IsDefusing,
						IsPlanting:   pl.IsPlanting,
						Money:        pl.Money(),
						Kills:        pl.Kills(),
						Deaths:       pl.Deaths(),
						IsInBuyzone:  pl.IsInBuyZone(),
					}

					players = append(players, e)
				}
				renderedTicks = append(renderedTicks, domain.Tick{
					Tick:              tick,
					Participants:      players,
					TotalRoundsPlayed: p.GameState().TotalRoundsPlayed(),
				})
			}
		})

	p.ParseToEnd()

	log.Print("Parsing finished")

	return domain.Demo{
		Header: domain.Header{
			MapName:      header.MapName,
			TickRate:     header.FrameRate(),
			SnapshotRate: 1,
		},
		Ticks: renderedTicks,
	}
}
