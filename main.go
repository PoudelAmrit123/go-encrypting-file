package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/PoudelAmrit123/goFleEncryption/filecrypt"
	"github.com/gohugoio/hugo/common/terminal"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	fmt.Println("Hello Golang!")

	if len(os.Args) < 2 {
		printHelp()
	}

	function := os.Args[1]

	switch function {
	case "help":
		printHelp()
	case "encrypt":
		encryptHandle()
	case "decrypt":
		decryptHandle()
	default:
		fmt.Println("Run encrypt to encrypt a file , and decrypt to decrypt the file ")
		os.Exit(1)

	}
}

func printHelp() {

	fmt.Println("File encryption")
	fmt.Println("Simple file encryptor for your day-to-day needs.")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("\tgo run . encrypt /path/to/your/file")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("\t encrypt \t Encrypt a file given a password")
	fmt.Println("\t decrypt \t Tries to decrypt a file using a password")
	fmt.Println("\t help \t\t Display help text")
	fmt.Println("")

}

func encryptHandle() {
	//if the user pass go run . encrypt
	// that means the length is less than 3
	if len(os.Args) < 3 {
		println("missing the path to the file . for more info , run  go run . help  ")
	}
	file := os.Args[2]
	if !validateFile(file) {
		panic("File not found")
	}
	password := getPassword()
	fmt.Println("\n Encrypting.....")
	filecrypt.Encryption(file, password)
	fmt.Println("\n File successfully Protected")

}

func decryptHandle() {

}

func getPassword() {
	fmt.Println("Enter the password")

	password := terminal.ReadPassword(0)
	fmt.Println("Confirm password")
	passwordConfirm := terminal.ReadPassword(0)
	if !validatePassword(password, passwordConfirm) {
		fmt.Println("\nPassword do not match. Please try again \n")
		return getPassword()
	}
	return password

}

func validatePassword(password []byte, password2 []byte) bool {

	if !bytes.Equal(password, password2) {
		return false
	}
	return true
}

func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true

}
