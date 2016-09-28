package cryptorsa

import (
	"crypto/rsa"
	"crypto/sha1"
	"hash"
	"io/ioutil"
)

// getPrivateKey returns private the key from a file
func getPrivateKey(privkp string) (rsa.PrivateKey, error) {
	var (
		kd  []byte
		err error
		key rsa.PrivateKey
	)
	if kd, err = ioutil.ReadFile(privkp); err != nil {
		return key, err
	}
	if key, err = parseRSAPrivateKeyFromPEM(kd); err != nil {
		return key, err

	}
	return key, nil
}

// getPublicKey returns the public key from a file
func getPublicKey(pubkp string) (rsa.PublicKey, error) {
	var (
		kd  []byte
		err error
		key rsa.PublicKey
	)

	if kd, err = ioutil.ReadFile(pubkp); err != nil {
		return key, err
	}
	if key, err = parseRSAPublicKeyFromPEM(kd); err != nil {
		return key, err

	}
	return key, nil
}

func getHash() hash.Hash {

	h := sha1.New()
	h.Write([]byte("00112233445566778899"))
	//return hex.EncodeToString(h.Sum(nil))
	return h
}
