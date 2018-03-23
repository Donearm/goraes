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
	"encoding/json"
	"os"
	"fmt"
)

type Paths struct {
	OutFile		string
	InFile		string
}

// Load config file
func LoadConfig() Paths {
	var config Paths

	file, oErr := os.Open("conf.json")
	if oErr != nil {
		panic("Couldn't open config file!")
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config = Paths{}

	err := decoder.Decode(&config)
	if err != nil {
		panic(err)
	}

	fmt.Println(config)

	return config
}