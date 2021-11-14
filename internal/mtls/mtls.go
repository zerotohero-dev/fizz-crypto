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

package mtls

import (
	"github.com/zerotohero-dev/fizz-crypto/internal/service"
)

type SpireArgs struct {
	ServerAddress  string
	SocketPath     string
	AppPrefix      string
	AppNameDefault string
	AppNameCrypto  string
	AppNameIdm     string
	AppNameMailer  string
	AppTrustDomain string
}

func ListenAndServe(cryptoArgs service.Args, spireArgs SpireArgs) {
	runSpireMtlSServer(cryptoArgs, spireArgs)
}
