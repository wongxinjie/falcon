package routers

import (
	"github.com/gin-gonic/gin"

	"falcon/web/middleware"
)

func NoAuthChain(handler gin.HandlerFunc) gin.HandlersChain {
	return gin.HandlersChain{
		middleware.SetApiInfra(),
		handler,
	}
}

func StrictAuthChain(handler gin.HandlerFunc) gin.HandlersChain {
	return gin.HandlersChain{
		middleware.SetApiInfra(),
		middleware.LoginRequired(),
		handler,
	}
}
