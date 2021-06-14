/*
 *  \
 *  \\,
 *   \\\,^,.,,.                    “Zero to Hero”
 *   ,;7~((\))`;;,,               <zerotohero.dev>
 *   ,(@') ;)`))\;;',    stay up to date, be curious: learn
 *    )  . ),((  ))\;,
 *   /;`,,/7),)) )) )\,,
 *  (& )`   (,((,((;( ))\,
 */

package service

import (
	"crypto/aes"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/zerotohero-dev/fizz-entity/pkg/user"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
	"time"
)

type claims struct {
	User    []byte
	jwt.StandardClaims
}

func (c cryptoService) JwtCreate(user user.User) string {
	key := c.env.Crypto.JwtKey

	// TODO: to constants.
	expires := time.Now().Add(30 * 24 * time.Hour)

	email, err := aes.Encrypt([]byte(user.Email))

	if err != nil {
		LogErr(
			fmt.Sprintf("JwtSign: Error encrypting user email (%s).", user.Email),
			err.Error(),
		)

		return ""
	}

	cl := &claims{
		User:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expires.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cl)

	tokenString, err := token.SignedString([]byte(key))

	if err != nil {
		log.Err("JwtSign: Error computing signed string", err.Error())

		return ""
	}

	return tokenString
}

func (c cryptoService) JwtVerify(authToken string) (valid bool, expiresAt int64, email string) {
	panic("implement me")
}
