package redisenum

import "fmt"

const (
	UserEndPrefix = "falcon:user-end:"
)

var (
	loginUserCacheKeyFormat = UserEndPrefix + "login:%d"
)

func GetLoginUserCacheKey(userID int64) string {
	return fmt.Sprintf(loginUserCacheKeyFormat, userID)
}
