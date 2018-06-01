package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"

	"github.com/spf13/viper"
)

// AES ...
type AES struct {
	encoded string
	decoded string
	err     error
}

var key = []byte(fmt.Sprintf("%v", viper.Get("aes.key")))

// Encrypt ...
func Encrypt(data string) *AES {

	block, err := aes.NewCipher(key)

	if err != nil {
		return &AES{err: errors.New("Error ocurred:" + err.Error())}
	}

	cypher := []byte(data)

	blocks := make([]byte, aes.BlockSize+len(cypher))
	iv := blocks[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return &AES{err: errors.New("Error ocurred:" + err.Error())}
	}

	stream := cipher.NewCFBEncrypter(block, iv)

	stream.XORKeyStream(blocks[aes.BlockSize:], cypher)

	return &AES{encoded: base64.URLEncoding.EncodeToString(blocks)}
}

// Decrypt ...
func Decrypt(cypher string) *AES {

	blocks, err := base64.URLEncoding.DecodeString(cypher)

	if err != nil {
		return &AES{err: errors.New("Error ocurred:" + err.Error())}
	}

	block, err := aes.NewCipher(key)

	if err != nil {
		return &AES{err: errors.New("Error ocurred:" + err.Error())}
	}

	if len(blocks) < aes.BlockSize {
		return &AES{err: errors.New("AES blocksize invalid")}
	}

	iv := blocks[:aes.BlockSize]
	blocks = blocks[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(blocks, blocks)

	return &AES{decoded: fmt.Sprintf("%s", blocks)}

}
