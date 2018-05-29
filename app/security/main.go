package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// EncryptWithAES ...
func EncryptWithAES(text string) (string, error) {

	block, err := aes.NewCipher([]byte("as"))

	if err != nil {
		return "", err
	}

	textToByte := []byte(text)

	cipherByte := make([]byte, aes.BlockSize+len(textToByte))

	iv := cipherByte[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {

		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherByte[aes.BlockSize:], textToByte)

	fmt.Println(base64.URLEncoding.EncodeToString(cipherByte))

	return base64.URLEncoding.EncodeToString(cipherByte), nil

}

// DecryptWithAES ...
func DecryptWithAES(crypto string) (string, error) {

	blocks, err := base64.URLEncoding.DecodeString(crypto)

	if err != nil {
		return "", err
	}

	key := []byte("86F7E437FAA5A7FCE15D1DDCB9EAEAEA377667B81B7FEEA3")
	block, err := aes.NewCipher(key)

	if err != nil {
		return "", err
	}

	if len(blocks) < aes.BlockSize {
		return "", err // CAMBIAR
	}

	iv := blocks[:aes.BlockSize]
	blocks = blocks[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(blocks, blocks)

	return fmt.Sprintf("%s", blocks), nil

}

/*
userPassword1 := "some user-provided password"

	// Generate "hash" to store from user password
	hash, err := bcrypt.GenerateFromPassword([]byte(userPassword1), bcrypt.DefaultCost)
	if err != nil {
		// TODO: Properly handle error
		log.Fatal(err)
	}
	fmt.Println("Hash to store:", string(hash))
	// Store this "hash" somewhere, e.g. in your database

	// After a while, the user wants to log in and you need to check the password he entered
	userPassword2 := "some user-provided password"
	hashFromDatabase := hash

	// Comparing the password with the hash
	if err := bcrypt.CompareHashAndPassword(hashFromDatabase, []byte(userPassword2)); err != nil {
		// TODO: Properly handle error
		log.Fatal(err)
	}

	fmt.Println("Password was correct!")

	fmt.Println(bcrypt.CompareHashAndPassword(hashFromDatabase, []byte(userPassword2)))


*/
