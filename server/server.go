package callisto

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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

	s.router.Handle("GET", "/hello", helloHandler())
	s.router.Run(":3000")

	return nil
}

func helloHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, now time is: %s", time.Now().Format("03:04:05 PM"))
	}
}
