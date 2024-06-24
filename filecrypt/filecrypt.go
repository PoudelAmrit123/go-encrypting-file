package filecrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
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

	_, err = dstFile.Write(cipherText)
	Errors(err)

}

func Errors(err error) error {
	panic(err.Error())

}

func Decrypt(source string, password []byte) {

}
