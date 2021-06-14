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
	"github.com/zerotohero-dev/fizz-entity/pkg/user"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
)

// import "strconv"

// var saltRounds, _ = strconv.Atoi(env.Env.BcryptHashRounds)

type CryptoService interface {
	TokenCreate() (string, error)

	HashCreate(pwd string) (string, error)
	HashVerify(pass, hash string) bool

	JwtCreate(user user.User) string
	JwtVerify(authToken string) (valid bool, expiresAt int64, email string)
}

type cryptoService struct{
	env env.FizzEnv
}


func NewCryptoService(e env.FizzEnv) CryptoService {
	return cryptoService{
		env: e,
	}
}