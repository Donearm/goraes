// Handle Goraes configuration file
package config

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
	"os/user"
)

// The paths struct. A InFile and OutFile, both strings, are the only elements
type Paths struct {
	InFile		string	`json:"InFile"`
	OutFile		string	`json:"OutFile"`
}

// Get user's HOME directory
func getHomeDir() string {
	u, err := user.Current()
	if err != nil {
		return ""
	}
	return u.HomeDir
}

// Load the config file
func LoadConfig() Paths {
	var config Paths
	var homedir string

	homedir = getHomeDir()

	file, oErr := os.Open(homedir + "/.goraes.json")
	if oErr != nil {
		// there's no config file or it is not readable. Skip it...
		return config
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config = Paths{}

	err := decoder.Decode(&config)
	if err != nil {
		// File isn't json or is wrong json, return nothing
		return config
	}

	return config
}
