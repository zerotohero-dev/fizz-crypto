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
	"context"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/zerotohero-dev/fizz-crypto/internal/service"
	"log"
	"net"
)

func runSpireMtlSServer(svcArgs service.Args, spireArgs SpireArgs) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	trustDomain := spireArgs.AppTrustDomain
	appPrefix := spireArgs.AppPrefix
	idmAppName := spireArgs.AppNameIdm
	mailerAppName := spireArgs.AppNameMailer
	anyAppName := spireArgs.AppNameDefault

	var ids []spiffeid.ID
	if svcArgs.IsDevelopment {
		anyId, _ := spiffeid.New(trustDomain, appPrefix, anyAppName)
		ids = []spiffeid.ID{anyId}
	} else {
		appId, _ := spiffeid.New(trustDomain, appPrefix, idmAppName)
		mailerId, _ := spiffeid.New(trustDomain, appPrefix, mailerAppName)
		ids = []spiffeid.ID{appId, mailerId}
	}

	listener, err := spiffetls.ListenWithMode(ctx, "tcp", spireArgs.ServerAddress,
		spiffetls.MTLSServerWithSourceOptions(
			tlsconfig.AuthorizeOneOf(ids...),
			workloadapi.WithClientOptions(workloadapi.WithAddr(spireArgs.SocketPath)),
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

	svc := service.New(svcArgs, ctx)

	for {
		conn, err := listener.Accept()
		if err != nil {
			go handleError(err)
		}
		go handleConnection(conn, svc)
	}
}
