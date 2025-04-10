package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const (
	signingKey = "qrkj#%@FNSAzpZ!@M<24FjH" // Replace this with your actual secret key
	tokenName  = "token"                   // Name of the cookie storing the token
)

// AuthMiddleware is the middleware function to verify JWT tokens stored in cookies.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Extract the token from the cookie
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

		// 2. Parse and validate the token
		token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			// Ensure the signing method is correct
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

		// 3. Extract claims from the token
		claims, ok := token.Claims.(*jwt.StandardClaims)
		if !ok {
			c.Redirect(http.StatusSeeOther, "/")
			c.Abort()
			return
		}

		// 4. Pass user information to the next handler
		email := claims.Subject
		role := claims.Audience

		// Store user information in the context for downstream handlers
		ctx := context.WithValue(c.Request.Context(), "email", email)
		ctx = context.WithValue(ctx, "role", role)
		c.Request = c.Request.WithContext(ctx)

		// Proceed to the next handler
		c.Next()
	}
}
