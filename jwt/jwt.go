package jwt

import (
	"fmt"
	"net/http"
	"time"

	gojwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	Payload map[string]interface{} `json:"payload"`
	gojwt.StandardClaims
}

type JWT struct {
	SecretKey     string
	ExpiredHour   uint64
	TokenHeadName string
	Authenticator func(*gin.Context) (map[string]interface{}, error)
	Verification  func(map[string]interface{}) (bool, error)
}

type JWTResponse struct {
	Expire       time.Time `json:"expire"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
}

func (jwt *JWT) LoginHandler(c *gin.Context) {
	payload, err := jwt.Authenticator(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}
	token, err := jwt.genToken(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code": http.StatusInternalServerError,
			"msg":  "Internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, token)
}

func (jwt *JWT) genToken(payload map[string]interface{}) (JWTResponse, error) {
	var res JWTResponse
	now := time.Now()
	exp := now.Add(time.Hour * time.Duration(jwt.ExpiredHour))
	claims := Claims{
		payload,
		gojwt.StandardClaims{
			ExpiresAt: exp.Unix(),
		},
	}

	result := gojwt.NewWithClaims(gojwt.SigningMethodHS256, claims)
	token, err := result.SignedString([]byte(jwt.SecretKey))
	if err != nil {
		return res, err
	}

	res.Token = token
	res.Expire = exp

	exp = now.Add(time.Hour * time.Duration(jwt.ExpiredHour+72))
	claims = Claims{
		payload,
		gojwt.StandardClaims{
			ExpiresAt: exp.Unix(),
		},
	}
	result = gojwt.NewWithClaims(gojwt.SigningMethodHS256, claims)
	refreshToken, err := result.SignedString([]byte(jwt.SecretKey))
	if err != nil {
		return res, err
	}
	res.RefreshToken = refreshToken
	return res, err
}

func (jwt *JWT) parseToken(ss string) (map[string]interface{}, error) {
	token, err := gojwt.ParseWithClaims(ss, &Claims{}, func(token *gojwt.Token) (interface{}, error) {
		return []byte(jwt.SecretKey), nil
	})

	if err != nil {
		if ve, ok := err.(*gojwt.ValidationError); ok {
			if ve.Errors&gojwt.ValidationErrorMalformed != 0 {
				return nil, fmt.Errorf("Invalid token")
			} else if ve.Errors&(gojwt.ValidationErrorExpired) != 0 {
				return nil, fmt.Errorf("Token expired")
			}
		}
		return nil, fmt.Errorf("Invalid token")
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("Invalid token")
	}
	return claims.Payload, nil
}
