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

	app.Configure(e, appName, e.Crypto.HoneybadgerApiKey, e.SanitizeCrypto)

	r := mux.NewRouter()
	api.InitializeEndpoints(r)
	app.RouteHealthEndpoints(r)

	app.ListenAndServe(e, appName, e.Crypto.Port, nil)
}
