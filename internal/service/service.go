/*
 *  \
 *  \\,
 *   \\\,^,.,,.                     Zero to Hero
 *   ,;7~((\))`;;,,               <zerotohero.dev>
 *   ,(@') ;)`))\;;',    stay up to date, be curious: learn
 *    )  . ),((  ))\;,
 *   /;`,,/7),)) )) )\,,
 *  (& )`   (,((,((;( ))\,
 */

package service

import (
	"context"
	"github.com/zerotohero-dev/fizz-entity/pkg/data"
	"time"
)

type Service interface {
	TokenCreate() (string, error)

	HashCreate(pwd string) (string, error)
	HashVerify(pass, hash string) bool

	JwtCreate(user data.User) string
	JwtVerify(authToken string) (valid bool, expiresAt int64, email string)
}

type Args struct {
	JwtKey           string
	AesPassphrase    string
	JwtExpiration    time.Duration
	RandomByteLength int
	BcryptHashRounds int
	IsDevelopment    bool
}

type service struct {
	args Args
	ctx  context.Context
}

func New(
	args Args,
	ctx context.Context,
) Service {
	return &service{
		args: args,
		ctx:  ctx,
	}
}
