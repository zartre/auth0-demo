package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func MustAuthen(c *gin.Context) {
	ses := sessions.Default(c)
	if ses.Get("profile") == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Next()
}
