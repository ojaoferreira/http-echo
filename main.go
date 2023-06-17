package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	text string = "bar"
)

type ResponseJSON struct {
	Text string `json:"text"`
}

func main() {

	flag.StringVar(&text, "text", "bar", "a string var")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-type") == "application/json" {
			w.Header().Add("Content-Type", "application/json")
			reponseJSON := ResponseJSON{Text: text}
			json.NewEncoder(w).Encode(reponseJSON)
			return
		}
		fmt.Fprint(w, text)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
