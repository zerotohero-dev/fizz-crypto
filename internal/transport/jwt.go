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

package transport

import (
	"context"
	"encoding/json"
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
	"net/http"
)

func DecodeJwtCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request reqres.JwtCreateRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Err("decodeJwtSignRequest: error decoding: %s", err.Error())

		request.Err = "decodeJwtSignRequest: Problem decoding JSON."
	}

	return request, nil
}

func DecodeJwtVerifyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request reqres.JwtVerifyRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Err("decodeJwtVerifyRequest: error decoding: %s", err.Error())

		request.Err = "decodeJwtVerifyRequest: Problem decoding JSON."
	}

	return request, nil
}
