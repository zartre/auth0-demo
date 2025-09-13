package router

import (
	"encoding/gob"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/nzartre/auth0-demo/internal/authen"
	web "github.com/nzartre/auth0-demo/web/controller"
)

func New(auth *authen.Authenticator) *gin.Engine {
	r := gin.Default()

	// We will store JSON in the cookie.
	gob.Register(map[string]any{})

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("auth_session", store))

	r.LoadHTMLGlob("web/template/*")

	r.GET("/", web.HomeHandler)
	r.GET("/callback", web.AuthCallbackHandler(auth))
	r.GET("/login", web.LoginHandler(auth))
	r.GET("/profile", web.ProfileHandler)

	return r
}
