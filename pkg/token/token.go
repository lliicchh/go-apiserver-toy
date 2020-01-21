package token

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lliicchh/apiserver/pkg/errno"
	"github.com/spf13/viper"
	"time"
)

type Context struct {
	ID       uint64
	Username string
}

// gets the token from the header
// and pass it to the Parse function
func ParseRequest(c *gin.Context) (*Context, error) {
	header := c.Request.Header.Get("Authorization")
	secret := viper.GetString("jwt_secret")

	if len(header) == 0 {
		return &Context{}, errno.ErrTokenInvalid// nil is wrong, the correct error to return is MissingHeader
	}
	var t string
	fmt.Sscanf(header, "Bearer %s", &t)
	return Parse(t, secret)

}

// validates the secret format
func secreFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil

	}

}

// validates the token with given secret,
// and returns the context if the token is valid
func Parse(t string, secret string) (*Context, error) {
	ctx := &Context{}

	token, err := jwt.Parse(t, secreFunc(secret))

	if err != nil {
		return ctx, err
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.ID = uint64(claims["id"].(float64))
		ctx.Username = claims["username"].(string)
		return ctx, nil
	} else {
		return ctx, err
	}
}

func Sign(ctx *gin.Context, c Context, secret string) (string, error) {
	if secret == "" {
		secret = viper.GetString("jwt_secret")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       c.ID,
		"username": c.Username,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, err
}
