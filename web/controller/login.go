package controller

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/nzartre/auth0-demo/internal/authen"
)

func LoginHandler(auth *authen.Authenticator) gin.HandlerFunc {
	return func(c *gin.Context) {
		state, err := GenerateRandomState()
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Error generating state: %v", err))
			return
		}

		// The browser will need the same state for subsequent requests.
		ses := sessions.Default(c)
		ses.Set("oauth_state", state)
		if err := ses.Save(); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Error saving session: %v", err))
			return
		}

		c.Redirect(http.StatusTemporaryRedirect, auth.AuthCodeURL(state))
	}
}

// GenerateRandomState creates a cryptographically secure random string
// suitable for use as an OAuth 2.0 state parameter.
func GenerateRandomState() (string, error) {
	// This length provides sufficient entropy (256 bits)
	// to make guessing the state infeasible.
	b := make([]byte, 32)

	_, err := rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}

	return base64.URLEncoding.EncodeToString(b), nil
}
