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

	return request, nil
}
