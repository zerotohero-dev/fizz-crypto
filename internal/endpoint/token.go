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

package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/zerotohero-dev/fizz-crypto/internal/service"
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
)

func MakeTokenCreateEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		gr, ok := request.(reqres.ContentTypeProblemRequest)

		if ok {
			return reqres.TokenCreateResponse{
				Err: gr.Err,
			}, nil
		}

		req := request.(reqres.TokenCreateRequest)

		if req.Err != "" {
			return reqres.TokenCreateResponse{
				Token: "",
				Err:   req.Err,
			}, nil
		}

		token, err := svc.TokenCreate()

		if err != nil {
			log.Err("MakeTokenCreateEndpoint: %s", err.Error())

			return reqres.TokenCreateResponse{
				Token: "",
				Err:   "MakeTokenCreateEndpoint: Error generating token",
			}, nil
		}

		return reqres.TokenCreateResponse{
			Token: token,
		}, nil
	}
}
