package callisto

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// Server is a gin server wrapper
type Server struct {
	router *gin.Engine
}

// Serve the http request all from here
func Serve() error {
	s := &Server{
		router: gin.Default(),
	}

	s.router.Use(static.Serve("/", static.LocalFile(os.Getenv("SITE"), true)))
	s.router.Handle("GET", "/hello", helloHandler())

	return s.router.Run()
}

func helloHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, now time is: %s", time.Now().Format("03:04:05 PM"))
	}
}
