package main

import "github.com/graphql-go/graphql"

type Header struct {
  MapName      string  `json:"map"`
  TickRate     float64 `json:"tickRate"`         // How many ticks per second
  SnapshotRate int     `json:"snapshotRate"` // How many ticks per snapshot
  ClientName   string  `json:"clientName"`
}

var HeaderType = graphql.NewObject(graphql.ObjectConfig{
  Name: "header",
  Fields: graphql.Fields{
    "mapName": &graphql.Field{
      Type: graphql.String,
    },
    "tickRate": &graphql.Field{
      Type: graphql.Int,
    },
    "snapShotRate": &graphql.Field{
      Type: graphql.Float,
    },
  },
})
