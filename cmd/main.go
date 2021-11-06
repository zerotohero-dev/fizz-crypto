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

package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/zerotohero-dev/fizz-app/pkg/app"
	"github.com/zerotohero-dev/fizz-crypto/internal/api"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
	"log"
	"net"
)

/*

TODO:
x Install SPIRE server locally.
* Launch crypto
* Launch idm
* register these two.
* try establishing mTLS between them.
/tmp/spire-server/private/api.sock
 */

//fmt.Println("in go func")
//// #region mTLS server
//ctx := context.Background()
//fmt.Println("creating the server…")
//listener, err := spiffetls.Listen(ctx, "tcp", "127.0.0.1:8443", tlsconfig.AuthorizeAny())
//fmt.Println(listener)
//if err != nil {
//	panic(err)
//} else {
//	fmt.Println("Everything is awesome!")
//}
//fmt.Println("out of go func")
//// #engregion mTLS server

// Client:
// https://github.com/spiffe/go-spiffe/blob/main/v2/examples/spiffe-tls/client/main.go

// Server:
// https://github.com/spiffe/go-spiffe/blob/main/v2/examples/spiffe-tls/server/main.go

// SPIFFE ID for IDM:     spiffe://fizzbuzz.pro/app/idm
// SPIFFE ID for Crypto:  spiffe://fizzbuzz.pro/app/crypto


func handleConnection(conn net.Conn) {
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
	log.Printf("Client says: %q", req)

	// Send a response back to the client
	if _, err = conn.Write([]byte("Hello client\n")); err != nil {
		log.Printf("Unable to send response: %v", err)
		return
	}
}

func handleError(err error) {
	log.Printf("Unable to accept connection: %v", err)
}

// TODO: this will likely be behind a service.
const (
	socketPath    = "unix:///tmp/spire-agent/public/api.sock"
	serverAddress = "localhost:55553"
)

func runSpireMtlSServer() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// TODO: maybe use "default" ?
		clientId := spiffeid.Must("fizzbuzz.pro", "app", "default")

		fmt.Println("runSpireMtlsServer", clientId)

		listener, err := spiffetls.ListenWithMode(ctx, "tcp", serverAddress,
			spiffetls.MTLSServerWithSourceOptions(
				tlsconfig.AuthorizeID(clientId),
				// tlsconfig.AuthorizeAny(),
				workloadapi.WithClientOptions(workloadapi.WithAddr(socketPath)),
			))

		if err != nil {
			log.Fatalf("Unable to create TLS listener: %v", err)
		}

		defer func(listener net.Listener) {
			err := listener.Close()
			if err != nil {
				panic(err.Error())
			}
		}(listener)

		for {
			conn, err := listener.Accept()
			if err != nil {
				go handleError(err)
			}
			go handleConnection(conn)
		}
}

func main() {
	e := *env.New()

	appEnv := e.Crypto
	svcName := appEnv.ServiceName

	app.Configure(e, svcName, appEnv.HoneybadgerApiKey, appEnv.Sanitize)

	r := mux.NewRouter()
	api.InitializeEndpoints(e, r)
	app.RouteHealthEndpoints(e.Crypto.PathPrefix, r)

	// TODO: in this setup ordering will probably matter. Think about how to
	// sort that out. As in: server needs to start first before the clients can
	// accept connections. Maybe instead of that, using an HTTPS restful tls
	// could be more practical. -- I need to think about this.
	go runSpireMtlSServer()

	app.ListenAndServe(e, svcName, "9012", r)
}
