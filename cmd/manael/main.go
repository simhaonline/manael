package main // import "manael.org/x/manael/cmd/manael"

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/gorilla/handlers"
	"manael.org/x/manael"
)

var bind = flag.String("bind", "0.0.0.0", "")
var port = flag.String("port", "8080", "")
var upstreamURL = flag.String("upstream-url", "http://localhost:9000", "")

func main() {
	flag.Parse()

	runtime.GOMAXPROCS(runtime.NumCPU())

	h := &manael.Handler{
		UpstreamURL: *upstreamURL,
	}

	addr := fmt.Sprintf("%s:%s", *bind, *port)
	err := http.ListenAndServe(addr, handlers.LoggingHandler(os.Stdout, h))
	if err != nil {
		log.Fatal(err)
	}
}