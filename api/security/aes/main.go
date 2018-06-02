package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/viper"
)

// Encrypt ...
func Encrypt(data string) string {

	key := fmt.Sprintf("%v", viper.Get("aes.key"))

	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		log.Panic(errors.New("Error ocurred:" + err.Error()))
		os.Exit(1)
	}

	cypher := []byte(data)

	blocks := make([]byte, aes.BlockSize+len(cypher))
	iv := blocks[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		log.Panic(errors.New("Error ocurred:" + err.Error()))
		os.Exit(1)
	}

	stream := cipher.NewCFBEncrypter(block, iv)

	stream.XORKeyStream(blocks[aes.BlockSize:], cypher)

	return base64.URLEncoding.EncodeToString(blocks)
}

// Decrypt ...
func Decrypt(cypher string) string {

	var key = fmt.Sprintf("%v", viper.Get("aes.key"))

	blocks, err := base64.URLEncoding.DecodeString(cypher)

	if err != nil {
		log.Panic(errors.New("Error ocurred:" + err.Error()))
		os.Exit(1)
	}

	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		log.Panic(errors.New("Error ocurred:" + err.Error()))
		os.Exit(1)
	}

	if len(blocks) < aes.BlockSize {
		log.Panic(errors.New("Error ocurred:" + err.Error()))
		os.Exit(1)
	}

	iv := blocks[:aes.BlockSize]
	blocks = blocks[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(blocks, blocks)

	return fmt.Sprintf("%s", blocks)

}
