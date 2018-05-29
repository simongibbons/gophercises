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

	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), http.NotFoundHandler())
	if err != nil {
		log.Fatalf("Error parsing yaml config: %s", err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func parseFlags() {
	flag.StringVar(&flagYAMLPath, "yaml", "redirects.yaml", "yaml file to read redirects from")
	flag.Parse()
}
