package main

// import (
// 	"crypto/rand"
// 	"crypto/rsa"
// 	"crypto/x509"
// 	"encoding/pem"
// 	"os"
// )

// func main() {
// 	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
// 	if err != nil {
// 		panic(err)
// 	}

// 	publickKey := &privateKey.PublicKey
// 	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
// 	privateKeyPem := pem.EncodeToMemory(&pem.Block{
// 		Type:  "RSA PRIVATE KEY",
// 		Bytes: privateKeyBytes,
// 	})
// 	err = os.WriteFile("private.pem", privateKeyPem, 0644)
// 	if err !=nil{
// 		panic(err)
// 	}

// 	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publickKey)
// 	if err != nil {
// 		panic(err)
// 	}

// 	publicKeyPem := pem.EncodeToMemory(&pem.Block{
// 		Type:  "RSA PUBLICK KEY",
// 		Bytes: publicKeyBytes,
// 	})

// 	err = os.WriteFile("publick.pem", publicKeyPem, 0644)
// 	if err != nil {
// 		panic(err)
// 	}

// }
