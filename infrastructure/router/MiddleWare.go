package router

import (
	aAuth "aBet/adapters/auth"
	iAuth "aBet/infrastructure/auth"
	"encoding/json"
	"errors"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

/*
Hàm lấy thông tin token và kiểm tra
- hợp lệ
- Xác thực
*/
func getMiddleWareConfig(authObject *aAuth.AuthObject) middleware.JWTConfig {
	return middleware.JWTConfig{
		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
			keyFunc := func(t *jwt.Token) (interface{}, error) {
				signingKey := iAuth.GetPublicKey()
				return []byte(signingKey), nil
			}
			// var re = regexp.MustCompile(`[a-z0-9A-X]*::`)
			// newAuth := re.ReplaceAllString(auth, "")
			token, _ := jwt.Parse(auth, keyFunc)
			if !token.Valid {
				return nil, errors.New("invalid token")
			}
			claims, _ := json.Marshal(token.Claims)

			*authObject, _ = iAuth.NewAuthObject(claims, "Bearer "+token.Raw)

			return token, nil
		},
	}
}
