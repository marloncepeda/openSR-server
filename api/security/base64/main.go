package aes

import (
	"encoding/base64"
	"log"
)

// Encrypt ...
func Encrypt(data string) string {

	return base64.StdEncoding.EncodeToString([]byte(data))

	/*
		key := fmt.Sprintf("%v", viper.Get("aes.key"))

		block, err := aes.NewCipher([]byte(key))

		if err != nil {
			log.Panic(errors.New("Error ocurred:" + err.Error()))
			os.Exit(1)
		}

		cypher := []byte(data)

		blocks := make([]byte, aes.BlockSize+len(cypher))
		iv := make([]byte, 16)

		fmt.Println(base64.URLEncoding.EncodeToString(iv))

		if _, err := io.ReadFull(rand.Reader, iv); err != nil {
			log.Panic(errors.New("Error ocurred:" + err.Error()))
			os.Exit(1)
		}

		stream := cipher.NewCFBEncrypter(block, iv)

		stream.XORKeyStream(blocks[aes.BlockSize:], cypher)

		return base64.URLEncoding.EncodeToString(blocks)
	*/
}

// Decrypt ...
func Decrypt(cypher string) string {

	data, err := base64.StdEncoding.DecodeString(cypher)

	if err != nil {
		log.Fatal("error:", err)
	}

	return base64.URLEncoding.EncodeToString(data)

	/*

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

		iv := make([]byte, 16)
		blocks = blocks[aes.BlockSize:]

		stream := cipher.NewCFBDecrypter(block, iv)

		stream.XORKeyStream(blocks, blocks)

		return fmt.Sprintf("%s", blocks)

	*/

}
