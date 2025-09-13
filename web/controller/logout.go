package controller

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LogoutHandler(c *gin.Context) {
	endpoint, err := url.Parse(fmt.Sprintf("https://%s/v2/logout", os.Getenv("AUTH0_DOMAIN")))
	if err != nil {
		c.String(http.StatusInternalServerError, "Error parsing logout URL: %v", err)
		return
	}

	returnTo, err := url.Parse(fmt.Sprintf("http://localhost:%s", os.Getenv("APP_PORT")))
	if err != nil {
		c.String(http.StatusInternalServerError, "Error parsing return URL: %v", err)
		return
	}

	params := make(url.Values)
	params.Add("returnTo", returnTo.String())
	params.Add("client_id", os.Getenv("AUTH0_CLIENT_ID"))
	endpoint.RawQuery = params.Encode()

	ses := sessions.Default(c)
	ses.Delete("auth_session")
	ses.Save()

	c.Redirect(http.StatusTemporaryRedirect, endpoint.String())
}
