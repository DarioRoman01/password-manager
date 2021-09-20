package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func Encrypt(key, text []byte) {
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
		return
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
		return
	}

	encrypted := gcm.Seal(nonce, nonce, text, nil)

	file, err := os.Create("claves.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	if _, err := file.Write(encrypted); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("claves encryptadas uwu")
}

func EncryptFile(name string, key []byte) {
	text, err := os.ReadFile(fmt.Sprintf("./%s", name))
	if err != nil {
		fmt.Println(err)
		return
	}

	Encrypt(key, text)
}

func Desencrypt(key []byte) ([]byte, error) {
	ciphertext, err := ioutil.ReadFile("./claves.txt")
	if err != nil {
		return nil, err
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, err
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)

}

func AddPassword(newPassword, key []byte) {
	content, err := Desencrypt(key)
	if err != nil {
		fmt.Println(err)
	}

	p := []byte(newPassword)
	content = append(content, p...)

	Encrypt(key, content)
}
