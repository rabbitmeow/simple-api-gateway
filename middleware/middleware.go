package middleware

import (
	"fmt"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//Middleware is
type Middleware struct{}

//ValidateToken is used for validate the jwt token
func (w *Middleware) ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.Request.Header.Get("Authorization")
		if bearer != "" {
			strSplit := strings.Split(bearer, " ")
			if strSplit[0] == "Bearer" && strSplit[1] != "" {
				secretKey := viper.GetString("token.jwt_server_key")
				token, err := jwt.Parse(strSplit[1], func(token *jwt.Token) (interface{}, error) {
					if jwt.GetSigningMethod("HS512") != token.Method {
						return nil, fmt.Errorf("Unexpected signing method")
					}

					return []byte(secretKey), nil
				})
				_, ok := token.Claims.(jwt.MapClaims)
				if ok && token.Valid && err == nil {
					c.Next()
				} else {
					c.JSON(401, gin.H{
						"status":  401,
						"message": "unauthorized",
					})
					c.AbortWithStatus(401)
				}
			} else {
				c.JSON(401, gin.H{
					"status":  401,
					"message": "unauthorized",
				})
				c.AbortWithStatus(401)
			}
		} else {
			c.JSON(401, gin.H{
				"status":  401,
				"message": "unauthorized",
			})
			c.AbortWithStatus(401)
		}
	}
}
