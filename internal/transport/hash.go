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

func DecodeHashCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request reqres.HashCreateRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Err("DecodeHashCreateRequest: error decoding: %s", err.Error())

		request.Err = "DecodeHashCreateRequest: Problem decoding JSON."
	}

	return request, nil
}

func DecodeHashVerifyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request reqres.HashVerifyRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Err("DecodeHashVerifyRequest: error decoding: %s", err.Error())

		request.Err = "DecodeHashVerifyRequest: Problem decoding JSON."
	}

	return request, nil
}
