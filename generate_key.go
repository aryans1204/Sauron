package main

import (
	"crypto/rand"
	"fmt"
	"os"
)

func main() {
	var keypath string
	var cipher string
	if len(os.Args) > 1 {
		keypath = os.Args[1]
		cipher = os.Args[2]
	} else {
		home_var := os.Getenv("HOME")
		keypath = home_var
		cipher = "AES"
	}

	if cipher == "AES" {
		key := make([]byte, 32) //256 bit key for Rijndael
	} else {
		key := make([]byte, 16) //128 bit keysize for Blowfish and Feistel
	}

	_, err := rand.Read(key) //create 16 bytes or 256 bits long key for use with AES cipher
	if err != nil {
		fmt.Println("There seems to be a problem, try again")
	}

	err = os.WriteFile(keypath+"/key.txt", []byte(key), 0777)

}
