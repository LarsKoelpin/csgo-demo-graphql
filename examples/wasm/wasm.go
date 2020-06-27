package main

import (
	"bytes"
	"fmt"
	"syscall/js"

	"github.com/larskoelpin/csgo-demo-graphql/pkg/domain"
	"github.com/larskoelpin/csgo-demo-graphql/pkg/usecase"
)

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	fmt.Println("WASM Go Initialized XDD")
	<-c
}

func registerCallbacks() {
	js.Global().Set("parse", js.FuncOf(parse))
}

func uint8ArrayToBytes(value js.Value) []byte {
	s := make([]byte, value.Get("byteLength").Int())
	js.CopyBytesToGo(s, value)
	return s
}

func parse(this js.Value, args []js.Value) interface{} {
	parseInternal(args[0], args[1])
	return nil
}

func parseInternal(data js.Value, callback js.Value) {
	b := bytes.NewBuffer(uint8ArrayToBytes(data))
	demoRepository := domain.DemoRepository{}
	schema := usecase.SchemaFromDemo(b, demoRepository)
	json := usecase.CreateJson(schema, `
{
  demo(fps: 20) {
    header {
      mapName
      tickRate
      fps
    }
    ticks {
      tick
      totalRoundsPlayed
      smokes {
        position {
          x
          y
        }
      }
      infernos {
        hull {
          x
          y
        }
      }
      players {
        entityId
        firing
        name
        armor
        angleX
        angleY
        hp
        deaths
        kills
        equipment {
          ammoInMagazine
          ammoReserve
          type
        }
        flashDuration
        hasDefuseKit
        hasHelmet
        kills
        money
        name
        team
        position {
          x
          y
        }
        equipment {
          ammoInMagazine
          ammoReserve
          type
        }
      }
    }
    events(
      type: [
        "MATCH_STARTED"
        "ROUND_STARTED"
        "ROUND_ENDED"
        "FLASH_EXPLOSION"
        "HE_EXPLOSION"
      ]
    ) {
     ... on HEExploded {
         name
         tick
         position {
            x
            y
         }
      }
      ... on FlashExploded {
        name
        tick
        position {
          x
          y
        }
      }
      ... on MatchStarted {
        name
        tick
      }

      ... on RoundEnded {
        name
        tick
      }

      ... on RoundStarted {
        name
        tick
      }
    }
  }
}

  `)
	fmt.Println("parsed")

	callback.Invoke(string(json))
}
