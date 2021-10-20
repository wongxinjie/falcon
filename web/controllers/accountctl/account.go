package accountctl

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"falcon/config"
	"falcon/enum/apienum"
	"falcon/instance/loginst"
	"falcon/mapper/db/loginaccountmapper"
	"falcon/pkg/token"
	"falcon/service/loginaccountsvc"
	"falcon/web/middleware"
)

func MobileLoginApi(c *gin.Context) {

	var requestBody MobileLoginRequest
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, apienum.ErrorInvalidArgument)
		return
	}
	if !requestBody.IsValid() {
		c.AbortWithStatusJSON(http.StatusBadRequest, apienum.ErrorInvalidArgument)
		return
	}

	ifr := middleware.GetInfra(c)
	ctx := context.Background()
	db := ifr.DB

	loginMapper := loginaccountmapper.New(db)
	loginAccount, err := loginMapper.PasswordLogin(ctx, requestBody.Mobile, requestBody.Password)
	if err != nil {
		loginst.Inst().Warn("passwordLogin error", zap.Any("requestBody", requestBody))
		c.AbortWithStatusJSON(http.StatusInternalServerError, apienum.DBError)
		return
	}
	if loginAccount == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, apienum.ErrorAccountOrPasswordNotMatch)
		return
	}

	loginCacheUser := &loginaccountsvc.LoginUserCacheData{
		UserID:  loginAccount.Id,
		LoginAt: time.Now().UTC().Unix(),
		Via:     requestBody.Via,
		Status:  0,
	}
	err = loginaccountsvc.CacheLoginUser(c, ifr, loginCacheUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, apienum.DBError)
		return
	}

	loginToken, err := token.GenerateToken(loginAccount.Id, requestBody.Via, []byte(config.UserJwtSecret))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, apienum.ErrorUnKnown)
		return
	}

	response := &LoginResponse{Token: loginToken}
	c.JSON(http.StatusOK, response)
}

func AccountApi(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, apienum.ErrorUnKnown)
		return
	}

	ifr := middleware.GetInfra(c)
	accountMapper := loginaccountmapper.New(ifr.DB)
	account, err := accountMapper.OneByID(c, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, apienum.DBError)
		return
	}

	response := AccountResponse{
		UserID:       account.Id,
		UserName:     account.Phone,
		RegisteredAt: account.CreatedAt.Unix(),
		Level:        0,
	}

	c.JSON(http.StatusOK, response)
}
