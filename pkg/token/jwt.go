package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	issuer = "falcon"
)

type UserTokenClaims struct {
	jwt.StandardClaims
	UserID  int64  `json:"user_id"`
	LoginAt int64  `json:"login_at"`
	Via     string `json:"via"`
}

func ParseToken(tokenStr string, secretKey []byte) (bool, *UserTokenClaims) {
	token, _ := jwt.ParseWithClaims(tokenStr, &UserTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if token == nil {
		return false, nil
	}
	claims, ok := token.Claims.(*UserTokenClaims)
	if !ok || !token.Valid {
		return false, nil
	}
	if !claims.VerifyExpiresAt(time.Now().UTC().Unix(), true) {
		return false, nil
	}
	return true, claims
}

// GenerateToken 生成token
func GenerateToken(userID int64, via string, secretKey []byte) (string, error) {
	var tokenClaims UserTokenClaims

	current := time.Now().UTC()
	expiredAt := current.Add(24 * time.Hour).Unix()

	tokenClaims = UserTokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt,
			IssuedAt:  current.Unix(),
			Issuer:    issuer,
		},
		UserID:  userID,
		LoginAt: current.Unix(),
		Via:     via,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	return token.SignedString(secretKey)
}
