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
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spiffe/go-spiffe/v2/spiffetls"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/zerotohero-dev/fizz-app/pkg/app"
	"github.com/zerotohero-dev/fizz-crypto/internal/api"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
)

/*

TODO:
x Install SPIRE server locally.
* Launch crypto
* Launch idm
* register these two.
* try establishing mTLS between them.
/tmp/spire-server/private/api.sock
 */

func main() {
	e := *env.New()

	go func() {
		fmt.Println("in go func")
		// #region mTLS server
		ctx := context.Background()
		fmt.Println("creating the server…")
		listener, err := spiffetls.Listen(ctx, "tcp", "127.0.0.1:8443", tlsconfig.AuthorizeAny())
		fmt.Println(listener)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Everything is awesome!")
		}
		fmt.Println("out of go func")
		// #engregion mTLS server
	}()

	appEnv := e.Crypto
	svcName := appEnv.ServiceName

	app.Configure(e, svcName, appEnv.HoneybadgerApiKey, appEnv.Sanitize)

	r := mux.NewRouter()
	api.InitializeEndpoints(e, r)
	app.RouteHealthEndpoints(e.Crypto.PathPrefix, r)

	app.ListenAndServe(e, svcName, appEnv.Port, r)
}
