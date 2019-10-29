package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"errors"
	"io"
)

func encryptStream(key string, iv []byte) cipher.Stream {
	block := newCipherBlock(key)
	return cipher.NewCFBEncrypter(block, iv)
}

// EncryptWriter will return a writer that will write encrypted data to
// the original writer.
func EncryptWriter(key string, w io.Writer) (*cipher.StreamWriter, error) {
	iv := make([]byte, aes.BlockSize)
	io.ReadFull(rand.Reader, iv)
	stream := encryptStream(key, iv)
	_, err := w.Write(iv)
	if err != nil {
		return nil, err
	}
	return &cipher.StreamWriter{S: stream, W: w}, nil
}

func decryptStream(key string, iv []byte) cipher.Stream {
	block := newCipherBlock(key)
	return cipher.NewCFBDecrypter(block, iv)
}

// DecryptReader will return a reader that will decrypt data from the
// provided reader and give the user a way to read that data as it if was
// not encrypted.
func DecryptReader(key string, r io.Reader) (*cipher.StreamReader, error) {
	iv := make([]byte, aes.BlockSize)
	n, err := r.Read(iv)
	//apart from nil check we should also ensure that
	//number of bytes that are read must be 16
	if n < len(iv) || err != nil {
		return nil, errors.New("encrypt: unable to read the full iv")
	}
	stream := decryptStream(key, iv)

	return &cipher.StreamReader{S: stream, R: r}, nil
}

//newCipherBlock return cipher block containing hashed version of key
func newCipherBlock(key string) cipher.Block {
	hasher := md5.New()
	hasher.Write([]byte(key))
	cipherKey := hasher.Sum(nil)
	block, _ := aes.NewCipher(cipherKey)
	return block
}
