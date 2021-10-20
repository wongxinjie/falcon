package routers

import (
	"github.com/gin-gonic/gin"

	"falcon/web/controllers/accountctl"
	"falcon/web/controllers/poetryctl"
)

const (
	ExternalRouteGroup = "api"
)

const (
	externalPoetryApiGroupName = "/v1/poetry"
	externalUserApiGroupName   = "/v1/user"
)

func setExternalPoetryRoutes(vg *gin.RouterGroup) {
	g := vg.Group(externalPoetryApiGroupName)

	g.GET("/search", NoAuthChain(poetryctl.SearchApi)...)
	g.GET("/detail/:id", NoAuthChain(poetryctl.DetailApi)...)
}

func setExternalUserRoutes(vg *gin.RouterGroup) {
	g := vg.Group(externalUserApiGroupName)

	g.POST("/mobile/login", NoAuthChain(accountctl.MobileLoginApi)...)
	g.GET("/account", StrictAuthChain(accountctl.AccountApi)...)
}

func setExternalGroup(r *gin.Engine) {
	gr := r.Group(ExternalRouteGroup)

	setExternalPoetryRoutes(gr)
	setExternalUserRoutes(gr)
}

func SetRoutes(router *gin.Engine) {

	setExternalGroup(router)
}
