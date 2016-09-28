package cryptorsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"os"
)

var noLabel = []byte("")

// EncryptRSA encrypts any data with the RSA key
func EncryptRSA(data, privkp string) (string, error) {
	var (
		privKey rsa.PrivateKey
		err     error
	)

	if privKey, err = getPrivateKey(privkp); err != nil {
		return "", err
	}

	//hash := getHash()
	//fmt.Println("hash : ", hash)
	hash := sha1.New()
	//fmt.Println("hash1 : ", hash)
	//hash := sha1.New()
	msg := []byte(data)
	//r := rand.Reader

	encryptedData, err := rsa.EncryptOAEP(hash, rand.Reader, &privKey.PublicKey, msg, noLabel)
	if err != nil {
		return "", err
	}
	encodedData := base64.URLEncoding.EncodeToString(encryptedData)

	//return string(encryptedData), nil
	return encodedData, nil
}

func writeFile(data, file string) {
	f, fError := os.Create(file)
	if fError != nil {
		fmt.Println(fError)
	}
	defer f.Close()
	f.WriteString(data)
}

// DecryptRSA any data with the RSA key
func DecryptRSA(data, privkp string) (string, error) {
	var (
		privKey       rsa.PrivateKey
		encryptedData []byte
		err           error
	)

	if privKey, err = getPrivateKey(privkp); err != nil {
		return "", err
	}

	if encryptedData, err = base64.URLEncoding.DecodeString(data); err != nil {
		return "", err
	}
	//hash := getHash()
	//fmt.Println("hash : ", hash)
	hash := sha1.New()
	//fmt.Println("hash1 : ", hash)
	decryptedData, err := rsa.DecryptOAEP(hash, rand.Reader, &privKey, encryptedData, noLabel)
	if err != nil {
		return "", err
	}
	return string(decryptedData), nil
}
