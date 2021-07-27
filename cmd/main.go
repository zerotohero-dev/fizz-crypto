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

package main

import (
	"github.com/gorilla/mux"
	"github.com/zerotohero-dev/fizz-app/pkg/app"
	"github.com/zerotohero-dev/fizz-crypto/internal/api"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
)

const appName = "fizz-crypto"

func main() {
	e := *env.New()

	appEnv := e.Crypto

	// TODO: maybe instead of passing two separate parameters, create a
	// new interface `SanitizableEnv` that has a `Sanitize()` method by
	// contract, and pass that one to this method instead.
	app.Configure(e, appName, appEnv.HoneybadgerApiKey, appEnv.Sanitize)

	r := mux.NewRouter()
	api.InitializeEndpoints(e, r)
	app.RouteHealthEndpoints(e.Crypto.PathPrefix, r)

	app.ListenAndServe(e, appName, appEnv.Port, r)
}
