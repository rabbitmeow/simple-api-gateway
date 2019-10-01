package main

import (
	"net/http"
	"net/http/httputil"
	"simple-api-gateway/controllers"
	"simple-api-gateway/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	serviceController := new(controllers.ServiceController)
	// authController := new(controllers.AuthController)
	middleware := new(middleware.Middleware)
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetString("mode") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:3008"}
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, service"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	router.NoRoute(middleware.ValidateToken(), func(c *gin.Context) {
		serviceReq := c.Request.Header.Get("service")
		if serviceReq == "" {
			c.JSON(400, gin.H{
				"status":  400,
				"message": "header not complete",
			})
		} else {
			targetURL := serviceController.Match(serviceReq)
			if targetURL == "" {
				c.JSON(404, gin.H{
					"status":  404,
					"message": "service not found",
				})
			} else {
				target := targetURL
				director := func(req *http.Request) {
					req.URL.Scheme = "http"
					req.URL.Host = target
				}
				proxy := &httputil.ReverseProxy{Director: director}
				proxy.ServeHTTP(c.Writer, c.Request)
			}
		}
	})
	router.Run(viper.GetString("server.host") + ":" + viper.GetString("server.port"))

}
