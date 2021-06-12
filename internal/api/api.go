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
)


func route(router *mux.Router, handler *http.Server, method string, path string) {
	router.Methods(method).Path(path).Handler(handler)
}

func InitializeEndpoints(router *mux.Router) {

	// Create a cryptographic hash.
	route(
		router, http.NewServer(nil, nil, nil),
		"POST", "/v1/hash",
	)

	// Verify the hash.
	route(
		router, http.NewServer(nil, nil, nil),
		"GET", "/v1/hash/{hashToVerify}",
	)

	// Create a JSON Web Token.
	route(
		router, http.NewServer(nil, nil, nil),
		"POST", "/v1/jwt",
	)

	// Verify the JSON Web Token.
	route(
		router, http.NewServer(nil, nil, nil),
		"GET", "/v1/jwt/{jwtToVerify}",
	)

	// Create a random token.
	route(
		router, http.NewServer(nil, nil, nil),
		"GET", "/v1/token",
	)
}