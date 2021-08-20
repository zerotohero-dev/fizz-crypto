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

package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/zerotohero-dev/fizz-crypto/internal/service"
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
)

func MakeHashCreateEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		gr, hasContentTypeProblem := request.(reqres.ContentTypeProblemRequest)
		if hasContentTypeProblem {
			return reqres.HashCreateResponse{
				Err: gr.Err,
			}, nil
		}

		req := request.(reqres.HashCreateRequest)
		if req.Err != "" {
			return reqres.HashCreateResponse{
				Hash: "",
				Err:  req.Err,
			}, nil
		}

		if req.Value == "" {
			return reqres.HashCreateResponse{
				Hash: "",
				Err:  "makeHashEndpoint: attempting to hash an empty string",
			}, nil
		}

		hash, err := svc.HashCreate(req.Value)
		if err != nil {
			return reqres.HashCreateResponse{
				Hash: "",
				Err:  "makeHashEndpoint: error hashing value",
			}, nil
		}

		return reqres.HashCreateResponse{
			Hash: hash,
		}, nil
	}
}

func MakeHashVerifyEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		gr, ok := request.(reqres.ContentTypeProblemRequest)

		if ok {
			return reqres.HashVerifyResponse{
				Err: gr.Err,
			}, nil
		}

		req := request.(reqres.HashVerifyRequest)

		if req.Err != "" {
			return reqres.HashVerifyResponse{
				Verified: false,
				Err:      req.Err,
			}, nil
		}

		verified := svc.HashVerify(req.Value, req.Hash)

		return reqres.HashVerifyResponse{
			Verified: verified,
		}, nil
	}
}
