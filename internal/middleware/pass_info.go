package middleware

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func PassUserData() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie(tokenName)
		if err != nil {
			ctx := c.Request.Context()
			ctx = context.WithValue(ctx, "email", "")
			ctx = context.WithValue(ctx, "role", "")
			c.Request = c.Request.WithContext(ctx)

			c.Next()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				fmt.Println("Unexpected signing method")
				c.Next()
				return nil, nil
			}
			return []byte(signingKey), nil
		})

		if err != nil || !token.Valid {
			fmt.Println("Invalid or unparsable token:", err)

			ctx := c.Request.Context()
			ctx = context.WithValue(ctx, "email", "")
			ctx = context.WithValue(ctx, "role", "")
			c.Request = c.Request.WithContext(ctx)

			c.Next()
			return
		}

		claims, ok := token.Claims.(*jwt.StandardClaims)
		if !ok {
			fmt.Println("Failed to extract claims")

			ctx := c.Request.Context()
			ctx = context.WithValue(ctx, "email", "")
			ctx = context.WithValue(ctx, "role", "")
			c.Request = c.Request.WithContext(ctx)

			c.Next()
			return
		}

		email := claims.Subject
		role := claims.Audience

		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, "email", email)
		ctx = context.WithValue(ctx, "role", role)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
