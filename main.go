package main

import (
	"fmt"
	"github.com/garbotron/donkeytownsfolk"
	"github.com/garbotron/giddeongarber.info"
	"github.com/garbotron/goshots"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	_ "net/http/pprof"
)

const httpPort = 80

func main() {

	r := mux.NewRouter()

	if err := giddeongarber.Init(r); err != nil {
		log.Fatal(err)
		return
	}

	if err := goshots.Init(r); err != nil {
		log.Fatal(err)
		return
	}

	if err := donkeytownsfolk.Init(r); err != nil {
		log.Fatal(err)
		return
	}

	http.Handle("/", r)
	http.ListenAndServe(fmt.Sprintf(":%d", httpPort), nil)
}
