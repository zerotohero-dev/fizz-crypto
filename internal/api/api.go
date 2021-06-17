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

package api

import (
	"github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/zerotohero-dev/fizz-app/pkg/app"
	"github.com/zerotohero-dev/fizz-crypto/internal/endpoint"
	"github.com/zerotohero-dev/fizz-crypto/internal/service"
	"github.com/zerotohero-dev/fizz-crypto/internal/transport"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
)

func route(router *mux.Router, handler *http.Server, method string, path string) {
	router.Methods(method).Path(path).Handler(handler)
}

func InitializeEndpoints(e env.FizzEnv, router *mux.Router) {
	svc := service.NewCryptoService(e)

	// Create a cryptographic hash.
	route(
		router, http.NewServer(
			endpoint.MakeHashCreateEndpoint(svc),
			app.ContentTypeValidatingMiddleware(transport.DecodeHashCreateRequest),
			app.EncodeResponse,
		),
		"POST", "/v1/hash",
	)

	// Verify the hash.
	route(
		router, http.NewServer(
			endpoint.MakeHashVerifyEndpoint(svc),
			app.ContentTypeValidatingMiddleware(transport.DecodeHashVerifyRequest),
			app.EncodeResponse,
		),
		"POST", "/v1/hash/verify",
	)

	// Create a JSON Web Token.
	route(
		router, http.NewServer(
			endpoint.MakeJwtCreateEndpoint(svc),
			app.ContentTypeValidatingMiddleware(transport.DecodeJwtCreateRequest),
			app.EncodeResponse,
		),
		"POST", "/v1/jwt",
	)

	// Verify the JSON Web Token.
	route(
		router, http.NewServer(
			endpoint.MakeJwtVerifyEndpoint(svc),
			app.ContentTypeValidatingMiddleware(transport.DecodeJwtVerifyRequest),
			app.EncodeResponse,
		),
		"POST", "/v1/jwt/verify",
	)

	// Create a random token.
	route(
		router, http.NewServer(
			endpoint.MakeTokenCreateEndpoint(svc),
			app.ContentTypeValidatingMiddleware(transport.DecodeTokenCreateRequest),
			app.EncodeResponse,
		),
		"GET", "/v1/token",
	)
}
