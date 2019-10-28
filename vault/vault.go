package vault

import (
	"bytes"
	"crypt"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// NewVault It returns new valut
func NewVault(encryptionKey, filePath string) *Vault {
	return &Vault{
		encodingKey: encryptionKey,
		filePath:    filePath,
	}
}

// Vault It defines structure for vault
type Vault struct {
	// Key or Passphrase
	encodingKey string

	// Map to set a cipher text with some key
	keyValues map[string]string

	//To store
	filePath string
}

func (v *Vault) setKeyValues() error {

	var sb bytes.Buffer
	enc := json.NewEncoder(&sb)
	enc.Encode(v.keyValues)

	f, err := os.OpenFile("/home/gslab/Coding/goworkspace/src/abc.txt", os.O_RDWR|os.O_CREATE, 0755)
	defer f.Close()
	if err != nil {
		return err
	}
	_, err = f.Write(sb.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func (v *Vault) loadKeyValues() error {

	fileInfo, err := os.Stat(v.filePath)
	if os.IsNotExist(err) {
		v.keyValues = make(map[string]string)
		return errors.New("No such directory or file found")
	}

	if fileInfo.Size() == 0 {
		v.keyValues = make(map[string]string)
		return errors.New("Empty file")
	}

	file, err := os.Open(v.filePath)
	defer file.Close()
	if err != nil {
		v.keyValues = make(map[string]string)
		return err
	}

	v.keyValues = make(map[string]string)

	dec := json.NewDecoder(file)
	err = dec.Decode(&v.keyValues)
	if err != nil {
		return errors.New(fmt.Sprintln("Error occured while decoding. Error desc:", err))
	}
	return nil
}

// Get Get metjod takes in key name and check whether there is cipher text with that key name.
// And if it is present then it decrypts it into plain text and returns it back
func (v *Vault) Get(keyName string) (string, error) {
	err := v.loadKeyValues()
	if err != nil {
		return "", err
	}

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

	v.loadKeyValues()
	encryptedValue, err := crypt.Encrypt(textToEncode, v.encodingKey)
	if err != nil {
		return err
	}
	v.keyValues[keyName] = encryptedValue
	err = v.setKeyValues()
	if err != nil {
		return err
	}
	return nil
}
