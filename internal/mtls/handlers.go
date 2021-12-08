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
	"github.com/zerotohero-dev/fizz-entity/pkg/data"
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
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

func handleSecureHashVerify(conn net.Conn, svc service.Service, request string) error {
	return nil
}

func handleJwt(conn net.Conn, svc service.Service, request string) error {
	body, err := payload(request)
	if err != nil {
		return errors.Wrap(err, "handleJwt: problem with unmarshal")
	}

	var jcr reqres.JwtCreateRequest
	err = json.Unmarshal([]byte(body), &jcr)
	if err != nil {
		return errors.Wrap(err, "handleJwt: problem with unmarshal")
	}

	result := svc.JwtCreate(data.User{Email:jcr.Email})
	return send(conn, result)
}

func handleSecureHash(conn net.Conn, svc service.Service, request string) error {
	body, err := payload(request)
	if err != nil {
		return errors.Wrap(err, "handleSecureHash: problem with unmarshal")
	}

	var hcr reqres.HashCreateRequest
	err = json.Unmarshal([]byte(body), &hcr)
	if err != nil {
		return errors.Wrap(err, "handleSecureHash: problem with unmarshal")
	}

	hash, err := svc.HashCreate(hcr.Value)
	if err != nil {
		return errors.Wrap(err, "handleSecureHash: problem creating hash")
	}

	return send(conn, reqres.HashCreateResponse{
		Hash: hash,
	})
}

func handleSecureToken(conn net.Conn, svc service.Service) error {
	token, _ := svc.TokenCreate()
	res := &reqres.TokenCreateResponse{
		Token: token,
	}

	return send(conn, res)
}

func handleUnknown(conn net.Conn, svc service.Service) error {
	log.Warning("Unknown request")
	return nil
}
