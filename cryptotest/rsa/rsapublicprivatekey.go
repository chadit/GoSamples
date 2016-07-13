package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

func main() {
	prikey, pubkey, _ := GeneratingKey(4048)
	writeFile(prikey, "C:\\temp\\private_key")
	writeFile(pubkey, "C:\\temp\\public_key.pub")
	// fmt.Println(prikey)
	// fmt.Println("--------------")
	// fmt.Println(pubkey)
	// f, fError := os.Create("C:\\temp\\private_key")
	// if fError != nil {
	// 	fmt.Println(fError)
	// }
	// defer f.Close()
	// f.WriteString(prikey)

	//fmt.Println(strings.Contains(prikey, "-----BEGIN RSA PRIVATE KEY-----"))
	//fmt.Println(strings.Contains(pubkey, "-----BEGIN RSA PUBLIC KEY-----"))
}

func writeFile(data, file string) {
	f, fError := os.Create(file)
	if fError != nil {
		fmt.Println(fError)
	}
	defer f.Close()
	f.WriteString(data)
}

// GeneratingKey generates an RSA keypair of the given bit size.
func GeneratingKey(bits int) (privateKey string, publicKey string, err error) {
	prikey, pubkey, err := GenerateKey(bits)
	return string(prikey), string(pubkey), err
}

// GenerateKey generates an RSA keypair of the given bit size.
func GenerateKey(bits int) (privateKey []byte, publicKey []byte, err error) {
	key, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	prikey := marshalPrivateKey(key)
	pubkey, err := marshalPublicKey(&key.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	return prikey, pubkey, nil
}

// marshalPrivateKey converts a private key to ASN.1 DER encoded form.
func marshalPrivateKey(privateKey *rsa.PrivateKey) []byte {
	marshaled := x509.MarshalPKCS1PrivateKey(privateKey)
	encoded := pem.EncodeToMemory(&pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   marshaled,
	})
	return encoded
}

// marshalPublicKey serialises a public key to DER-encoded PKIX format.
func marshalPublicKey(publicKey *rsa.PublicKey) ([]byte, error) {
	marshaled, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	encoded := pem.EncodeToMemory(&pem.Block{
		Type:    "RSA PUBLIC KEY",
		Headers: nil,
		Bytes:   marshaled,
	})
	return encoded, nil
}

// parsePrivateKey returns an RSA private key from its ASN.1 PKCS#1 DER encoded form.
func parsePrivateKey(privateKey []byte) (*rsa.PrivateKey, error) {
	decoded, _ := pem.Decode(privateKey)
	if decoded == nil {
		return nil, errors.New("private key error")
	}
	prikey, err := x509.ParsePKCS1PrivateKey(decoded.Bytes)
	if err != nil {
		return nil, err
	}
	return prikey, nil
}

// parsePublicKey parses a DER encoded public key.
// These values are typically found in PEM blocks with "BEGIN PUBLIC KEY".
func parsePublicKey(publicKey []byte) (*rsa.PublicKey, error) {
	decoded, _ := pem.Decode(publicKey)
	if decoded == nil {
		return nil, errors.New("public key error")
	}
	pubkey, err := x509.ParsePKIXPublicKey(decoded.Bytes)
	if err != nil {
		return nil, err
	}
	return pubkey.(*rsa.PublicKey), nil
}
