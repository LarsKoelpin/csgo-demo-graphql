# CSGO Demo Renderer
[![Go Report Card](https://goreportcard.com/badge/github.com/LarsKoelpin/csgo-demo-graphql)](https://goreportcard.com/report/github.com/LarsKoelpin/csgo-demo-graphql)
![GitHub](https://img.shields.io/github/license/LarsKoelpin/csgo-demo-graphql)

A CSGO Demo renderer using GraphQL. It takes a CSGO Demo in binary format and encodes it into JSON.
It only respects the attributes, which the user specifies in the query, therefore not over or underfetching
any data resulting in a bloated JSON.

## How to run

## Build
If you want to build the app, clone the repository and run
```bash
go run cmd/main.go --query ./examples/prod.graphql --demo examples/test.dem
```

For building the webserver run 

```bash
go build cmd/webserver.go --query ./examples/prod.graphql --port 8080
```

It builds a binary called **main** or **webserver**.

## Execute the application

Run using

```bash
csgodemo --query query.query --demo demo.dem
```

where the **query.query** is a file containing a graphQL query.

```graphql
{
  demo(freq: 0.2) {
    header {
      mapName
    }
    ticks {
      participants {
        entityId
      }
    }
  }
}
```

Where **freq** is the recording FPS.

It creates a File named out.json containing all data.

```bash
ls
+ out.json
```

For exploring purposes, you can use the interactive Graphiql tool available at

https://larskoelpin.github.io/csgo-demo-graphql/

### Run as Webserver
There is also 
```bash 
csgodemo --query query.query --port 8080
```

This hosts a simple HTTP-Server listening on Port 8080. It accepts any request containing a CS:GO Binary.
Please do not use this for public facing apps!

## Considerations

If you want to send the json over the wire, try deflateing it using gzip

creates a File named out.json

```bash
gzip out.json //out.json.gz
```

inflate using

```bash
gzip -d out.json.gz
```

## Supported Events

For full reference see graphiql at https://larskoelpin.github.io/csgo-demo-graphql/

### List of Events

The event is the identifier of the event (This is also the "name" Attribute of the Event).
The GraphQL Type is the Type Name of the Event. This is e.g. needed, when you want to query
the union type of events.

| Event           |    GraphQL     |                  Description |
| --------------- | :------------: | ---------------------------: |
| SMOKE_STARTED   |  SmokeStarted  |      A Smokegrenade exploded |
| SMOKED_EXPIRED  |  SmokeExpired  | Smokegreande effect wore off |
| FIRE_STARTED    |  FireStarted   |        Fire Grenade exploded |
| FIRE_EXPIRED    |  FireExpired   |         Fire effect wore off |
| FLASH_EXPLOSION | FlashExplosion |          Flashbang explosion |
| MATCH_STARTED   |  MatchStarted  |              A Match started |
| ROUND_STARTED   |  RoundStarted  |              A Round started |
| ROUND_ENDED     |   RoundEnded   |                A Round ended |
| WEAPON_FIRED    |  WeaponFired   |               Somebody shot. |
| BOMB_PLANTED    |  BombPlanted   |         The bomb was planted |
