package main

import (
  "encoding/json"
  "fmt"
  "github.com/graphql-go/graphql"
  dem "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs"
  "log"
  "os"
)

func main() {
  f, _ := os.Open("/home/lars/devel/src/github.com/markus-wa/cs-demo-minifier/cmd/csminify/test.dem")
  p := dem.NewParser(f)
  header, _ := p.ParseHeader()

  demo := Demo{
    Header: Header{
      MapName:      header.MapName,
      TickRate:     header.FrameRate(),
      SnapshotRate: 1,
      ClientName:   header.ClientName,
    },
  };
  demoType := CreateDemoType(demo);


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
