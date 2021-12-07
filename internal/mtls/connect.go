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
	"bufio"
	"encoding/json"
	"github.com/zerotohero-dev/fizz-crypto/internal/service"
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
	"net"
)

func handleConnection(conn net.Conn, svc service.Service) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			panic(err.Error())
		}
	}(conn)

	req, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Info("Error reading incoming data %v", err)
		return
	}

	apiRequest := &reqres.MtlsApiRequest{}

	_ = json.Unmarshal([]byte(req), apiRequest)

	log.Info("Client says: %q %s %s", req, apiRequest.Endpoint, apiRequest.Method)

	multiplex(apiRequest.Endpoint, apiRequest.Method, conn, svc)
}

func handleError(err error) {
	log.Info("Unable to accept connection: %v", err)
}
