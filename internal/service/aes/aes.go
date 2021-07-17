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

package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func Encrypt(aesPassphrase string, data []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(aesPassphrase))

	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)

	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, data, nil), nil
}

func Decrypt(aesPassphrase string, data []byte) ([]byte, error) {
	key := []byte(aesPassphrase)
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)

	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()

	nonce, ct := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ct, nil)

	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
