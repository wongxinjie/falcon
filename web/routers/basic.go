package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusNotFound, "not found")
		return
	}
}
