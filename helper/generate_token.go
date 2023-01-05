package helper

import (
	"absensi-karyawan-api/constant"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtClaim struct {
	UserID   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	Division string `json:"division"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// GenerateToken is used to create new token and will return token and time expired token
func GenerateToken(claim JwtClaim) (string, time.Time, error) {
	jwtKey := []byte(constant.JwtKey)
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := claim
	claims.StandardClaims = jwt.StandardClaims{ExpiresAt: expirationTime.Unix()}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, expirationTime, err
}

// GenerateLogoutToken is used to generate dummy user token and will return dummy token and time expired token
func GenerateLogoutToken() (string, time.Time, error) {
	jwtKey := []byte(constant.JwtKey)
	expirationTime := time.Now().Add(0)
	logoutToken := JwtClaim{
		UserID:   0,
		UserName: "emptyUser",
		Division: "emptyDivision",
		Role:     "emptyRole",
	}

	logoutToken.StandardClaims = jwt.StandardClaims{ExpiresAt: expirationTime.Unix()}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, logoutToken)
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, expirationTime, err
}
