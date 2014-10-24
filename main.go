package main

import (
	"fmt"
	"github.com/garbotron/giddeongarber.info"
	"github.com/garbotron/goshots"
	"log"
	"net/http"
	_ "net/http/pprof"
)

const httpPort = 80

func main() {

	if err := giddeongarber.Init(); err != nil {
		log.Fatal(err)
		return
	}

	if err := goshots.Init(); err != nil {
		log.Fatal(err)
		return
	}

	http.ListenAndServe(fmt.Sprintf(":%d", httpPort), nil)
}
