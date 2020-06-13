package usecase

import (
  "io/ioutil"
  "log"
  "os"
)

func ReadQuery(file string) string {
  file2, err := os.Open(file)
  if err != nil {
    log.Fatal(err)
  }
  defer file2.Close()

  b, err := ioutil.ReadAll(file2)
  return string(b)
}
