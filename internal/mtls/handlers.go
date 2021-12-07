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
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/zerotohero-dev/fizz-crypto/internal/service"
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
	"net"
)

func handleSecureHashVerify(conn net.Conn, svc service.Service) error {
	return nil
}

func handleJwt(conn net.Conn, svc service.Service) error {
	return nil
}

func handleSecureHash(conn net.Conn, svc service.Service) error {
	return nil
}

func handleSecureToken(conn net.Conn, svc service.Service) error {
	token, _ := svc.TokenCreate()
	res := &reqres.TokenCreateResponse{
		Token: token,
	}

	// TODO: handle serialization errors.
	serialized, _ := json.Marshal(res)

	if _, err := conn.Write([]byte(string(serialized) + "\n")); err != nil {
		return errors.Wrap(err, "Unable to send a response")
	}

	return nil
}

func handleUnknown(conn net.Conn, svc service.Service) error {
	return nil
}
