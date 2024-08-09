package infrastructure

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Auth struct{
	JWTService JWT
}

func NewAuthMiddleware (jwt JWT)*Auth{
	return &Auth{
		JWTService: jwt,
	}
}
type AuthInterface interface{
	AuthenticationMiddleware() gin.HandlerFunc
}
func NewAuth(jwtService JWT) AuthInterface{
	return &Auth{
		JWTService: jwtService,
	}
}

func (authenticate *Auth) AuthenticationMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		authHeader := c.GetHeader("Authorization")
		if authHeader == ""{
			c.JSON(http.StatusUnauthorized, gin.H{
				"message" : "Authorization header required",
			})
			c.Abort()
			return
		}
		auth := strings.Split(authHeader, " ")
		if len(auth) != 2 || strings.ToLower(auth[0]) != "bearer"{
			c.JSON(http.StatusUnauthorized, gin.H{
				"message" : "Invalid authorization header.",
			})
			c.Abort()
			return
		}
		claims, err := authenticate.JWTService.ValidateToken(auth[1])
		if err != nil{
			c.JSON(http.StatusUnauthorized, gin.H{
				"message" : err,
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