package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const (
	signingKey = "qrkj#%@FNSAzpZ!@M<24FjH" // placeholder, i need to change it to secret
	tokenName  = "token"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie(tokenName)
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				c.Redirect(http.StatusSeeOther, "/")
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading cookie"})
			}
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(signingKey), nil
		})

		if err != nil || !token.Valid {
			c.Redirect(http.StatusSeeOther, "/")
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*jwt.StandardClaims)
		if !ok {
			c.Redirect(http.StatusSeeOther, "/")
			c.Abort()
			return
		}

		email := claims.Subject
		role := claims.Audience

		ctx := context.WithValue(c.Request.Context(), "email", email)
		ctx = context.WithValue(ctx, "role", role)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
