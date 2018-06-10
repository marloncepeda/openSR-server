package security

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// Encrypt ...
func Encrypt(data string) string {

	key := fmt.Sprintf("%v", viper.Get("encrypt.key"))
	keys := strings.Split(key, ",")

	//Convert string to byte.
	dataAsByte := []byte(data)

	//Convert byte to hex
	hexData := hex.EncodeToString(dataAsByte)

	//Convert hex to base64
	base := base64.StdEncoding.EncodeToString([]byte(hexData))

	//Merge with the encryption key
	return keys[0] + base + keys[1]
}

// Decrypt ...
func Decrypt(cypher string) string {

	key := fmt.Sprintf("%v", viper.Get("encrypt.key"))
	keys := strings.Split(key, ",")

	var replacer = strings.NewReplacer(keys[0], "", keys[1], "")

	//Delete the keys
	base := replacer.Replace(cypher)

	// Convert base64 to hex
	hexData, _ := base64.StdEncoding.DecodeString(base)

	//Convert byte slide to string hex
	str := fmt.Sprintf("%s", hexData)

	//Convert hex to initial []byte
	dataAsByte, _ := hex.DecodeString(str)

	//Convert byte to string
	data := string(dataAsByte[:])

	return data
}
