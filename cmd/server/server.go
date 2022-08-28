package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"

	"github.com/yael-castro/agrak/internal/dependency"
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

	engine := &gin.Engine{}

	err := dependency.NewInjector(dependency.Default).Inject(engine)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf(`http server is running on port "%v" %v`, port, "🤘\n")
	log.Fatal(http.ListenAndServe(":"+port, engine))
}
