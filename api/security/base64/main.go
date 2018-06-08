package aes

import (
	"encoding/base64"
	"log"
)

// Encrypt ...
func Encrypt(data string) string { return base64.StdEncoding.EncodeToString([]byte(data)) }

// Decrypt ...
func Decrypt(cypher string) string {

	data, err := base64.StdEncoding.DecodeString(cypher)

	if err != nil {
		log.Fatal("error:", err)
	}

	return base64.URLEncoding.EncodeToString(data)

}
