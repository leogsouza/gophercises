package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/leogsouza/gophercises/urlshort"
)

func main() {
	// define flag for file to read yaml content
	ymlFile := flag.String("yaml-file", "./urls.yml", "The file which has yaml code")
	jsonFile := flag.String("json-file", "./urls.json", "The file which has json code")
	flag.Parse()
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml, err := ioutil.ReadFile(*ymlFile)
	if err != nil {
		panic(err)
	}

	_, err = urlshort.YAMLHandler(yaml, mapHandler)
	if err != nil {
		panic(err)
	}

	json, err := ioutil.ReadFile(*jsonFile)
	if err != nil {
		panic(err)
	}

	jsonHandler, err := urlshort.JSONHandler(json, mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
