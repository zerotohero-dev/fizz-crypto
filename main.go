/*
 *  \
 *  \\,
 *   \\\,^,.,,.                    “Zero to Hero”
 *   ,;7~((\))`;;,,               <zerotohero.dev>
 *   ,(@') ;)`))\;;',    stay up to date, be curious: learn
 *    )  . ),((  ))\;,
 *   /;`,,/7),)) )) )\,,
 *  (& )`   (,((,((;( ))\,
 */

package main

import (
	"fmt"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
)

func main() {
	e := env.New()

	e.SanitizeCrypto()

	fmt.Println("crypto port", e.Crypto.PortSvcCrypto)
}
