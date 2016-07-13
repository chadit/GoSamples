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

var (
	folderPath string
)

func getFolderPath() string {
	var f string
	fmt.Print("Please enter a folder to save files (Ex: c:\\Temp) : ")
	if _, err := fmt.Scan(&f); err != nil {
		fmt.Println("input error due to : ", err)
		getFolderPath()
	}

	if f == "" {
		f = "C:\\Temp"
	}
	folderPath = f
	return f
}

func getInput() int {
	fmt.Print("Please enter an integer (Ex: 4048) : ")
	var i int
	if _, err := fmt.Scan(&i); err != nil {
		fmt.Println("input error due to : ", err)
		getInput()
	}
	return i
}

func main() {
	createFolder(getFolderPath())

	prikey, pubkey, _ := GeneratingKey(getInput())
	writeFile(prikey, folderPath+"\\private_key")
	writeFile(pubkey, folderPath+"\\public_key.pub")
	fmt.Println("RSA keys where created at : ", folderPath)
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

// fileFolderExists returns whether the given file or directory exists or not
func fileFolderExists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			// FileFolderExists does not exist
			return false, nil
		}
		return true, nil
	}
	return true, nil
}

// createFolder create a folder if it does not already exist
func createFolder(path string) (bool, error) {
	results, err := fileFolderExists(path)
	if err != nil {
		return false, err
	}

	if results == false {
		fileErr := os.MkdirAll(path, 0x711)
		if fileErr != nil {
			return false, err
		}
		return true, nil
	}
	return false, err
}
