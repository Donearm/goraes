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
	"io/ioutil"
	"flag"
	"log"
	_ "encoding/json"

	"github.com/zfeldt/gencrypt"
	"github.com/julienroland/copro/prompt"
)

const (
	// some constants to colorize terminal output with ANSI color codes
	ansiRed = "\033[0;31m"
	ansiBlue = "\033[1;34m"
	ansiReset = "\033[0m"
)

var SearchKey, InFile, OutFile string
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

/*	if SearchKey == "" && Decrypt == false && Encrypt == false {
/*		flag.Usage();
/*	} */

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


func main() {
	// initialize cli arguments
	flag.Parse()
	//textToEncrypt := "abcdeFUUUUUUUU"
	//keyForEncryption := []byte("example key 1234")

	// Load file to encrypt in memory
	file := openFile(InFile)
	defer file.Close()
	fileToEncrypt, rerr := ioutil.ReadAll(file)
	if rerr != nil {
		panic(rerr)
	}

	// ask the user for a password
	keyForEncryption := []byte(askForPassword())

	// Get the GCM
	gcm, err := gencrypt.NewGCM(keyForEncryption)
	if err != nil {
		panic(err)
	}

	// Encrypt the data
	encryptedText, eerr := gcm.AESEncrypt(fileToEncrypt)
	if eerr != nil {
		panic(eerr)
	}

	// Decrypt the data
	/*
	decryptedText, derr := gcm.AESDecrypt(encryptedText)
	if derr != nil {
		panic(derr)
	}
	*/

	// write back file in encrypted format
	werr := ioutil.WriteFile(OutFile, encryptedText, 0644)
	if werr != nil {
		panic(err)
	}
}

