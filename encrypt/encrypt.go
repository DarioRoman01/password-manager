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

func Encrypt(key, text []byte) error {
	c, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	encrypted := gcm.Seal(nonce, nonce, text, nil)

	file, err := os.Create("claves.txt")
	if err != nil {
		return err
	}

	if _, err := file.Write(encrypted); err != nil {
		return err
	}

	return nil
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
	ciphertext, err := ioutil.ReadFile("../pwd")
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

func AddPassword(newPassword, key []byte) error {
	content, err := Desencrypt(key)
	if err != nil {
		return err
	}

	p := []byte(newPassword)
	content = append(content, p...)

	return Encrypt(key, content)
}
