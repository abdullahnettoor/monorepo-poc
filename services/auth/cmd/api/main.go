package main

import (
	"context"
	"log"
	"net/http"
	"os"

	authmiddleware "github.com/abdullahnettoor/monorepo-poc/libs/go/auth-middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	issuerURL := os.Getenv("OIDC_ISSUER_URL")
	clientID := os.Getenv("OIDC_CLIENT_ID")

	if issuerURL == "" {
		// Default to local Keycloak for development
		issuerURL = "http://localhost:8080/realms/monorepo"
	}

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// protected group
	protected := r.Group("/")
	protected.Use(authMiddleware(issuerURL, clientID))
	{
		protected.GET("/profile", func(c *gin.Context) {
			claims, _ := c.Get("claims")
			c.JSON(http.StatusOK, gin.H{
				"message": "You are authorized",
				"claims":  claims,
			})
		})
	}

	log.Println("Starting Auth Service on :8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func authMiddleware(issuerURL, clientID string) gin.HandlerFunc {
	// We initialize the validator once. In production, handle error/retries more gracefully.
	// context.Background is used here but ideally should be managed better.
	validator, err := authmiddleware.NewTokenValidator(context.Background(), issuerURL, clientID)
	if err != nil {
		log.Printf("WARNING: Failed to initialize OIDC validator: %v. Auth will fail.", err)
	}

	return func(c *gin.Context) {
		if validator == nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Auth provider unavailable"})
			return
		}

		tokenString, err := authmiddleware.ExtractTokenFromHeader(c.Request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid authorization token"})
			return
		}

		token, err := validator.ValidateToken(c.Request.Context(), tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "details": err.Error()})
			return
		}

		// Store claims in context for handlers to use
		var claims map[string]interface{}
		if err := token.Claims(&claims); err == nil {
			c.Set("claims", claims)
		}

		c.Next()
	}
}
