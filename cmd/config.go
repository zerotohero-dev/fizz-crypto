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
	"github.com/zerotohero-dev/fizz-app/pkg/app"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
)

func configureApp(e env.FizzEnv) {
	app.Configure(app.ConfigureOptions{
		AppName:           e.Crypto.ServiceName,
		DeploymentType:    e.Deployment.Type,
		HoneybadgerApiKey: e.Crypto.HoneybadgerApiKey,
		LogDestination:    e.Log.Destination,
		SanitizeFn: func() {
			e.Crypto.Sanitize()
			e.Log.Sanitize()
		},
	})
}
