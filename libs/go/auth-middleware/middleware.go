package authmiddleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
)

type TokenValidator struct {
	verifier *oidc.IDTokenVerifier
}

// NewTokenValidator initializes the OIDC provider and verifier
func NewTokenValidator(ctx context.Context, issuerURL string, clientID string) (*TokenValidator, error) {
	provider, err := oidc.NewProvider(ctx, issuerURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get provider: %w", err)
	}

	config := &oidc.Config{
		ClientID: clientID,
	}

	// If ClientID is empty, we skip client ID check (useful if multiple clients access same resource)
	if clientID == "" {
		config.SkipClientIDCheck = true
	}

	return &TokenValidator{
		verifier: provider.Verifier(config),
	}, nil
}

// ValidateToken parses and validates the raw JWT string
func (v *TokenValidator) ValidateToken(ctx context.Context, rawToken string) (*oidc.IDToken, error) {
	return v.verifier.Verify(ctx, rawToken)
}

// ExtractTokenFromHeader helper to get token from Authorization header
func ExtractTokenFromHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("authorization header missing")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", fmt.Errorf("invalid authorization header format")
	}

	return parts[1], nil
}
