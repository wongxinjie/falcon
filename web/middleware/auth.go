package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"falcon/config"
	"falcon/enum/apienum"
	token2 "falcon/pkg/token"
	"falcon/service/loginaccountsvc"
)

func LoginRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := c.Request
		token := r.Header.Get("Authorization")
		if len(token) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, apienum.ErrorUnAuthorized)
			return
		}

		ok, claims := token2.ParseToken(token, []byte(config.UserJwtSecret))
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, apienum.ErrorUnKnown)
			return
		}

		ifr := GetInfra(c)
		if ifr == nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, apienum.ErrorUnKnown)
			return
		}
		loginUser, err := loginaccountsvc.FetchCachedLoginUser(c, ifr, claims.UserID, claims.Via)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, apienum.ErrorUnKnown)
			return
		}

		if loginUser.UserID == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, apienum.ErrorUnAuthorized)
			return
		}
		if claims.LoginAt < loginUser.LoginAt {
			c.AbortWithStatusJSON(http.StatusUnauthorized, apienum.ErrorUnAuthorized)
			return
		}

		c.Set(UserIdKey, claims.UserID)
		c.Set(ViaKey, claims.Via)
	}
}
