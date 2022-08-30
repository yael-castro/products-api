package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"

	"github.com/yael-castro/products-api/internal/dependency"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	log.SetFlags(log.Flags() | log.Lshortfile)

	if os.Getenv("GIN_MODE") == "" {
		gin.SetMode(gin.TestMode)
	}

	var h http.Handler

	err := dependency.NewInjector(dependency.Default).Inject(&h)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf(`http server is running on port "%v" %v`, port, "🤘\n")
	log.Fatal(http.ListenAndServe(":"+port, h))
}
