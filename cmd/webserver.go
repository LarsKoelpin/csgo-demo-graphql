package main

import (
	"bytes"
	json2 "encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/larskoelpin/csgo-demo-graphql/pkg/domain"
	"github.com/larskoelpin/csgo-demo-graphql/pkg/usecase"
)

var userQuery string

func main() {
	fl := new(flag.FlagSet)
	fl.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage of convertcsgo:")
		fl.PrintDefaults()
		fmt.Fprintln(os.Stderr)
	}

	queryFilePtr := fl.String("query", "", "The query file")
	httpPort := fl.Int("port", 8080, "The HTTP Port of the Webserver. Only considered when 'server' is set to true")
	err := fl.Parse(os.Args[1:])
	if err != nil {
		// Some parsing problem, the flag.Parse() already prints the error to stderr
		return
	}

	if *queryFilePtr == "" {
		log.Fatal("No query file specified. try using --query=/path/to/ile  :(", queryFilePtr)
	}

	log.Print("Reading file " + *queryFilePtr)
	userQuery = usecase.ReadQuery(*queryFilePtr)

	log.Println("Started a WebServer")
	http.HandleFunc("/", HTTPHandler)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*httpPort), nil))
}

func HTTPHandler(w http.ResponseWriter, r *http.Request) {
	var fileAsBytes = StreamToByte(r.Body)

	if fileAsBytes == nil {
		return
	}

	byteReader := bytes.NewReader(fileAsBytes)

	log.Print("Creating a schema ...")
	demo := domain.RecordDemo(byteReader, 24, domain.DefaultDemoTemplate)
	log.Print("Marshalling")
	json, _ := json2.Marshal(demo)
	log.Print("Sending ...")
	w.Write([]byte(json))
}

func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(stream)

	if err != nil {
		return nil
	}

	return buf.Bytes()
}
