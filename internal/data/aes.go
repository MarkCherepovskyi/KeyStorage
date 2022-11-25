package data

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"log"
)

func Encryption(text, key []byte) ([]byte, error) {

	c, err := aes.NewCipher(key)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(gcm.Seal(nonce, nonce, text, nil))
	return gcm.Seal(nonce, nonce, text, nil), nil
}

func Decryption(ciphertext, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		fmt.Println(err)
		return nil, err
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(string(plaintext)) // todo delete it
	return plaintext, nil
}
