package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"time"
)

const (
	Claims = "claims"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	Id       string
	Username string
	Phone    string
	// 角色
	Role int32
}

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{SigningKey: []byte("fsaefwegergerge")}
}

func (j *JWT) GenerateJWT(claims *CustomClaims) (string, error) {

	claims.ExpiresAt = &jwt.NumericDate{Time: time.Now().Add(time.Duration(1000 * 3600 * 24))}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(j.SigningKey)
	if err != nil {
		zap.S().Error("生成JWT错误:" + err.Error())
		return "", err
	}
	return tokenStr, nil
}

func (j *JWT) ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	// 判断各种解析错误
	if err != nil {
		if result, ok := err.(jwt.ValidationError); ok {
			switch result.Errors {
			case jwt.ValidationErrorMalformed:
				return nil, errors.New("解析Token失败")
			case jwt.ValidationErrorExpired:
				return nil, errors.New("Token过期")
			case jwt.ValidationErrorNotValidYet:
				return nil, errors.New("Token不正确")
			default:
				return nil, errors.New("Token不合法")
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, errors.New("Token不合法")
	}

	return nil, errors.New("Token不合法")
}

// 刷新token
func (j *JWT) RefreshToken(tokenStr string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if err != nil {
		return "", err
	}

	// 合法就加一个小时
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(24 * time.Hour))
		return j.GenerateJWT(claims)
	}

	return "", errors.New("Token不合法")
}

func GetClaim(c *gin.Context) *CustomClaims {
	value, exists := c.Get(Claims)
	if !exists {
		return nil
	}
	return value.(*CustomClaims)
}
