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
	"encoding/json"
	"crypto/aes"
)

const (
	// some constants to colorize terminal output with ANSI color codes
	ansiRed = "\033[0;31m"
	ansiBlue = "\033[1;34m"
	ansiReset = "\033[0m"
)

var usageMessage  = `

`

// parse flags
func init() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usageMessage)
	}

	const (
		defSearchKey	= ""
		defInFile		= "t.json"
		defOutFile		= "/tmp/t"
	)

	flag.StringVar(&SearchKey, "searchkey", defSearchKey, "")
	flag.StringVar(&SearchKey, "s", defSearchKey, "")
	flag.StringVar(&InFile, "inputfile", defInFile, "")
	flag.StringVar(&InFile, "i", defInFile, "")
	flag.StringVar(&OutFile, "outputfile", defOutFile, "")
	flag.StringVar(&OutFile, "o", defOutFile, "")

}

func main() {
	// initialize cli arguments
	flag.Parse()
}

