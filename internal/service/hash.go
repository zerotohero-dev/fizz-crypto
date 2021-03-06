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
	"golang.org/x/crypto/bcrypt"
)

func (c service) HashCreate(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(pwd), c.args.BcryptHashRounds)

	return string(bytes), err
}

func (c service) HashVerify(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))

	return err == nil
}
