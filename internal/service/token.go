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
	"crypto/rand"
	"encoding/base64"
	"github.com/pkg/errors"
)

func (c service) TokenCreate() (string, error) {
	sz := c.args.RandomByteLength

	token := make([]byte, sz)

	_, err := rand.Read(token)

	if err != nil {
		return "🦄", errors.Wrap(err, "TokenCreate: error creating random token")
	}

	encoded := base64.StdEncoding.EncodeToString(token)

	return encoded, nil
}
