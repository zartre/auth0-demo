package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ProfileHandler(c *gin.Context) {
	ses := sessions.Default(c)
	profile := ses.Get("profile")

	c.HTML(http.StatusOK, "profile.html", profile)
}
