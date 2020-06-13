package main

import (
  usecase "github.com/larskoelpin/csgo-demo-graphql/usecase"
  "io/ioutil"
  "log"
  "os"
  "fmt"
)

func main() {
  userQuery := readQuery("data.query");
  fmt.Print(userQuery)

	demo := usecase.CreateDemo("/home/lars/devel/src/github.com/markus-wa/cs-demo-minifier/cmd/csminify/test.dem");
  schema := usecase.CreateGraphqlSchema(demo);
  json := usecase.CreateJson(schema, userQuery);

  file, _ := os.Create("/home/lars/test.json")
  fmt.Println("WRITE")
  file.Write([]byte(json));
}

func readQuery(file string) string {
  file2, err := os.Open(file)
  if err != nil {
    log.Fatal(err)
  }
  defer file2.Close()

  b, err := ioutil.ReadAll(file2)
  return string(b);
}
