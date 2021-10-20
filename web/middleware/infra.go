package middleware

import (
	"github.com/gin-gonic/gin"

	"falcon/infra"
)

func SetApiInfra() gin.HandlerFunc {

	return func(c *gin.Context) {
		ifr := infra.Inst()
		if ifr == nil {
			panic("infra is nil")
		}
		c.Set(InfraKey, ifr)
		c.Next()
	}
}
