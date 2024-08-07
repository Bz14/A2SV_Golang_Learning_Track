package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)
var jwtSecret = []byte("secret")

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context){
		authHeader := c.GetHeader("Authorization")
		if authHeader == ""{
			c.JSON(http.StatusUnauthorized, gin.H{
				"error" : "Authorization header required.",
			})
			c.Abort()
		}
		auth := strings.Split(authHeader, " ")
		fmt.Println(auth)
		if len(auth) != 2 || auth[0] != "Bearer"{
			c.JSON(http.StatusUnauthorized, gin.H{
				"error" : "Invalid authorization header.",
			})
			c.Abort()
			return
		}
		token, err := jwt.Parse(auth[1], func(token *jwt.Token)(interface{}, error){
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok{
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})
		if err != nil || !token.Valid{
			c.JSON(http.StatusUnauthorized, gin.H{
				"error" : "Invalid authorization header.",
			})
			c.Abort()
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok{
			c.JSON(http.StatusUnauthorized, gin.H{
				"error" : "Can't access claims.",
			})
			c.Abort()
			return
		}
		c.Set("user_id" , claims["user_id"])
		c.Set("username", claims["username"])
		c.Set("role", claims["role"])
		c.Next()

		// Can check the expiration time of the token if it is valid or not
	}
}