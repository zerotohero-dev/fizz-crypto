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
	"github.com/zerotohero-dev/fizz-crypto/internal/mtls"
	"github.com/zerotohero-dev/fizz-crypto/internal/service"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
)

func listenAndServeApp(e env.FizzEnv) {
	go func() {
		svcName := e.Crypto.ServiceName

		r := mux.NewRouter()
		// api.InitializeEndpoints(e, r)

		app.RouteHealthEndpoints(e.Crypto.PathPrefix, r)
		app.ListenAndServe(e, svcName, e.Crypto.Port, r)
	}()
}

func listenAndServeMtls(e env.FizzEnv) {
	// go func() {
	mtls.ListenAndServe(service.Args{
		JwtKey:           e.Crypto.JwtKey,
		AesPassphrase:    e.Crypto.AesPassphrase,
		JwtExpiration:    e.Crypto.JwtExpiration,
		RandomByteLength: e.Crypto.RandomByteLength,
		BcryptHashRounds: e.Crypto.BcryptHashRounds,
		IsDevelopment:    e.Deployment.Type == env.Development,
	},
		mtls.SpireArgs{
			ServerAddress:  e.Crypto.MtlsServerAddress,
			SocketPath:     e.Spire.SocketPath,
			AppTrustDomain: e.Spire.AppTrustDomainFizz,
			AppPrefix:      e.Spire.AppPrefixFizz,
			AppNameDefault: e.Spire.AppNameFizzDefault,
			AppName:        e.Crypto.ServiceName,
			AppNameIdm:     e.Idm.ServiceName,
			AppNameMailer:  e.Mailer.ServiceName,
		},
	)
	// }()
}
