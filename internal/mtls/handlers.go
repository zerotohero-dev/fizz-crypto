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
	"fmt"
	"github.com/zerotohero-dev/fizz-crypto/internal/service"
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
	"log"
	"net"
)

func handleConnection(conn net.Conn, svc service.Service) {
	fmt.Println("handle connection")
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			panic(err.Error())
		}
	}(conn)

	req, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Printf("Error reading incoming data: %v", err)
		return
	}

	apiRequest := &reqres.MtlsApiRequest{}

	_ = json.Unmarshal([]byte(req), apiRequest)

	log.Printf("Client says: %q", req)
	log.Println(apiRequest.Endpoint) // /v1/token
	log.Println(apiRequest.Method)   // "GET"

	token, _ := svc.TokenCreate()
	res := &reqres.TokenCreateResponse{
		Token: token,
	}

	serialized, _ := json.Marshal(res)

	// Send a response back to the client
	if _, err = conn.Write([]byte(string(serialized) + "\n")); err != nil {
		log.Printf("Unable to send response: %v", err)
		return
	}
}

func handleError(err error) {
	log.Printf("Unable to accept connection: %v", err)
}
