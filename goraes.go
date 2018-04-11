package main

////////////////////////////////////////////////////////////////////////////////
// Copyright (c) 2018, Gianluca Fiore
//
//    This program is free software: you can redistribute it and/or modify
//    it under the terms of the GNU General Public License as published by
//    the Free Software Foundation, either version 3 of the License, or
//    (at your option) any later version.
//
////////////////////////////////////////////////////////////////////////////////


import (
	"fmt"
	"os"
	"io"
	"io/ioutil"
	"flag"
	"log"
	_ "encoding/json"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"

	_ "github.com/zfeldt/gencrypt"
	"github.com/julienroland/copro/prompt"
)

const (
	// some constants to colorize terminal output with ANSI color codes
	ansiRed = "\033[0;31m"
	ansiBlue = "\033[1;34m"
	ansiReset = "\033[0m"
)

var SearchKey, InFile, OutFile, Password string
var Decrypt, Encrypt bool

// parse flags
func init() {
	var usageMessage  = `
goraes [-i|-inputfile <file>] [-o|-outputfile <file>] [-s|-searchkey <word>] [-d|-decrypt|-e|-encrypt]

goraes encrypts/decrypts a JSON file containing login credentials.

Arguments:
	-s|-searckey <word>
		Search for matching account names

	-i|-inputfile <file>
		The input file

	-o|-outputfile <file>
		The output file

	-d|-decrypt
		Set program to decrypt mode

	-e|-encrypt
		Set program to encrypt mode

	-p|-password
		Give encryption/decryption password directly on the command line

`

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usageMessage)
	}

	const (
		defSearchKey	= ""
		defInFile		= "t.json"
		defOutFile		= "/tmp/t"
		defDecrypt		= false
		defEncrypt		= false
		defPassword		= ""
	)

	flag.StringVar(&SearchKey, "searchkey", defSearchKey, "")
	flag.StringVar(&SearchKey, "s", defSearchKey, "")
	flag.StringVar(&InFile, "inputfile", defInFile, "")
	flag.StringVar(&InFile, "i", defInFile, "")
	flag.StringVar(&OutFile, "outputfile", defOutFile, "")
	flag.StringVar(&OutFile, "o", defOutFile, "")
	flag.BoolVar(&Decrypt, "d", defDecrypt, "")
	flag.BoolVar(&Decrypt, "--decrypt", defDecrypt, "")
	flag.BoolVar(&Encrypt, "e", defEncrypt, "")
	flag.BoolVar(&Encrypt, "--encrypt", defEncrypt, "")
	flag.StringVar(&Password, "p", defPassword, "")
	flag.StringVar(&Password, "password", defPassword, "")


}

func openFile(f string) *os.File {
	fl, err :=  os.OpenFile(f, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
	if err != nil {
		log.Fatal(err)
	}

	return fl
}

func askForPassword() string {
	// this uses julienroland/copro/prompt library
	ask := prompt.NewPassword()
	ask.Question = "Enter Password"

	result, err := ask.Run()
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func checkPwdLength(s string) string {
	// check that password is at least 32 bytes. If not, add enough characters 
	// to make it 32 bytes long.
	// Clearly this is no NSA-safe but it'll suffice...
	if len(s) < 32 {
		difference := 32 - len(s)
		var pad string
		for i := 0; i < difference; i++ {
			pad += "F"
		}
		return s + pad
	}
	// password is at least 32 bytes long, return it untouched
	return s
}



func main() {
	// initialize cli arguments
	flag.Parse()
	var keyForEncryption []byte

	// Check if we have the minimum amount of parameters and not both have been 
	// set at the same time
	if (Decrypt == false && Encrypt == false) || (Decrypt == true && Encrypt == true) {
		flag.Usage()
		os.Exit(1)
	}

	// Load file to encrypt in memory
	file := openFile(InFile)
	defer file.Close()
	fileToWork, rerr := ioutil.ReadAll(file)
	if rerr != nil {
		panic(rerr)
	}

	if Password == "" {
		// ask the user for a password
		keyForEncryption = []byte(checkPwdLength(askForPassword()))
	} else {
		keyForEncryption = []byte(checkPwdLength(Password))
	}

	if Encrypt {
		// Encrypt the data
		c, cerr := aes.NewCipher(keyForEncryption)
		if cerr != nil {
			panic(cerr)
		}

		gcm, err := cipher.NewGCM(c)
		if err != nil {
			panic(err)
		}

		nonce := make([]byte, gcm.NonceSize())
		if _, nerr := io.ReadFull(rand.Reader, nonce); nerr != nil {
			panic(nerr)
		}
		encryptedText := gcm.Seal(nonce, nonce, fileToWork, nil)

		werr := ioutil.WriteFile(OutFile, encryptedText, 0644)
		if werr != nil {
			panic(werr)
		}

		os.Exit(0)
	} else {
		// Decrypt the data
		c, cerr := aes.NewCipher(keyForEncryption)
		if cerr != nil {
			panic(cerr)
		}

		gcm, gerr := cipher.NewGCM(c)
		if gerr != nil {
			panic(gerr)
		}

		nonceSize := gcm.NonceSize()
		if len(fileToWork) < nonceSize {
			panic("ciphertext too short")
		}

		nonce, decryptedText := fileToWork[:nonceSize], fileToWork[nonceSize:]
		decryptedText, derr := gcm.Open(nil, nonce, decryptedText, nil)
		if derr != nil {
			panic(derr)
		}

		// write back file in plaintext format
		werr := ioutil.WriteFile(OutFile, decryptedText, 0644)
		if werr != nil {
			panic(werr)
		}
		os.Exit(0)
	}

}

