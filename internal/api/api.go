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
	"github.com/zerotohero-dev/fizz-crypto/internal/service"
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
			transport.DecodeHashRequest,
			transport.EncodeResponse,
		),
		"POST", "/v1/hash",
	)

	// Verify the hash.
	route(
		router, http.NewServer(
			endpoint.MakeHashVerifyEndpoint(svc),
			transport.DecodeVerifyHashRequest,
			transport.EncodeResponse,
		),
		"GET", "/v1/hash/{hashToVerify}",
	)

	// Create a JSON Web Token.
	route(
		router, http.NewServer(
			endpoint.MakeJwtCreateEndpoint(svc),
			transport.DecodeJwtCreateRequest,
			transport.EncodeResponse,
		),
		"POST", "/v1/jwt",
	)

	// Verify the JSON Web Token.
	route(
		router, http.NewServer(
			endpoint.MakeJwtVerifyEndpoint,
			transport.DecodeJtwVerifyRequesty,
			transport.EncodeResponse,
		),
		"GET", "/v1/jwt/{jwtToVerify}",
	)

	// Create a random token.
	route(
		router, http.NewServer(
			enpoint.MakeTokenCreateEndpoint,
			transport.DecodeTokenCreateRequest,
			transport.EncodeResponse,
		),
		"GET", "/v1/token",
	)
}