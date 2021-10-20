package middleware

import (
	"github.com/gin-gonic/gin"

	"falcon/infra"
)

const (
	InfraKey  = "infra"
	UserIdKey = "user_id"
	ViaKey    = "via"
)

func GetInfra(c *gin.Context) *infra.Infra {
	if c == nil {
		return nil
	}

	v, ok := c.Get(InfraKey)
	if !ok {
		return nil
	}

	ifr, ok := v.(*infra.Infra)
	if !ok {
		return nil
	}
	return ifr
}

func GetUserID(c *gin.Context) int64 {
	if c == nil {
		return 0
	}

	v, ok := c.Get(UserIdKey)
	if !ok {
		return 0
	}
	userID, ok := v.(int64)
	if !ok {
		return 0
	}
	return userID
}
