package cryptorsa

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	testmsg := "this is a test"
	privKeyPath := "/home/chadit/Projects/src/github.com/chadit/GoSamples/cryptotest/cryptorsa/test_data/private_key"
	r, err := EncryptRSA(testmsg, privKeyPath)
	if err != nil {
		fmt.Println(err)
	}

	if r == "" {
		t.Fatal("no data returned")
	}

	rd, rderr := DecryptRSA(r, privKeyPath)
	if rderr != nil {
		fmt.Println(rderr)
	}
	//fmt.Println("decrypted string : ", rd)

	if testmsg != rd {
		t.Fatal("no data returned")
	}
}

func TestDecrypt(t *testing.T) {
	testmsg := "this is a test"
	encryptedmsg := "Ci_8qH-2TN9m_EaxkdLFUzoihxL7llZEkk6LNIoGL7VTd5ngKzpogT9ReCIFNv4R3adT9QWC-7-lRtl12H8-h6qiRex3PH7Y1WbOxsKhHaEv5DrsNIrXrN8qYFsfvY3fRmqA2cPtwyqgdWmatyDneIGVFWha4Ya7olHAHUWhPe-FJe9Z1WQgerB5NCSl5GmQWbtQUncKC97OJwMD3pRpXPpxbibzyoLEwMNPjlD6Qkg_KNcKNkxv_UpcF9sFFSeu3a14qIPBwLJMMZEUKRidRjzBgiCVUBRSTt2pK1iEwz4zxUsE31Bk3zpAeY-eWOBf0Kgxp6ELXrUjgH75fwap3qYxSM0Z7wmmgSH3wJqOfNGJxGs0xoktG1EwTVB5wfuYxxq1brpRIRH2vtfUJ-np9UNnUsUlriEENj8X8wQYDt4KsagTY3XJ8xa5ymjuegpkv_QPgUC5p3VHHO2AMYI8Jjjqg1OiUyhSc-_nXAW9XzQ66-4Sy8ClfWHrFIe8ZbitsmAh1BzEXNy5PshN46XAl7BWmNHzCdUmnZJCtoAcOZjQ-5h8WJvsOPooFOAykzUqYnXMgSW4C2UTyJ4m3r6qyG5eMOOllUFM2xS36UNLPqlqQQIwbu_sfMOZjwk-UjRJMBwRcJ2_Ij7rMkuti0lS-D3fSpJebhaCmzc="
	privKeyPath := "/home/chadit/Projects/src/github.com/chadit/GoSamples/cryptotest/cryptorsa/test_data/private_key"

	rd, rderr := DecryptRSA(encryptedmsg, privKeyPath)
	if rderr != nil {
		fmt.Println(rderr)
	}
	//fmt.Println("decrypted string111 : ", rd)

	if testmsg != rd {
		t.Fatal("no data returned")
	}

}
