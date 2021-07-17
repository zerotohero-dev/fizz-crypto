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
	"github.com/zerotohero-dev/fizz-entity/pkg/data"
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
)

// MakeJwtCreateEndpoint creates an endpoint that generates a JSON Web Token
// when provided an email. Note that no verification is done whether an active
// user with that email exists. The assumption is that the service that’s calling
// this endpoint has already done its due diligence and already checked that an
// active user with the email exists in the database.
func MakeJwtCreateEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		gr, ok := request.(reqres.ContentTypeProblemRequest)

		if ok {
			return reqres.JwtCreateResponse{
				Err: gr.Err,
			}, nil
		}

		req := request.(reqres.JwtCreateRequest)

		if req.Err != "" {
			return reqres.JwtCreateResponse{
				Token: "",
				Err:   req.Err,
			}, nil
		}

		u := data.User{
			Info: data.Info{
				Email: req.Email,
			},
		}

		token := svc.JwtCreate(u)

		return reqres.JwtCreateResponse{
			Token: token,
		}, nil
	}
}

func MakeJwtVerifyEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		gr, hasContentTypeIssues := request.(reqres.ContentTypeProblemRequest)

		if hasContentTypeIssues {
			return reqres.JwtVerifyResponse{
				Valid: false,
				Err:   gr.Err,
			}, nil
		}

		req := request.(reqres.JwtVerifyRequest)

		if req.Err != "" {
			return reqres.JwtVerifyResponse{
				Valid: false,
				Err:   "Error verifying token.",
			}, nil
		}

		valid, expires, email := svc.JwtVerify(req.Token)

		return reqres.JwtVerifyResponse{
			Valid:   valid,
			Email:   email,
			Expires: expires,
		}, nil
	}
}
