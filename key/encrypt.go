package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {

	publickKeyPem, err := os.ReadFile("publick.pem")
	if err != nil {
		panic(err)
	}
	publickKeyBlock, _ := pem.Decode(publickKeyPem)
	publickKey, err := x509.ParsePKIXPublicKey(publickKeyBlock.Bytes)
	if err != nil {
		panic(err)
	}
	plainText := []byte("Hello World!")
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publickKey.(*rsa.PublicKey), plainText)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Encrypted: %x\n", cipherText)

}
