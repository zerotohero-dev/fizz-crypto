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
	"github.com/zerotohero-dev/fizz-crypto/internal/service"
	"github.com/zerotohero-dev/fizz-entity/pkg/data"
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
	"github.com/zerotohero-dev/fizz-mtls/pkg/mtls/ext"
	"net"
)

func handleJwtVerify(conn net.Conn, svc service.Service, body string) {
	body, err := ext.Payload(body)
	if err != nil {
		log.Err("handleJwtVerify: problem with unmarshal: %s", err.Error())
		sendErr := ext.Send(conn, reqres.JwtVerifyResponse{
			Err: "handleJwtVerify: problem with unmarshal",
		})
		if sendErr != nil {
			log.Err("handleJwtVerify: could not send: %s", sendErr.Error())
		}
		return
	}

	var jvr reqres.JwtVerifyRequest
	err = json.Unmarshal([]byte(body), &jvr)
	if err != nil {
		log.Err("handleJwtVerify: problem with unmarshal: %s", err.Error())
		sendErr := ext.Send(conn, reqres.JwtVerifyResponse{
			Err: "handleJwtVerify: problem with unmarshal",
		})
		if sendErr != nil {
			log.Err("handleJwtVerify: could not send: %s", sendErr.Error())
		}
		return
	}

	valid, expiresAt, email := svc.JwtVerify(jvr.Token)
	sendErr := ext.Send(conn, reqres.JwtVerifyResponse{
		Valid: valid,
		Expires: expiresAt,
		Email: email,
	})
	if sendErr != nil {
		log.Err("handleJwtVerify: could not send: %s", sendErr.Error())
	}
}

func handleJwt(conn net.Conn, svc service.Service, body string) {
	body, err := ext.Payload(body)
	if err != nil {
		log.Err("handleJwt: problem with unmarshal: %s", err.Error())
		sendErr := ext.Send(conn, reqres.JwtCreateResponse{
			Err: "handleJwt: problem with unmarshal",
		})
		if sendErr != nil {
			log.Err("handleJwt: could not send: %s", sendErr.Error())
		}
		return
	}

	var jcr reqres.JwtCreateRequest
	err = json.Unmarshal([]byte(body), &jcr)
	if err != nil {
		log.Err("handleJwt: problem with unmarshal: %s", err.Error())

		sendErr := ext.Send(conn, reqres.JwtCreateResponse{
			Err: "handleJwt: problem with unmarshal",
		})
		if sendErr != nil {
			log.Err("handleJwt: could not send: %s", sendErr.Error())
		}
		return
	}

	token := svc.JwtCreate(data.User{Email: jcr.Email})
	sendErr := ext.Send(conn, reqres.JwtCreateResponse{
		Token: token,
	})
	if sendErr != nil {
		log.Err("handleJwt: could not send: %s", sendErr.Error())
	}
}

func handleSecureHashVerify(conn net.Conn, svc service.Service, body string) {
	body, err := ext.Payload(body)

	if err != nil {
		log.Err("handleSecureHashVerify: unmarshal problem: %s", err.Error())

		sendErr := ext.Send(conn, reqres.HashVerifyResponse{
			Verified: false,
			Err:      "handleSecureHashVerify: problem with unmarshal",
		})

		if sendErr != nil {
			log.Err("handleSecureHashVerify: could not send: %s", sendErr.Error())
		}

		return
	}

	var hwr reqres.HashVerifyRequest
	err = json.Unmarshal([]byte(body), &hwr)
	if err != nil {
		log.Err("handleSecureHashVerify: unmarshal problem: %s", err.Error())
		sendErr := ext.Send(conn, reqres.HashVerifyResponse{
			Verified: false,
			Err:      "handleSecureHashVerify: problem with unmarshal",
		})

		if sendErr != nil {
			log.Err("handleSecureHashVerify: could not send: %s", sendErr.Error())
		}

		return
	}

	verified := svc.HashVerify(hwr.Value, hwr.Hash)
	sendErr := ext.Send(conn, reqres.HashVerifyResponse{
		Verified: verified,
	})
	if sendErr != nil {
		log.Err("handleSecureHashVerify: could not send: %s", sendErr.Error())
	}
}

func handleSecureHash(conn net.Conn, svc service.Service, body string) {
	body, err := ext.Payload(body)
	if err != nil {
		sendErr := ext.Send(conn, reqres.HashCreateResponse{
			Err: "handleSecureHash: problem with unmarshal",
		})
		if sendErr != nil {
			log.Err("handleSecureHash: could not send: %s", sendErr.Error())
		}
		return
	}

	var hcr reqres.HashCreateRequest
	err = json.Unmarshal([]byte(body), &hcr)
	if err != nil {
		log.Err("handleSecureHash: problem with unmarshal: %s", err.Error())
		sendErr := ext.Send(conn, reqres.HashCreateResponse{
			Err: "handleSecureHash: problem with unmarshal",
		})
		if sendErr != nil {
			log.Err("handleSecureHash: could not send: %s", sendErr.Error())
		}
		return
	}

	hash, err := svc.HashCreate(hcr.Value)
	if err != nil {
		log.Err("handleSecureHash: problem creating hash: %s", err.Error())
		sendErr := ext.Send(conn, reqres.HashCreateResponse{
			Hash: hash,
		})
		if sendErr != nil {
			log.Err("handleSecureHash: could not send: %s", sendErr.Error())
		}
		return
	}

	sendErr := ext.Send(conn, reqres.HashCreateResponse{
		Hash: hash,
	})
	if sendErr != nil {
		log.Err("handleSecureHash: could not send: %s", sendErr.Error())
	}
}

func handleSecureToken(conn net.Conn, svc service.Service, body string) {
	token, _ := svc.TokenCreate()
	res := &reqres.TokenCreateResponse{
		Token: token,
	}

	sendErr := ext.Send(conn, res)
	if sendErr != nil {
		log.Err("handleSecureToken: could not send: %s", sendErr.Error())
	}
}

func handleUnknown(conn net.Conn, svc service.Service, body string) {
	log.Warning("Unknown request")
}
