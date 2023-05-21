package main

import (
	"crypto/rand"
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

		rand.Read(key)
		os.WriteFile(keypath+"/key.txt", []byte(key), 0777)
	} else {
		key := make([]byte, 16) //128 bit keysize for Blowfish and Feistel
		rand.Read(key)
		os.WriteFile(keypath+"/key.txt", []byte(key), 0777)

	}

}
