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
	"github.com/zerotohero-dev/fizz-entity/pkg/endpoint"
	"github.com/zerotohero-dev/fizz-entity/pkg/method"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
	"net"
)

func multiplex(
	apiEndpoint string,
	apiMethod method.Method,
	body string,
	conn net.Conn, svc service.Service,
) {
	switch {
	case apiEndpoint == endpoint.Crypto.SecureHashVerify && apiMethod == method.Post:
		err := handleSecureHashVerify(conn, svc, body)
		if err != nil {
			log.Err("error handling crypto secure hash verification %s", err.Error())
		}
	case apiEndpoint == endpoint.Crypto.Jwt && apiMethod == method.Post:
		err := handleJwt(conn, svc, body)
		if err != nil {
			log.Err("error handling jwt: %s", err.Error())
		}
	case apiEndpoint == endpoint.Crypto.JwtVerify && apiMethod == method.Post:
		err := handleJwtVerify(conn, svc, body)
		if err != nil {
			log.Err("error handling jwt: %s", err.Error())
		}
	case apiEndpoint == endpoint.Crypto.SecureHash && apiMethod == method.Post:
		err := handleSecureHash(conn, svc, body)
		if err != nil {
			log.Err("error handling secure hash: %s", err.Error())
		}
	case apiEndpoint == endpoint.Crypto.SecureToken && apiMethod == method.Get:
		err := handleSecureToken(conn, svc, body)
		if err != nil {
			log.Err("error handling secure token: %s", err.Error())
		}
	default:
		err := handleUnknown(conn, svc, body)
		if err != nil {
			log.Err("error handling unknown request: %s", err.Error())
		}
	}
}