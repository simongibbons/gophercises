package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/simongibbons/gophercises/urlshort"
	"io/ioutil"
	"log"
)

var (
	flagYAMLPath string
)

func main() {
	parseFlags()

	yaml, err := ioutil.ReadFile(flagYAMLPath)
	if err != nil {
		log.Fatalf("Couldn't read yaml config: %s. Error: %s", flagYAMLPath, err)
	}

	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), defaultMux())
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func parseFlags() {
	flag.StringVar(&flagYAMLPath, "yaml", "redirects.yaml", "yaml file to read redirects from")
	flag.Parse()
}
