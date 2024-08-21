package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	/* genKey() */
	login()
}

func login() {
	// Generate a new RSA key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Error generating RSA key pair:", err)
		return
	}
	fmt.Printf("privateKey: %v\n", privateKey)

	// Create a public key from the generated private key
	publicKey := &privateKey.PublicKey

	fmt.Printf("publicKey: %v\n", publicKey)

	// Create some data to be signed
	data := []byte("anousone@gmail.com-device11223-948281912")

	// Hash the data using SHA-256
	hashed := sha256.Sum256(data)

	// Sign the hashed data using the private key
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Println("Error signing data:", err)
		return
	}
	fmt.Printf("signature: %v\n", signature)

	// Verify the signature using the public key
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		fmt.Println("Signature verification failed:", err)
		return
	}

	fmt.Println("Signature verification successful!")
}
