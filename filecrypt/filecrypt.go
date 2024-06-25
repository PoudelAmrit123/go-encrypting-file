package filecrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"

	"golang.org/x/crypto/pbkdf2"
)

func Encrypt(source string, password []byte) {
	if _, err := os.Stat(source); os.IsNotExist(err) {
		panic(err.Error())
	}
	//Opening the file and closing the file
	srcFile, err := os.Open(source)
	if err != nil {
		panic(err.Error())
	}
	defer srcFile.Close()
	//read the file
	plaintext, err := io.ReadAll(srcFile)
	if err != nil {
		panic(err.Error())

	}

	key := password
	//creating the empty nounce and randomizing the nounce

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	dk := pbkdf2.Key(key, nonce, 4096, 32, sha1.New)
	cipherBlock, err := aes.NewCipher(dk)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		panic(err.Error())
	}

	cipherText := aesgcm.Seal(key, nonce, plaintext, nil)
	//we append the nonce in the cipher text for the decryption part
	cipherText = append(cipherText, nonce...)

	dstFile, err := os.Create(source)
	Errors(err)
	defer dstFile.Close()

	_, err = dstFile.Write(cipherText)
	Errors(err)

}

func Errors(err error) error {
	panic(err.Error())

}

func Decrypt(source string, password []byte) {

	if _, err := os.Stat(source); os.IsNotExist(err) {
		panic(err.Error())
	}
	srcFile, err := os.Open(source)
	Errors(err)
	defer srcFile.Close()
	//reading the source file
	cipherText, err := io.ReadAll(srcFile)
	Errors(err)

	key := password
	//finding out the nounce that is appended in the encryption

	salt := cipherText[len(cipherText)-12:]

	str := hex.EncodeToString(salt)
	nonce, err := hex.DecodeString(str)
	if err != nil {
		panic(err.Error())

	}
	dk := pbkdf2.Key(key, nonce, 4096, 32, sha1.New)
	cipherblock, err := aes.NewCipher(dk)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(cipherblock)
	if err != nil {
		panic(err.Error())
	}
	plainText, err := aesgcm.Open(key, nonce, cipherText[:len(cipherText)-12], nil)
	if err != nil {
		panic(err.Error())

	}

	dstFile, err := os.Create(source)
	if err != nil {
		panic(err.Error())

	}
	_, err = dstFile.Write(plainText)
	if err != nil {
		panic(err.Error())

	}

	defer dstFile.Close()

}
