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

func handleJwtVerify(conn net.Conn, svc service.Service, body string) error {
	body, err := payload(body)
	if err != nil {
		return errors.Wrap(err, "handleJwtVerify: problem with unmarshal")
	}

	var jvr reqres.JwtVerifyRequest
	err = json.Unmarshal([]byte(body), &jvr)
	if err != nil {
		return errors.Wrap(err, "handleJwtVerify: problem with unmarshal")
	}

	valid, expiresAt, email := svc.JwtVerify(jvr.Token)
	return send(conn, reqres.JwtVerifyResponse{
		Valid: valid,
		Expires: expiresAt,
		Email: email,
	})
}

func handleJwt(conn net.Conn, svc service.Service, body string) error {
	body, err := payload(body)
	if err != nil {
		return errors.Wrap(err, "handleJwt: problem with unmarshal")
	}

	var jcr reqres.JwtCreateRequest
	err = json.Unmarshal([]byte(body), &jcr)
	if err != nil {
		return errors.Wrap(err, "handleJwt: problem with unmarshal")
	}

	token := svc.JwtCreate(data.User{Email: jcr.Email})
	return send(conn, reqres.JwtCreateResponse{
		Token: token,
	})
}

func handleSecureHashVerify(conn net.Conn, svc service.Service, body string) error {
	body, err := payload(body)
	if err != nil {
		return errors.Wrap(err, "handleSecureHashVerify: problem with unmarshal")
	}

	var hwr reqres.HashVerifyRequest
	err = json.Unmarshal([]byte(body), &hwr)
	if err != nil {
		return errors.Wrap(err, "handleSecureHashVerify: problem with unmarshal")
	}

	verified := svc.HashVerify(hwr.Value, hwr.Hash)
	return send(conn, reqres.HashVerifyResponse{
		Verified: verified,
	})
}

func handleSecureHash(conn net.Conn, svc service.Service, body string) error {
	body, err := payload(body)
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

func handleSecureToken(conn net.Conn, svc service.Service, body string) error {
	token, _ := svc.TokenCreate()
	res := &reqres.TokenCreateResponse{
		Token: token,
	}

	return send(conn, res)
}

func handleUnknown(conn net.Conn, svc service.Service, body string) error {
	log.Warning("Unknown request")
	return nil
}
