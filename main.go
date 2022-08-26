package main

import (
	"github.com/yael-castro/layered-architecture/internal/handler"
	"log"
	"net/http"
	"os"

	"github.com/yael-castro/layered-architecture/internal/dependency"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	log.SetFlags(log.Flags() | log.Lshortfile)

	h := handler.Handler{}

	err := dependency.NewInjector(dependency.Default).Inject(&h)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf(`http server is running on port "%v" %v`, port, "🤘\n")
	log.Fatal(http.ListenAndServe(":"+port, h))
}
