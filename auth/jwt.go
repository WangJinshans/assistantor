package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

const (
	Secret          = "123#111" //salt
	ExpireTime      = 3600      //token expire time
	ErrorServerBusy = "server is busy"
	ErrorReLogin    = "relogin"
)

type JWTClaims struct {
	jwt.StandardClaims
	UserId string
}

//生成 jwt token
func GenerateToken(userId string) (string, error) {
	claims := &JWTClaims{
		UserId: userId,
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", errors.New(ErrorServerBusy)
	}
	return signedToken, nil
}

//验证jwt token
func VerifyToken(ctx *gin.Context) (*JWTClaims, error) {
	strToken := ctx.Request.Header.Get("token")
	token, err := jwt.ParseWithClaims(strToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return nil, errors.New(ErrorServerBusy)
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, errors.New(ErrorReLogin)
	}
	if err = token.Claims.Valid(); err != nil {
		return nil, errors.New(ErrorReLogin)
	}
	return claims, nil
}

func Refresh(c *gin.Context) (string, error) {
	claims, _ := VerifyToken(c)
	return GenerateToken(claims.UserId)
}
