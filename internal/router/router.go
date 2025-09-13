package router

import (
	"github.com/gin-gonic/gin"
	web "github.com/nzartre/auth0-demo/web/controller"
)

func New() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("web/template/*")

	r.GET("/", web.HomeHandler)

	return r
}
