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
	"github.com/zerotohero-dev/fizz-entity/pkg/data"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
)

type Service interface {
	TokenCreate() (string, error)

	HashCreate(pwd string) (string, error)
	HashVerify(pass, hash string) bool

	JwtCreate(user data.User) string
	JwtVerify(authToken string) (valid bool, expiresAt int64, email string)
}

type service struct {
	env env.FizzEnv
}

func New(e env.FizzEnv) Service {
	return &service{
		env: e,
	}
}
