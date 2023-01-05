package helper

import (
	"absensi-karyawan-api/constant"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// Auth is used to get token from
func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		bearerToken := context.GetHeader(constant.Authorization)
		if !strings.Contains(bearerToken, constant.Bearer) {
			res := GetResponse(http.StatusUnauthorized, false, constant.MSG_INVALID_TOKEN, nil)
			context.JSON(http.StatusUnauthorized, res)
			context.Abort()
			return
		}

		tokenString := ""
		arrayToken := strings.Split(bearerToken, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		err, id := ValidateUserToken(tokenString)
		if err != nil {
			res := GetResponse(http.StatusUnauthorized, false, err.Error(), nil)
			context.JSON(http.StatusUnauthorized, res)
			context.Abort()
			return
		}

		context.Set("id", id)
		context.Next()
	}
}

// ValidateUserToken is used to validation
func ValidateUserToken(signedToken string) (err error, ID int64) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(constant.JwtKey), nil
		},
	)
	if err != nil {
		return err, ID
	}
	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		err = errors.New(constant.MsgParseErr)
		return err, ID
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New(constant.MsgTokenExpired)
		return err, ID
	}

	ID = claims.UserID

	return nil, ID
}
