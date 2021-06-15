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
	"crypto/rand"
	"encoding/base64"
	"github.com/pkg/errors"
	"strconv"
)

func (c cryptoService) TokenCreate() (string, error) {
	sz, err := strconv.Atoi(c.env.Crypto.RandomByteLength)
	if err != nil {
		return "🦄", errors.Wrap(err, "TokenCreate: Problem converting random byte length")
	}

	token := make([]byte, sz)

	_, err = rand.Read(token)

	if err != nil {
		return "🦄", errors.Wrap(err, "TokenCreate: error creating random token")
	}

	encoded := base64.StdEncoding.EncodeToString(token)

	return encoded, nil
}
