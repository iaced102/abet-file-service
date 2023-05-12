package crypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

func generateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey) {
	// This method requires a random number of bits.
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// The public key is part of the PrivateKey struct
	return privateKey, &privateKey.PublicKey
}

// Export public key as a string in PEM format
func exportPubKeyAsPEMStr(pubkey *rsa.PublicKey) string {
	pubKeyPem := string(pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(pubkey),
		},
	))
	return pubKeyPem
}

// Export private key as a string in PEM format
func exportPrivKeyAsPEMStr(privkey *rsa.PrivateKey) string {
	privKeyPem := string(pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privkey),
		},
	))
	return privKeyPem

}

// Save string to a file
func saveKeyToFile(keyPem, filename string) {
	pemBytes := []byte(keyPem)
	ioutil.WriteFile(filename, pemBytes, 0400)
}

// Decode private key struct from PEM string
func ExportPEMStrToPrivKey(priv []byte) *rsa.PrivateKey {
	block, _ := pem.Decode(priv)
	key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	return key
}

// Decode public key struct from PEM string
func ExportPEMStrToPubKey(pub []byte) *rsa.PublicKey {
	block, _ := pem.Decode(pub)
	key, _ := x509.ParsePKCS1PublicKey(block.Bytes)
	return key
}

// Read data from file
func ReadKeyFromFile(filename string) []byte {
	key, _ := ioutil.ReadFile(filename)
	return key
}

func GetPrivateKey(filepath string) *rsa.PrivateKey {
	privKeyPEM := ReadKeyFromFile(filepath)
	privKeyFile := ExportPEMStrToPrivKey(privKeyPEM)
	return privKeyFile
}

func GetPrivateKeyByte(filepath string) []byte {
	privKeyPEM := ReadKeyFromFile(filepath)
	return privKeyPEM
}

func GetPublicKey(filepath string) *rsa.PublicKey {
	pubKeyPEM := ReadKeyFromFile(filepath)
	pubKeyFile := ExportPEMStrToPubKey(pubKeyPEM)
	return pubKeyFile
}
func EncryptByPublicKey(pubKeyFile *rsa.PublicKey, message []byte) string {

	cipherText, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		pubKeyFile,
		message,
		[]byte(""),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(cipherText, "                   cipherText           ")
	return base64.StdEncoding.EncodeToString(cipherText)
}
func DecryptByPrivateKey(privateKey *rsa.PrivateKey, cipherText string) string {
	text, e := base64.StdEncoding.DecodeString(cipherText)
	fmt.Println(text, "            texttexttext           ")
	if e != nil {
		panic(e)
	}
	decMessage, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, text, []byte(""))
	if err != nil {
		panic(err)
	}
	fmt.Println(decMessage)
	return string(decMessage)
}
func CreateAndSaveKeyPair() {
	// Generate a 2048-bits key
	privateKey, publicKey := generateKeyPair(2048)
	// Create PEM string

	privKeyStr := exportPrivKeyAsPEMStr(privateKey)
	pubKeyStr := exportPubKeyAsPEMStr(publicKey)
	fmt.Println(privKeyStr)

	saveKeyToFile(privKeyStr, "crypt/privkey.pem")
	saveKeyToFile(pubKeyStr, "crypt/pubkey.pem")
}
func CreateAndSaveKeyPairV2() {
	// Generate a 2048-bits key
	privateKey, publicKey := generateKeyPair(2048)
	// Create PEM string

	privKeyStr := exportPrivKeyAsPEMStr(privateKey)
	pubKeyStr := exportPubKeyAsPEMStr(publicKey)
	saveKeyToFile(privKeyStr, "crypt/privkeyv2.pem")
	saveKeyToFile(pubKeyStr, "crypt/pubkeyv2.pem")
}
