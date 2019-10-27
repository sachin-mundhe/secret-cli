package vault

import (
	"crypt"
	"errors"
)

// NewVault It returns new valut
func NewVault(encryptionKey string) Vault {
	return Vault{
		encodingKey: encryptionKey,
		keyValues:   make(map[string]string),
	}
}

// Vault It defines structure for vault
type Vault struct {
	// Key or Passphrase
	encodingKey string

	// Map to set a cipher text with some key
	keyValues map[string]string
}

// Get Get metjod takes in key name and check whether there is cipher text with that key name.
// And if it is present then it decrypts it into plain text and returns it back
func (v *Vault) Get(keyName string) (string, error) {
	cipherText, ok := v.keyValues[keyName]
	if !ok {
		return "", errors.New("No secret found")
	}
	plainText, err := crypt.Decrypt(cipherText, v.encodingKey)
	if err != nil {
		return "", errors.New("Error while decrypting ciphertext")
	}
	return plainText, nil
}

// Set Set method takes in key name and plain text to encode, and converts it into cipher text
func (v *Vault) Set(keyName, textToEncode string) error {
	encryptedValue, err := crypt.Encrypt(textToEncode, v.encodingKey)
	if err != nil {
		return err
	}
	v.keyValues[keyName] = encryptedValue
	return nil
}
