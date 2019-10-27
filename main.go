package main

import (
	"fmt"
	"vault"
)

func main() {
	key := "I Love My India!"

	v := vault.NewVault(key)
	err := v.Set("username", "P95046384")
	if err != nil {
		fmt.Println("Error:", err)
	}

	originalText, err := v.Get("username")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Original text:", originalText)
	// encryptedText, _ := crypt.Encrypt("Hello there!", key)
	// fmt.Println("Encrypted text:", encryptedText)
	// originalText, _ := crypt.Decrypt(encryptedText, key)
	// fmt.Println("Decrypted to:", originalText)
}
