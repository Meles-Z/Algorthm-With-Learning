// package main

// import (
//         "crypto/rand"
//         "crypto/rsa"
//         "crypto/x509"
//         "encoding/pem"
//         "os"
// )

// func main() {
//         // Read the public key
//         publicKeyBytes, err := os.ReadFile("publick.pem")
//         if err != nil {
//                 panic(err)
//         }

//         block, _ := pem.Decode(publicKeyBytes)
//         if block == nil || block.Type != "RSA PUBLICK KEY" {
//                 panic("failed to decode PEM public key")
//         }

//         publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
//         if err != nil {
//                 panic(err)
//         }
//         rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
//         if !ok {
//                 panic("not an rsa public key")
//         }

//         // Encrypt the message
//         plaintext := []byte("This is my secret message")
//         ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, plaintext)
//         if err != nil {
//                 panic(err)
//         }

//         // Write the ciphertext to a file
//         err = os.WriteFile("ciphertext.bin", ciphertext, 0644)
//         if err != nil {
//                 panic(err)
//         }
// }


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
        // 1. Read the private key
        privateKeyBytes, err := os.ReadFile("private.pem")
        if err != nil {
                panic(err)
        }

        block, _ := pem.Decode(privateKeyBytes)
        if block == nil || block.Type != "RSA PRIVATE KEY" {
                panic("failed to decode PEM private key")
        }

        privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
        if err != nil {
                panic(err)
        }

        // 2. Read the encrypted data (ciphertext). For this example, we will encrypt and then decrypt a simple string.
        ciphertext, err := os.ReadFile("ciphertext.bin")
        if err != nil {
                panic(err)
        }

        // 3. Decrypt the ciphertext
        plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
        if err != nil {
                panic(err)
        }

        fmt.Printf("Decrypted: %s\n", plaintext)
}