package users

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

var key = []byte("example key 1234")

func encrypt(data string) string {

	block, err := aes.NewCipher(key)

	if err != nil {
		panic(err)
	}

	dataAsByte := []byte(data)
	cipherData := make([]byte, aes.BlockSize+len(dataAsByte))

	iv := cipherData[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherData[aes.BlockSize:], dataAsByte)

	return base64.URLEncoding.EncodeToString(cipherData)

}
