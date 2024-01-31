package tjwt

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	smodel "github.com/octoposprime/op-be-shared/pkg/model"
	tconfig "github.com/octoposprime/op-be-shared/tool/config"
)

type Claims struct {
	UserId string `json:"user_id"`
	Data   any    `json:"data"`
	jwt.StandardClaims
}

func NewClaims(userId string, data any) *Claims {
	return &Claims{
		UserId: userId,
		Data:   data,
	}
}

type RefreshClaims struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

func (c *Claims) GenerateJWT() (tokenStr string, refreshTokenStr string, err error) {
	c.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	var mySigningKey = []byte(tconfig.GetJwtConfigInstance().Jwt.SecretKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenStr, err = token.SignedString(mySigningKey)
	if err != nil {
		return
	}

	refreshClaims := &RefreshClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 24).Unix(),
		},
		UserId: c.UserId,
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenStr, err = refreshToken.SignedString(mySigningKey)
	if err != nil {
		return
	}
	return
}

func DecodeJWT(token string) (userId string, data any, err error) {
	token = strings.Replace(token, "Bearer ", "", 1)
	var mySigningKey = []byte(tconfig.GetJwtConfigInstance().Jwt.SecretKey)
	jwtToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return
	}
	if claims, ok := jwtToken.Claims.(*Claims); ok && jwtToken.Valid {
		if !claims.StandardClaims.VerifyExpiresAt(time.Now().Unix(), true) {
			err = smodel.ErrorTokenExpired
			return
		}
		userId = claims.UserId
		data = claims.Data
		return
	} else {
		err = smodel.ErrorDecodeJwtFailed
		return
	}
}

func DecodeRefreshJWT(token string) (userId string, err error) {
	token = strings.Replace(token, "Bearer ", "", 1)
	var mySigningKey = []byte(tconfig.GetJwtConfigInstance().Jwt.SecretKey)
	jwtToken, err := jwt.ParseWithClaims(token, &RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return
	}
	if claims, ok := jwtToken.Claims.(*RefreshClaims); ok && jwtToken.Valid {
		if !claims.StandardClaims.VerifyExpiresAt(time.Now().Unix(), true) {
			err = smodel.ErrorTokenExpired
			return
		}
		userId = claims.UserId
		return
	} else {
		err = smodel.ErrorDecodeJwtFailed
		return
	}
}
