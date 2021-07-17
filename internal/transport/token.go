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

package transport

import (
	"context"
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
	"net/http"
)

// DecodeTokenCreateRequest decodes a token that is used for email verification.
func DecodeTokenCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request reqres.TokenCreateRequest

	//if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	//	log.Err("decodeTokenRequest: error decoding: %s", err.Error())
	//
	//	request.Err = "decodeTokenRequest: Problem decoding JSON."
	//}

	return request, nil
}
