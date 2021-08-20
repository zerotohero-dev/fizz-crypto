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

package api

import (
	"context"
	"github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/zerotohero-dev/fizz-app/pkg/app"
	"github.com/zerotohero-dev/fizz-crypto/internal/endpoint"
	"github.com/zerotohero-dev/fizz-crypto/internal/service"
	"github.com/zerotohero-dev/fizz-crypto/internal/transport"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
)

var urls = struct {
	SecureHash       string
	SecureHashVerify string
	Jwt              string
	JwtVerify        string
	SecureToken      string
}{
	SecureHash:       "/v1/hash",
	SecureHashVerify: "/v1/hash/verify",
	Jwt:              "/v1/jwt",
	JwtVerify:        "/v1/jwt/verify",
	SecureToken:      "/v1/token",
}

func InitializeEndpoints(e env.FizzEnv, router *mux.Router) {
	svc := service.New(e, context.Background())

	prefix := e.Crypto.PathPrefix

	// Create a cryptographic hash.
	app.RoutePrefixedPath(
		http.NewServer(
			endpoint.MakeHashCreateEndpoint(svc),
			app.ContentTypeValidatingMiddleware(transport.DecodeHashCreateRequest),
			app.EncodeResponse,
		),
		router, "POST", prefix, urls.SecureHash,
	)

	// Verify the hash.
	app.RoutePrefixedPath(
		http.NewServer(
			endpoint.MakeHashVerifyEndpoint(svc),
			app.ContentTypeValidatingMiddleware(transport.DecodeHashVerifyRequest),
			app.EncodeResponse,
		),
		router, "POST", prefix, urls.SecureHashVerify,
	)

	// Create a JSON Web Token.
	app.RoutePrefixedPath(
		http.NewServer(
			endpoint.MakeJwtCreateEndpoint(svc),
			app.ContentTypeValidatingMiddleware(transport.DecodeJwtCreateRequest),
			app.EncodeResponse,
		),
		router, "POST", prefix, urls.Jwt,
	)

	// Verify the JSON Web Token.
	app.RoutePrefixedPath(
		http.NewServer(
			endpoint.MakeJwtVerifyEndpoint(svc),
			app.ContentTypeValidatingMiddleware(transport.DecodeJwtVerifyRequest),
			app.EncodeResponse,
		),
		router, "POST", prefix, urls.JwtVerify,
	)

	// Create a random token.
	app.RoutePrefixedPath(
		http.NewServer(
			endpoint.MakeTokenCreateEndpoint(svc),
			app.ContentTypeValidatingMiddleware(transport.DecodeTokenCreateRequest),
			app.EncodeResponse,
		),
		router, "GET", prefix, urls.SecureToken,
	)
}
