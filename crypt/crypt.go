package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
)

// genPassPhrase This function take key[string] as a parameter and converts it into byte array of length 16.
func genPassPhrase(key string) []byte {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hasher.Sum(nil)
}

// Encrypt Encrypt will take in a key and plain text and returns
// hex representation of ciphertext
func Encrypt(textToEncrypt, key string) (string, error) {

	plaintext := []byte(textToEncrypt)
	passPhrase := genPassPhrase(key)

	// Error will never be called as we are passing 16 byte array always
	block, err := aes.NewCipher(passPhrase)
	if err != nil {
		return "", errors.New("invalid passphrase")
	}
	//Initialization Vector
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", errors.New("error occurred in initialization vector(IV)")
	}

	// Generating cipher text
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return hex.EncodeToString(ciphertext), nil
}

// Decrypt Decrypt will take in a key and hex represented ciphertext and converts
// into original plaintext
func Decrypt(cipherText, key string) (string, error) {

	passPhrase := genPassPhrase(key)
	ciphertext, _ := hex.DecodeString(cipherText)

	block, err := aes.NewCipher(passPhrase)
	if err != nil {
		return "", errors.New("invalid passphrase")
	}

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext), nil

}
