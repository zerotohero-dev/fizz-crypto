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
		handleSecureHashVerify(conn, svc, body)
	case apiEndpoint == endpoint.Crypto.Jwt && apiMethod == method.Post:
		handleJwt(conn, svc, body)
	case apiEndpoint == endpoint.Crypto.JwtVerify && apiMethod == method.Post:
		handleJwtVerify(conn, svc, body)
	case apiEndpoint == endpoint.Crypto.SecureHash && apiMethod == method.Post:
		handleSecureHash(conn, svc, body)
	case apiEndpoint == endpoint.Crypto.SecureToken && apiMethod == method.Get:
		handleSecureToken(conn, svc, body)
	default:
		handleUnknown(conn, svc, body)
	}
}