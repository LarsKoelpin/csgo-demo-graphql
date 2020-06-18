package main

import (
	"log"
  "os"

  "github.com/larskoelpin/csgo-demo-graphql/pkg/domain"
	"github.com/larskoelpin/csgo-demo-graphql/pkg/usecase"
)

/**
 * Works on my machine :eyebrows:
 */
var introspectionQuery = `
query IntrospectionQuery {
  __schema {
    queryType {
      name
    }
    mutationType {
      name
    }
    subscriptionType {
      name
    }
    types {
      ...FullType
    }
    directives {
      name
      description
      locations
      args {
        ...InputValue
      }
    }
  }
}

fragment FullType on __Type {
  kind
  name
  description
  fields(includeDeprecated: true) {
    name
    description
    args {
      ...InputValue
    }
    type {
      ...TypeRef
    }
    isDeprecated
    deprecationReason
  }
  inputFields {
    ...InputValue
  }
  interfaces {
    ...TypeRef
  }
  enumValues(includeDeprecated: true) {
    name
    description
    isDeprecated
    deprecationReason
  }
  possibleTypes {
    ...TypeRef
  }
}

fragment InputValue on __InputValue {
  name
  description
  type {
    ...TypeRef
  }
  defaultValue
}

fragment TypeRef on __Type {
  kind
  name
  ofType {
    kind
    name
    ofType {
      kind
      name
      ofType {
        kind
        name
        ofType {
          kind
          name
          ofType {
            kind
            name
            ofType {
              kind
              name
              ofType {
                kind
                name
              }
            }
          }
        }
      }
    }
  }
}
`

var completeQuery = `
{
      demo(fps: 0.1) {
			  header {
          mapName
          tickRate
          fps
        }
        ticks {
          tick
          totalRoundsPlayed
          players {
            name
            entityId
            team
            position {
              x
              y
            }
            angleX
            angleY
            hp
            armor
            flashDuration
            npc
            hasHelmet
            hasDefuseKit
            equipment {
              type
              ammoInMagazine
              ammoReserve
            }
            firing
            planting
            defusing
            inBuyzone
            money
            kills
            deaths
          }
          bomb {
           carrier {
             entityId
             name
            }
            lastOnGroundPosition {
              x
              y
            }
          }
        }
      }
}
`

func main() {
	demoRepository := domain.DemoRepository{}
	log.Print("Reading User query ...")
	log.Print("Creating a schema ...")
	file,_ := os.Open("test.dem");
	schema := usecase.SchemaFromDemo(file, demoRepository)
	introspection := usecase.CreateJson(schema, introspectionQuery)
	usecase.CreateJsonFile("/home/lars/devel/src/github.com/larskoelpin/csgo-demo-graphql/docs/api-explorer/src/schema.json", introspection)
	testdata := usecase.CreateJson(
		schema,
		completeQuery,
	)
	usecase.CreateJsonFile(
		"/home/lars/devel/src/github.com/larskoelpin/csgo-demo-graphql/docs/api-explorer/src/data.json",
		testdata,
	)

}
