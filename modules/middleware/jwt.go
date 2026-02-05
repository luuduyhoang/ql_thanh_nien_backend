package middleware

import (
	// "github.com/dgrijalva/jwt-go"
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatus(401)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatus(401)
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte("SECRET_KEY"), nil
		})

		if err != nil || !token.Valid {
			log.Println("JWT error:", err)
			c.AbortWithStatus(401)
			return
		}

		log.Printf("Token claims: %+v", token.Claims)

		c.Set("claims", token.Claims)
		c.Next()
	}
}


// func JWTMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		tokenString := c.GetHeader("Authorization")
// 		if tokenString == "" {
// 			c.AbortWithStatus(401)
// 			return
// 		}

// 		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
// 			return []byte("SECRET_KEY"), nil
// 		})

// 		log.Printf("hihihihihi")

// 		if err != nil || !token.Valid {
// 			c.AbortWithStatus(401)
// 			return
// 		}

// 		log.Printf("Token claims: %v", token.Claims)

// 		c.Set("claims", token.Claims)
// 		c.Next()
// 	}
// }
