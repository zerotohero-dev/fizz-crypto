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
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
	"net"
)

func send(conn net.Conn, result interface{}) error {
	serialized, _ := json.Marshal(result)
	if _, err := conn.Write([]byte(string(serialized) + "\n")); err != nil {
		return errors.Wrap(err, "Unable to send a response")
	}

	return nil
}

func payload(request string) (string, error) {
	var req reqres.MtlsApiRequest

	err := json.Unmarshal([]byte(request), &req)
	if err != nil {
		return "", errors.Wrap(err, "payload: problem with unmarshal")
	}

	body := req.Body
	return body, nil
}