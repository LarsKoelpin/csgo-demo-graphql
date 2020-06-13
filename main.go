package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"

	"github.com/graphql-go/graphql"
	dem "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
)

func main() {
	f, _ := os.Open("/home/lars/devel/src/github.com/markus-wa/cs-demo-minifier/cmd/csminify/test.dem")
	p := dem.NewParser(f)
	header, _ := p.ParseHeader()

	snapshotRate := int(math.Round(header.FrameRate() / 0.1))
	renderedTicks := make([]Tick, 0)
	p.RegisterEventHandler(
		func(e events.FrameDone) {
			tick := p.CurrentFrame()
			players := make([]Participant, 0)

			if tick%snapshotRate == 0 {
				for _, pl := range p.GameState().Participants().Playing() {
					e := Participant{
						Name:          pl.Name,
						EntityID:      pl.EntityID,
						Hp:            pl.Health(),
						Armor:         pl.Armor(),
						FlashDuration: 0.1, // Round to nearest 0.1 sec - saves space in JSON
						Position: Position{
							X: pl.Position().X,
							Y: pl.Position().Y,
							Z: pl.Position().Z,
						},
						AngleX:       int(pl.ViewDirectionX()),
						AngleY:       int(pl.ViewDirectionY()),
						HasHelmet:    pl.HasHelmet(),
						HasDefuseKit: pl.HasDefuseKit(),
						Equipment:    toEntityEquipment(pl.Weapons()),
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
				renderedTicks = append(renderedTicks, Tick{
					Tick:              tick,
					Participants:      players,
					TotalRoundsPlayed: p.GameState().TotalRoundsPlayed(),
				})
			}
		})

	p.ParseToEnd()

	demo := Demo{
		Header: Header{
			MapName:      header.MapName,
			TickRate:     header.FrameRate(),
			SnapshotRate: 1,
			ClientName:   header.ClientName,
		},
		Ticks: renderedTicks,
	}
	demoType := CreateDemoType(demo)

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: graphql.Fields{
		"demo": {
			Name: "Demo",
			Type: demoType,
			Args: nil,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return demo, nil
			},
			DeprecationReason: "",
			Description:       "",
		},
	}}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	query := `
		{
      demo {
			  header {
          mapName
        }
        ticks {
          tick
          totalRoundsPlayed
          participants {
            entityId
            team
            position {
              x
            }
          }
          bomb {
           carrier {
             entityId

            }
          }
        }
      }
		}
	`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON) // {"data":{"hello":"world"}}

	fmt.Println(header.ClientName)
}
