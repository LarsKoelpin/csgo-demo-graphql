package main

import (
	"log"

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
      demo(freq: 0.1) {
			  header {
          mapName
          tickRate
          snapshotRate
        }
        ticks {
          tick
          totalRoundsPlayed
          participants {
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
            isNpc
            hasHelmet
            hasDefuseKit
            equipment {
              type
              ammoInMagazine
              ammoReserve
            }
            isPlanting
            isDefusing
            isInBuyzone
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
	schema := usecase.SchemaFromDemo("test.dem", demoRepository)
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
