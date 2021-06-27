#!/usr/bin/env zsh

#  \
#  \\,
#   \\\,^,.,,.                    “Zero to Hero”
#   ,;7~((\))`;;,,               <zerotohero.dev>
#   ,(@') ;)`))\;;',    stay up to date, be curious: learn
#    )  . ),((  ))\;,
#   /;`,,/7),)) )) )\,,
#  (& )`   (,((,((;( ))\,

TAG="0.0.10"

# shellcheck disable=SC1090
source ~/.zprofile

docker run -e FIZZ_CRYPTO_SVC_PORT="$FIZZ_CRYPTO_SVC_PORT" \
-e FIZZ_CRYPTO_JWT_KEY="$FIZZ_CRYPTO_JWT_KEY" \
-e FIZZ_CRYPTO_RANDOM_BYTE_LENGTH="$FIZZ_CRYPTO_RANDOM_BYTE_LENGTH" \
-e FIZZ_CRYPTO_BCRYPT_HASH_ROUNDS="$FIZZ_CRYPTO_BCRYPT_HASH_ROUNDS" \
-e FIZZ_CRYPTO_AES_PASSPHRASE="$FIZZ_CRYPTO_AES_PASSPHRASE" \
-e FIZZ_DEPLOYMENT_TYPE="$FIZZ_DEPLOYMENT_TYPE" \
-e FIZZ_CRYPTO_HONEYBADGER_API_KEY="$FIZZ_CRYPTO_HONEYBADGER_API_KEY" \
-e FIZZ_LOG_DESTINATION="$FIZZ_LOG_DESTINATION" \
-e FIZZ_CRYPTO_JWT_EXPIRY_HOURS="$FIZZ_CRYPTO_JWT_EXPIRY_HOURS" \
zerotohero-dev/fizz-crypto:$TAG
