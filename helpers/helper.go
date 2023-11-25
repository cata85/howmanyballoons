package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func Config() map[string]interface{} {
	arg := os.Args[1]
	filename, _ := filepath.Abs("config/config-" + arg + ".yml")
	yamlFile, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var config map[string]interface{}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	return config
}
func Get(this interface{}, key string) interface{} {
	return this.(map[string]interface{})[key]
}
func String(payload interface{}) string {
	var load string
	if pay, oh := payload.(string); oh {
		load = pay
	} else {
		load = ""
	}
	return load
}

func GetPassword(encKey string, iv string, encPass string) string {
	ciphertext, err := base64.StdEncoding.DecodeString(encPass)
	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher([]byte(encKey))
	if err != nil {
		panic(err)
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, []byte(iv))
	mode.CryptBlocks(ciphertext, ciphertext)

	password := ""
	for _, r := range string(ciphertext) {
		x := int(r)
		if x != 8 {
			password += string(r)
		}
	}
	return password
}
