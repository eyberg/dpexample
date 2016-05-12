package main

import (
	"fmt"
	"github.com/deferpanic/deferclient/deferstats"
	"net/http"
	"time"
)

// fast test
func fastHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this request is fast")
}

// slow test
func slowHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(6 * time.Second)
	fmt.Fprintf(w, "this request is slow")
}

// panic test
func panicHandler(w http.ResponseWriter, r *http.Request) {
	panic("There is no need for panic")
	fmt.Fprintf(w, "this request is panic")
}

func main() {
	dps := deferstats.NewClient("v00L0K6CdKjE4QwX5DL1iiODxovAHUfo")
	dps.Setenvironment("unikernel")

	go dps.CaptureStats()

	// no need to change these?
	http.HandleFunc("/fast", dps.HTTPHandlerFunc(fastHandler))
	http.HandleFunc("/slow", dps.HTTPHandlerFunc(slowHandler))
	http.HandleFunc("/panic", dps.HTTPHandlerFunc(panicHandler))

	http.ListenAndServe(":3000", nil)
}
