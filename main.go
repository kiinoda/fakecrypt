package main

import "github.com/kiinoda/fakecrypt/encrypt"
import "github.com/kiinoda/fakecrypt/decrypt"
import "fmt"

func main() {
	str := "KodeKloud"
	fmt.Println("Original  string:", str)

	encryptedStr := encrypt.Nimbus(str)
	fmt.Println("Encrypted string:", encryptedStr)

	decryptedStr := decrypt.Nimbus(encryptedStr)
	fmt.Println("Decrypted string:", decryptedStr)
}
