package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	//The key argument should be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
	aesEncrypt()
	bsString := "$2a$10$u0ir89cBxDlapUnU35nOl.3Txb3ZhPL0Ew/9n9qrglyqepKJN3oP6"
	bsString1 := "$2a$10$Cva6JkaDfJfxoKpg6NwTjOJO2ajquyuD6oPnUfhMiKu8uFXYZ7y7a"
	p := "mywifesnameandbirthday#1"
	p1 := "123456"

	// p := "todd"
	// // bcrypt.MinCost
	// bs, _ := bcrypt.GenerateFromPassword([]byte(p), 10)
	// fmt.Printf("PASSWORD ONE: %x \n", bs)
	//
	// p2 := "todd"
	// bs2, _ := bcrypt.GenerateFromPassword([]byte(p2), 10)
	// fmt.Printf("PASSWORD TWO: %x \n", bs2)

	bs, _ := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	//fmt.Println(bs)
	bsStringAuto := string(bs)
	fmt.Println(bsStringAuto)

	err := bcrypt.CompareHashAndPassword(bs, []byte(p))
	if err != nil {
		fmt.Println("Doesn't match")
	} else {
		fmt.Println("match")
	}

	err1 := bcrypt.CompareHashAndPassword([]byte(bsString), []byte(p))
	if err1 != nil {
		fmt.Println("Doesn't match")
	} else {
		fmt.Println("match")
	}

	err2 := bcrypt.CompareHashAndPassword([]byte(bsString1), []byte(p1))
	if err2 != nil {
		fmt.Println("Doesn't match")
	} else {
		fmt.Println("match")
	}
}

func aesEncrypt() {
	block, iv := getBlockAndIv()
	str := []byte("Hello World!")
	encrypter := cipher.NewCFBEncrypter(block, iv)
	encrypted := make([]byte, len(str))
	encrypter.XORKeyStream(encrypted, str)

	bsStringAuto := base64.URLEncoding.EncodeToString(encrypted)
	fmt.Println("encryptedString", bsStringAuto)

	//	fmt.Printf("%s \nencrypted to \n%v\n", str, encrypted)

	// decrypt

	// decrypter := cipher.NewCFBDecrypter(block, iv) // simple!
	//
	// decrypted := make([]byte, len(str))
	// decrypter.XORKeyStream(decrypted, encrypted)
	//
	// fmt.Printf("%v decrypt to %s\n", encrypted, decrypted)

}

func aesDecrypt() {

}

func getBlockAndIv() (cipher.Block, []byte) {
	key := "54a0106563fcd718944080f7"
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// 16 bytes for AES-128, 24 bytes for AES-192, 32 bytes for AES-256
	ciphertext := []byte("54a01b3863fcd61790afa4dd")
	iv := ciphertext[:aes.BlockSize] // const BlockSize = 16

	return block, iv

}
