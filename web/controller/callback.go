package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/nzartre/auth0-demo/internal/authen"
)

func AuthCallbackHandler(auth *authen.Authenticator) gin.HandlerFunc {
	return func(c *gin.Context) {
		ses := sessions.Default(c)

		// Auth0 sends `state` in the query params.
		if c.Query("state") != ses.Get("oauth_state") {
			c.String(http.StatusBadRequest, "State mismatched")
			return
		}

		// Request an access token from Auth0
		authToken, err := auth.Exchange(c.Request.Context(), c.Query("code"))
		if err != nil {
			c.String(http.StatusUnauthorized, "Failed to exchange token: %v", err)
			return
		}

		idToken, err := auth.ParseIDToken(c.Request.Context(), authToken)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to parse ID Token: %v", err)
			return
		}

		profile := make(map[string]any)
		if err := idToken.Claims(&profile); err != nil {
			c.String(http.StatusInternalServerError, "Failed to parse claims: %v", err)
			return
		}

		ses.Set("access_token", authToken.AccessToken)
		ses.Set("profile", profile)
		if err := ses.Save(); err != nil {
			c.String(http.StatusInternalServerError, "Failed to save session: %v", err)
			return
		}

		c.Redirect(http.StatusTemporaryRedirect, "/profile")
	}
}
