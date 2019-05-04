package callisto

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

// Server is a gin server wrapper
type Server struct {
	authConf *oauth2.Config
	router   *gin.Engine
}

// Serve the http request all from here
func Serve() error {
	godotenv.Load("../../site/.env")

	s := &Server{
		router: gin.Default(),
		authConf: &oauth2.Config{
			ClientID:     os.Getenv("CLIENT_ID"),
			ClientSecret: os.Getenv("CLIENT_SECRET"),
			Endpoint:     github.Endpoint,
		},
	}

	// s.router.Use(static.Serve("/", static.LocalFile(os.Getenv("SITE"), true)))
	s.router.Use(cors.Default())

	s.router.Handle("GET", "/hello", helloHandler())
	s.router.Handle("GET", "/api/auth", authHandler(s.authConf))

	return s.router.Run()
}

func helloHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, now time is: %s", time.Now().Format("03:04:05 PM"))
	}
}

func authHandler(config *oauth2.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Query("code")
		token, err := config.Exchange(oauth2.NoContext, code)
		if err != nil {
			c.AbortWithStatus(http.StatusNetworkAuthenticationRequired)
			return
		}

		client := config.Client(oauth2.NoContext, token)
		resp, err := client.Get("https://api.github.com/user")

		if err != nil || resp.StatusCode != http.StatusOK {
			c.AbortWithStatus(http.StatusNetworkAuthenticationRequired)
			return
		}

		defer resp.Body.Close()

		var user map[string]interface{}

		err = json.NewDecoder(resp.Body).Decode(&user)

		if err != nil {
			c.AbortWithStatus(http.StatusNetworkAuthenticationRequired)
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
